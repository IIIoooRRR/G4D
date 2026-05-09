## Static Event Processor
```go
    func (b *Bot) StaticEventProcessor() {
        Buffer := &b.CommandBuffer
for event := range b.Gateway.Queue {
    for _, command := range *Buffer {
        log.Println(event.Type, command.Trigger)
	    if event.Type != command.Trigger {
             continue
        }
        go func(cmd *Command, event *Connect.RawEvent) {
        err := cmd.Action.Execute(b, event)
        if err != nil {
        log.Println("[EVENT PROCESSOR] ", err)
		}
      }(&command, event)
	}
  }
}
```
Use a static event handler if you do not have commands that will be initialized while the bot is running under certain conditions.\
It does not check new commands again after launching the bot.CreateBot(...) consumes less RAM
## Dynamic Event Processor
```go
func (b *Bot) DynamicEventProcessor() {

	for event := range b.Gateway.Queue {
		for _, command := range b.CommandBuffer {
			log.Println(event.Type, command.Trigger)
			if event.Type != command.Trigger {
				continue
			}
			go func(cmd *Command, event *Connect.RawEvent) {
				err := cmd.Action.Execute(b, event)
				if err != nil {
					log.Println("[EVENT PROCESSOR] ", err)
				}
			}(&command, event)
		}
	}
}
```
Use a dynamic event processor if you initialize commands at run time.\
It consumes more memory and may be slower, but each iteration causes a copy of the current array.