package g4d

import (
	"sync"

	"github.com/IIIoooRRR/G4D/model/gateway"
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

func staticEventProcessor(b *Bot, limitSize uint) {
	limiter := make(chan struct{}, limitSize)
	cmdMap := prepareCommand(b.CommandBuffer)

	for event := range b.Gateway.Queue {
		wg := sync.WaitGroup{}
		eventType := types.Get(event.Type)
		if eventType == nil {
			continue
		}
		parse.AddEvent(event, &wg, len(cmdMap[event.Type]), eventType)
		ctx := b.newCtx()
		for _, cmd := range cmdMap[event.Type] {
			limiter <- struct{}{}
			go func(cmd CommandTemplate, event *gateway.RawEvent) {
				defer func() { <-limiter }()
				b.initCommand(cmd, event, &ctx)
			}(cmd, event)
		}
		wg.Wait()
		parse.DeleteEvent(event)
		ctx.Cancel()
	}
}
