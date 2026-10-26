package g4d

import (
	"fmt"
	"sync"

	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/parse/types"
)

func (b *Bot) dynamicEventProcessor(seq int, limiter chan struct{}) {
	for {
		select {
		case <-b.ctx.Done():
			b.Logger.Info(fmt.Sprintf("%d processor stopped", seq))
			return
		case event := <-b.Gateway.Queue:
			ctx := b.newCtx()
			wg := sync.WaitGroup{}

			t := types.Get(event.Type)
			if t == nil {
				b.Gateway.EventPool.Put(event.Reset())
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

			b.eventCache.AddEvent(event, &wg, seq, len(activeCmd), t)
			for _, cmd := range activeCmd {
				limiter <- struct{}{}
				go func(cmd CommandTemplate, event *parse.RawEvent) {
					defer func() { <-limiter }()
					b.initCommand(cmd, event, ctx, &wg)
				}(cmd, event)
			}
			wg.Wait()
			b.Gateway.EventPool.Put(event.Reset())
		}
	}
}
