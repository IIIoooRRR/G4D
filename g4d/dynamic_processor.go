package g4d

import (
	"sync"

	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/parse/types"
)

func (b *Bot) dynamicEventProcessor(seq int, limiter chan struct{}) {
	for {
		select {
		case <-b.ctx.Done():
			b.Logger.Info("processor stopped")
			return
		case event := <-b.Gateway.Queue:
			ctx := b.newCtx()
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

			b.eventCache.AddEvent(event, &wg, seq, len(activeCmd), eventType)
			for _, cmd := range activeCmd {
				limiter <- struct{}{}
				go func(cmd CommandTemplate, event *parse.RawEvent) {
					defer func() { <-limiter }()
					b.initCommand(cmd, event, ctx)
				}(cmd, event)
			}
			wg.Wait()
		}
	}
}
