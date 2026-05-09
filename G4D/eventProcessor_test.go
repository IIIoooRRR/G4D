package G4D

import (
	"log"

	"github.com/IIIoooRRR/G4D/Connect"
)

// DynamicEventProcessor reads b.CommandBuffer on every event.
// Allows adding/removing commands at runtime, but slightly slower.
// Use with sync.RWMutex if you modify buffer concurrently.

// StaticEventProcessor caches b.CommandBuffer once at start.
// Faster, but commands cannot be changed after launch.
func (b *Bot) ExampleBot_DynamicEventProcessor() {
	for event := range b.Gateway.Queue {
		for _, command := range b.CommandBuffer { //copies the current buffering command
			if event.Type != command.Trigger { //comparing it with the trigger declared during initialization
				continue
			}
			go func(cmd *Command, event *Connect.RawEvent) {
				err := cmd.Action.Execute(event) // executing the command
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(&command, event)
		}
	}
}

func (b *Bot) ExampleBot_StaticEventProcessor() {
	Buffer := b.CommandBuffer // copies the link to the buffer, which will then be referenced during operation
	for event := range b.Gateway.Queue {
		for _, command := range Buffer {
			if event.Type != command.Trigger { //comparing it with the trigger declared during initialization
				continue
			}
			go func(cmd *Command, event *Connect.RawEvent) {
				err := cmd.Action.Execute(event) // executing the command
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(&command, event)
		}
	}
}
