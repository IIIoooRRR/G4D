package g4d

import (
	"fmt"
	"sync"

	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/parse/types"
)

func prepareCommand(cmds []CommandTemplate) map[string][]CommandTemplate {
	CmdMap := make(map[string][]CommandTemplate)
	for _, command := range cmds {
		CmdMap[command.Trigger] = append(CmdMap[command.Trigger], command)
	}
	return CmdMap
}

func (b *Bot) staticEventProcessor(seq int, limiter chan struct{}) {
	cmdMap := prepareCommand(b.CommandBuffer)
	for {
		select {
		case <-b.ctx.Done():
			b.Logger.Info(fmt.Sprintf("%d processor stopped", seq))
			return
		case event := <-b.Gateway.Queue:
			ctx := b.newCtx()
			wg := sync.WaitGroup{}
			factory := types.Get(event.Type)
			if factory == nil {
				b.Gateway.EventPool.Put(event.Reset())
				continue
			}
			cmds := cmdMap[event.Type]
			b.eventCache.AddEvent(event, &wg, seq, len(cmds), factory)
			for _, cmd := range cmds {
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
