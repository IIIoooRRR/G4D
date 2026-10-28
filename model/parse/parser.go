//go:build !unsafe

package parse

import (
	"sync"

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
