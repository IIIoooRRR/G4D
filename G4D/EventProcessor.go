package G4D

import (
	"log"

	"github.com/IIIoooRRR/G4D/Connect"
)

func sortCommand(cmds []Command) map[string][]Command {
	CmdMap := make(map[string][]Command)
	for _, command := range cmds {
		CmdMap[command.Trigger] = append(CmdMap[command.Trigger], command)
	}
	return CmdMap
}

/*
func (b *Bot) StaticEventProcessor() {
	Buffer := b.CommandBuffer
	for event := range b.Gateway.Queue {
		for _, command := range Buffer {
			if event.Type != command.Trigger {
				continue
			}
			go func(cmd Command, event *Connect.RawEvent) {
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
			go func(command Command, event *Connect.RawEvent) {
				err := command.Action.Execute(event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(command, event)
		}
	}
}

func (b *Bot) DynamicEventProcessor() {
	for event := range b.Gateway.Queue {

		for _, command := range b.CommandBuffer {
			if event.Type != command.Trigger {
				continue
			}
			go func(cmd Command, event *Connect.RawEvent) {
				err := cmd.Action.Execute(event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(command, event)
		}
	}
}
