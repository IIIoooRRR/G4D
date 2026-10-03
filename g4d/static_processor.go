package g4d

import (
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
			b.Logger.Info("processor stopped")
			return
		case event := <-b.Gateway.Queue:
			ctx := b.newCtx()
			wg := sync.WaitGroup{}
			eventType := types.Get(event.Type)
			if eventType == nil {
				continue
			}
			cmds := cmdMap[event.Type]
			b.eventCache.AddEvent(event, &wg, seq, len(cmds), eventType)
			for _, cmd := range cmds {
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
