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

func (b *Bot) staticEventProcessor(seq int, limiter *chan struct{}) {
	cmdMap := prepareCommand(b.CommandBuffer)
	for event := range b.Gateway.Queue {
		ctx := b.newCtx()
		wg := sync.WaitGroup{}
		eventType := types.Get(event.Type)
		if eventType == nil {
			continue
		}
		b.eventCache.AddEvent(event, &wg, seq, len(cmdMap[event.Type]), eventType)
		for _, cmd := range cmdMap[event.Type] {
			go func(cmd CommandTemplate, event *parse.RawEvent) {
				*limiter <- struct{}{}
				defer func() { <-*limiter }()
				b.initCommand(cmd, event, &ctx)
			}(cmd, event)
		}
		wg.Wait()
	}
}
