package g4d

import (
	"sync"

	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/parse/types"
)

func dynamicEventProcessor(b *Bot, limitSize uint) {
	limiter := make(chan struct{}, limitSize)
	for event := range b.Gateway.Queue {
		wg := sync.WaitGroup{}
		eventType := types.Get(event.Type)
		if eventType == nil {
			continue
		}
		var activeCmd []CommandTemplate
		b.CommandMu.Lock()
		for _, cmd := range b.CommandBuffer {
			if cmd.Trigger != event.Type {
				continue
			}
			activeCmd = append(activeCmd, cmd)
		}
		b.CommandMu.Unlock()
		parse.AddEvent(event, &wg, len(activeCmd), eventType)
		for _, cmd := range activeCmd {
			limiter <- struct{}{}
			go func() {
				defer func() { <-limiter }()
				b.initCommand(cmd, event, b.Logger)
			}()
		}
		wg.Wait()
		parse.DeleteEvent(event)
	}
}
