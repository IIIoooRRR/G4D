package parse

import (
	"maps"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/shema"
	"go.uber.org/zap"
)

/*
I hope this code doesn't terrify you because of its curvature.
I wanted to speed up the program so that each team would not parse the message on its own and, most importantly,
WOULD NOT WAIT for other readers. so that all operations are isolated from each other. I solved it through:
1. Waiting groups. all commands must call GetEvent, which itself does wg.Done.
This is done as a weak protection of the tos discord and so that the entire bot does not get up in anticipation
if some command carries heavy calculation logic.
2. Atomic pointers to the map.
All insertion/deletion operations are done via cas (compare and swap operations).
this is the most advantageous strategy with < 10 processor instances for events.
Although, I think you won't need it.
3. Reflection. I hate her with all my heart,
I'm sorry. however, without it,there would be a lot of boiler-plate (switch-case, function wrapping),
however, it is used once per parsing (easy reflection) and during compilation.
I WOULD LIKE TO POINT OUT,
so if you want to add some structure that I didn't implement (by accident) or forgot to do it for parsing, use the methods in parse/types.
I wanted to be a bore, so use add to add this - it will add a new structure, and change to change it, you can only replace it without adding
*/
var (
	cache atomic.Pointer[map[*gateway.RawEvent]eventEntry]
)

type eventEntry struct {
	Data any
	Wg   *sync.WaitGroup
}

func init() {
	mp := make(map[*gateway.RawEvent]eventEntry)
	cache.Store(&mp)

}
func GetEvent[T any](event *gateway.RawEvent) T {
	defer (*cache.Load())[event].Wg.Done()
	return (*cache.Load())[event].Data.(T)

}

func AddEvent(event *gateway.RawEvent, wg *sync.WaitGroup, quantity int, t reflect.Type) {
	entry := eventEntry{
		Data: reflectParsing(event, t),
		Wg:   wg,
	}
	casAdd[eventEntry](&cache, event, entry)
	wg.Add(quantity)
}

func DeleteEvent(event *gateway.RawEvent) {
	casDelete[eventEntry](&cache, event)
}

func reflectParsing(event *gateway.RawEvent, t reflect.Type) any {
	d := reflect.New(t).Interface()

	err := Unmarshal(event.Data, d)
	if err != nil {
		logger.Error("unmarshal raw event", zap.Error(err))
		return nil
	}
	return reflect.ValueOf(d).Elem().Interface()
}

func casDelete[V any](m *atomic.Pointer[map[*gateway.RawEvent]V], event *gateway.RawEvent) {
	for {
		oldMapPtr := m.Load()
		newMap := maps.Clone(*oldMapPtr)
		delete(newMap, event)
		if m.CompareAndSwap(oldMapPtr, &newMap) {
			break
		}
	}
}
func casAdd[V any](m *atomic.Pointer[map[*gateway.RawEvent]V], event *gateway.RawEvent, val V) {
	for {
		oldMapPtr := m.Load()
		newMap := maps.Clone(*oldMapPtr)
		newMap[event] = val
		if m.CompareAndSwap(oldMapPtr, &newMap) {
			break
		}
	}
}

/* Channel */
func ToChannel(body []byte) (*shema.Channel, error) {
	var channel *shema.Channel
	err := Unmarshal(body, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
