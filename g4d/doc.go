// Custom processors
//
// G4D ships with two built-in processors:
//   - StaticEventProcessor: one goroutine per processor, waits for all commands
//   - DynamicEventProcessor: spawns a goroutine per command
//
// To add a new processor type, use RegisterProcessor before Bot.Run:
//
//	bot.RegisterProcessor("my-processor", func(b *g4d.Bot, seq int, limiter chan struct{}) {
//	    for {
//	        select {
//	        case <-b.Context().Done():
//	            return
//	        case event := <-b.Gateway.Queue:
//	            // your logic
//	        }
//	    }
//	})
//	bot.Run(..., g4d.ProcessorType("my-processor"), ...)
//
// The processor func receives *g4d.Bot as the first argument, so it
// has full access to bot fields (Gateway, Logger, etc.).
//
// Built-in processors (StaticEventProcessor, DynamicEventProcessor)
// cannot be replaced. RegisterProcessor returns an error if the type
// is already registered.
//
// The limiter channel must be respected: send before spawning a
// goroutine, receive when it finishes.
package g4d
