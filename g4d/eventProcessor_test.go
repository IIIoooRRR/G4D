package g4d_test

import (
	"log"
	"runtime/debug"

	gateway2 "github.com/IIIoooRRR/G4D/model/gateway"
)

// StaticEventProcessor pre-sorts commands by event type at startup.
// Faster than dynamic, but commands cannot be added/modified at runtime.
//
// Parameters:
//   - limitSize: max concurrent command executions (0 = unlimited)
func (b *Bot) ExampleBot_StaticEventProcessor(limitSize uint) {
	// Pre-sort commands by trigger type for O(1) lookup
	commands := prepareCommandMap(b.CommandBuffer)

	// Create semaphore for concurrency limiting
	limiter := make(chan struct{}, limitSize)

	// Main event loop
	for event := range b.Gateway.Queue {
		// Get commands that match this event type
		for _, cmd := range commands[event.Type] {
			// Acquire limiter slot (blocks if at limit)
			limiter <- struct{}{}

			// Execute command in goroutine
			go func(c CommandTemplate, e *gateway2.RawEvent) {
				defer func() { <-limiter }() // Release slot when done
				b.InitCommand(c, e)
			}(cmd, event)
		}
	}
}

// DynamicEventProcessor reads b.CommandBuffer on every event.
// Allows adding/removing commands at runtime, but slightly slower.
// Use sync.RWMutex if you modify CommandBuffer concurrently.
func (b *Bot) ExampleBot_DynamicEventProcessor(limitSize uint) {
	// Create semaphore for concurrency limiting
	limiter := make(chan struct{}, limitSize)

	// Main event loop
	for event := range b.Gateway.Queue {
		// Collect matching commands from current buffer
		var activeCmds []CommandTemplate
		b.commandMu.Lock()
		for _, cmd := range b.CommandBuffer {
			if event.Type == cmd.Trigger {
				activeCmds = append(activeCmds, cmd)
			}
		}
		b.commandMu.Unlock()

		// Execute all matching commands
		for _, cmd := range activeCmds {
			// Acquire limiter slot (blocks if at limit)
			limiter <- struct{}{}

			// Execute command in goroutine
			go func(c CommandTemplate, e *gateway2.RawEvent) {
				defer func() { <-limiter }() // Release slot when done
				b.InitCommand(c, e)
			}(cmd, event)
		}
	}
}

func (b *Bot) ExampleBot_InitCommand(cmd CommandTemplate, event *gateway2.RawEvent) {
	// Recover from any panic in the command
	defer func() {
		if r := recover(); r != nil {
			if b.PanicHandler != nil {
				b.PanicHandler.OnPanic(event, &cmd, r, debug.Stack())
			}
		}
	}()

	// Execute the command
	if err := cmd.Action.Execute(event); err != nil {
		log.Printf("[EVENT PROCESSOR] %v", err)
	}
}

func prepareCommandMap(cmds []CommandTemplate) map[string][]CommandTemplate {
	cmdMap := make(map[string][]CommandTemplate)
	for _, cmd := range cmds {
		cmdMap[cmd.Trigger] = append(cmdMap[cmd.Trigger], cmd)
	}
	return cmdMap
}
