package G4D

import (
	"log"

	"github.com/IIIoooRRR/G4D/Connect"
)

func (b *Bot) StaticEventProcessor() {
	Buffer := &b.CommandBuffer
	for event := range b.Gateway.Queue {
		for _, command := range *Buffer {
			if event.Type != command.Trigger {
				continue
			}
			go func(cmd *Command, event *Connect.RawEvent) {
				err := cmd.Action.Execute(event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(&command, event)
		}
	}
}
func (b *Bot) DynamicEventProcessor() {
	for event := range b.Gateway.Queue {
		for _, command := range b.CommandBuffer {
			if event.Type != command.Trigger {
				continue
			}
			go func(cmd *Command, event *Connect.RawEvent) {
				err := cmd.Action.Execute(event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(&command, event)
		}
	}
}
