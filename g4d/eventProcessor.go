package g4d

import (
	"runtime/debug"
	"sync"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/parse/types"
	"go.uber.org/zap"
)

func prepareCommand(cmds []CommandTemplate) map[string][]CommandTemplate {
	CmdMap := make(map[string][]CommandTemplate)
	for _, command := range cmds {
		CmdMap[command.Trigger] = append(CmdMap[command.Trigger], command)
	}
	return CmdMap
}

func (b *Bot) InitCommand(command CommandTemplate, event *gateway.RawEvent, logger *zap.Logger) {
	defer func() {
		if r := recover(); r != nil {
			if b.PanicHandler != nil {
				(*b.PanicHandler).OnPanic(event, &command, r, debug.Stack())
			}
		}
	}()
	err := command.Action.Execute(event)
	if err != nil {
		logger.Error("cmd execution error ", zap.Error(err))
	}
}

func (b *Bot) StaticEventProcessor(limitSize uint) {
	limiter := make(chan struct{}, limitSize)
	cmdMap := prepareCommand(b.CommandBuffer)
	for event := range b.Gateway.Queue {
		wg := sync.WaitGroup{}
		eventType := types.Get(event.Type)
		if eventType == nil {
			continue
		}
		parse.AddEvent(event, &wg, len(cmdMap[event.Type]), eventType)

		for _, cmd := range cmdMap[event.Type] {
			limiter <- struct{}{}
			go func(cmd CommandTemplate, event *gateway.RawEvent) {
				defer func() { <-limiter }()
				b.InitCommand(cmd, event, b.Logger)
			}(cmd, event)
		}
		wg.Wait()
		parse.DeleteEvent(event)
	}
}
func (b *Bot) DynamicEventProcessor(limitSize uint) {
	limiter := make(chan struct{}, limitSize)
	for event := range b.Gateway.Queue {
		wg := sync.WaitGroup{}
		eventType := types.Get(event.Type)
		if eventType == nil {
			continue
		}
		var activeCmd []CommandTemplate
		b.commandMu.Lock()
		for _, cmd := range b.CommandBuffer {
			if cmd.Trigger != event.Type {
				continue
			}
			activeCmd = append(activeCmd, cmd)
		}
		b.commandMu.Unlock()
		parse.AddEvent(event, &wg, len(activeCmd), eventType)
		for _, cmd := range activeCmd {
			limiter <- struct{}{}
			go func() {
				defer func() { <-limiter }()
				b.InitCommand(cmd, event, b.Logger)
			}()
		}
		wg.Wait()
		parse.DeleteEvent(event)
	}
}
