package parse

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/dependencies"
	"github.com/IIIoooRRR/G4D/model/schema"
	"go.uber.org/zap"
)

/*
I hope this code doesn’t scare you with its complexity.
I wanted to speed up the program so that each command wouldn’t parse the message on its own and, most importantly,
would NOT WAIT for other readers. so that all operations would be isolated from each other. I solved this using:
1. Wait groups. all commands must call GetEvent, which in itself makes wg.Done.
This is done as a weak defense against TOS Discord and to prevent the entire bot from freezing while waiting if some command contains complex calculation logic.
2. The parser itself. Starting from this commit, each bot has its own parser inside it. It is initialized during the Run function. All work is done without mutexes, since each processor is allocated its own cell during initialization. It works strictly with that cell.
3. Reflection. I hate it with all my heart.
I'm sorry. However, without it, there would be a lot of boilerplate code (switch-case, function wrapping),
but it is used only during parsing (simple reflection) and during compilation.
I WOULD LIKE TO POINT OUT
that if you want to add some structure that I didn’t implement (by accident) or forgot to do for parsing, use the methods in parse/types.
I wanted to be a stickler, so use add to add a new structure, and change to modify it. You can only replace it without adding.
*/

type Cache struct {
	Entry  []EventEntry
	noCopy noCopy
	logger *zap.Logger
}
type EventEntry struct {
	Data unsafe.Pointer
	Wg   *sync.WaitGroup
}

func InitCache(quantity _const.Quantity, logger *zap.Logger) *Cache {
	return &Cache{
		Entry:  make([]EventEntry, quantity),
		logger: logger,
	}
}

func GetEvent[T any](event *RawEvent) *T {
	defer (*event.cache)[event.idx].Wg.Done()
	return (*T)(
		(*event.cache)[event.idx].Data)
}

func (c *Cache) AddEvent(event *RawEvent, wg *sync.WaitGroup, seq, quantity int, t reflect.Type) {
	wg.Add(quantity)
	c.Entry[seq] = EventEntry{
		Data: c.reflectParsing(event, t),
		Wg:   wg,
	}
	//We set the values for the hidden fields so that we can retrieve the parsing values during reading.
	event.cache, event.idx = &c.Entry, seq
}
func (c *Cache) reflectParsing(event *RawEvent, t reflect.Type) unsafe.Pointer {
	d := reflect.New(t)
	err := Unmarshal(event.Data, d.Interface())
	if err != nil {
		c.logger.Error("unmarshal raw event", zap.Error(err))
		return nil
	}
	// #nosec G103
	return d.UnsafePointer()
}

/* Channel */
func ToChannel(body []byte) (*schema.Channel, error) {
	var channel *schema.Channel
	err := Unmarshal(body, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
func ToUser(body []byte) (*dependencies.User, error) {
	var user dependencies.User
	err := Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
