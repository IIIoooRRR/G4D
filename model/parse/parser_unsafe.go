//go:build unsafe

package parse

import (
	"sync"
	"unsafe"

	"go.uber.org/zap"
)

type EventEntry struct {
	Data unsafe.Pointer
	Wg   *sync.WaitGroup
}

// GetEvent returns a pointer to shared event data.
// Commands must NOT modify the returned value.
func GetEvent[T any](event *RawEvent) *T {
	return (*T)(event.cache[event.idx].Data)
}

func (c *Cache) reflectParsing(event *RawEvent, t any) unsafe.Pointer {
	err := Unmarshal(event.Data, t)
	if err != nil {
		c.logger.Error("unmarshal event data error", zap.Error(err))
	}
	return (*eface)(unsafe.Pointer(t)).Data
}
