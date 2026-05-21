package g4d

import (
	"log"

	"github.com/IIIoooRRR/G4D/gateway"
)

// DynamicEventProcessor reads b.CommandBuffer on every event.
// Allows adding/removing debug_commands at runtime, but slightly slower.
// Use with sync.RWMutex if you modify buffer concurrently.

// StaticEventProcessor caches b.CommandBuffer once at start.
// Faster, but debug_commands cannot be changed after launch.
func (b *Bot) ExampleBot_DynamicEventProcessor() {
	for event := range b.Gateway.Queue {
		for _, command := range b.CommandBuffer { //copies the current buffering command
			if event.Type != command.Trigger { //comparing it with the trigger declared during initialization
				continue
			}
			go func(cmd *Command, event *gateway.RawEvent) {
				err := cmd.Action.Execute(event) // executing the command
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(&command, event)
		}
	}
}

func (b *Bot) ExampleBot_StaticEventProcessor() {
	commands := sortCommand(b.CommandBuffer)
	for event := range b.Gateway.Queue {
		for _, command := range commands[event.Type] {
			go func(command Command, event *gateway.RawEvent) {
				err := command.Action.Execute(event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(command, event)
		}
	}
}
