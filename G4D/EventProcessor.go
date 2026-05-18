package G4D

import (
	"log"
	"runtime/debug"

	"github.com/IIIoooRRR/G4D/connect"
)

func sortCommand(cmds []Command) map[string][]Command {
	CmdMap := make(map[string][]Command)
	for _, command := range cmds {
		CmdMap[command.Trigger] = append(CmdMap[command.Trigger], command)
	}
	return CmdMap
}

func (b *Bot) InitCommand(command Command, event *connect.RawEvent) {
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

/*
func (b *Bot) StaticEventProcessor() {
	Buffer := b.CommandBuffer
	for event := range b.Gateway.Queue {
		for _, command := range Buffer {
			if event.Type != command.Trigger {
				continue
			}
			go func(cmd Command, event *connect.RawEvent) {
				err := cmd.Action.Execute(event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(command, event)
		}
	}
}
*/

func (b *Bot) StaticEventProcessor() {
	commands := sortCommand(b.CommandBuffer)
	for event := range b.Gateway.Queue {
		for _, command := range commands[event.Type] {
			go b.InitCommand(command, event)
		}
	}
}

func (b *Bot) DynamicEventProcessor() {
	for event := range b.Gateway.Queue {
		b.commandMu.Lock()
		bufCopy := b.CommandBuffer
		b.commandMu.Unlock()
		for _, command := range bufCopy {
			if event.Type != command.Trigger {
				continue
			}
			go b.InitCommand(command, event)
		}
	}
}
