//go:build !unsafe

package parse

import (
	"sync"

	"go.uber.org/zap"
)

type EventEntry struct {
	Data any
	Wg   *sync.WaitGroup
}

// GetEvent returns a pointer to shared event data.
// Commands must NOT modify the returned value.
func GetEvent[T any](event *RawEvent) *T {
	return event.cache[event.idx].Data.(*T)
}

func (c *Cache) reflectParsing(event *RawEvent, t any) any {
	err := Unmarshal(event.Data, t)
	if err != nil {
		c.logger.Error("unmarshal raw event", zap.String("event_type", event.Type), zap.Error(err), zap.Any("event", event.Data))
	}
	return t
}
