package g4d

import (
	"log"
	"runtime/debug"

	"github.com/IIIoooRRR/G4D/model/gateway"
)

func prepareCommand(cmds []CommandTemplate) map[string][]CommandTemplate {
	CmdMap := make(map[string][]CommandTemplate)
	for _, command := range cmds {
		CmdMap[command.Trigger] = append(CmdMap[command.Trigger], command)
	}
	return CmdMap
}

func (b *Bot) InitCommand(command CommandTemplate, event *gateway.RawEvent) {
	defer func() {
		if r := recover(); r != nil {
			if b.PanicHandler != nil {
				b.PanicHandler.OnPanic(event, &command, r, debug.Stack())
			}
		}
	}()
	err := command.Action.Execute(event)
	if err != nil {
		log.Println("[EVENT PROCESSOR] ", err)
	}
}

func (b *Bot) StaticEventProcessor(limitSize int) {
	limiter := make(chan struct{}, limitSize)
	commands := prepareCommand(b.CommandBuffer)
	for event := range b.Gateway.Queue {
		for _, command := range commands[event.Type] {
			limiter <- struct{}{}
			go func(command CommandTemplate, event *gateway.RawEvent) {
				defer func() { <-limiter }()
				b.InitCommand(command, event)
			}(command, event)
		}
	}
}

func (b *Bot) DynamicEventProcessor(limitSize int) {
	limiter := make(chan struct{}, limitSize)
	for event := range b.Gateway.Queue {
		var activeCmds []CommandTemplate
		b.commandMu.Lock()
		for _, command := range b.CommandBuffer {
			if event.Type == command.Trigger {
				activeCmds = append(activeCmds, command)
			}
		}
		b.commandMu.Unlock()
		for _, command := range activeCmds {
			limiter <- struct{}{}
			go func(localCmd CommandTemplate, event *gateway.RawEvent) {
				defer func() { <-limiter }()
				b.InitCommand(localCmd, event)
			}(command, event)
		}
	}
}
