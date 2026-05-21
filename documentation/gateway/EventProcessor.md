## Static Event Processor
```go
func (b *Bot) StaticEventProcessor() {
    commands := sortCommand(b.CommandBuffer)
    for event := range b.Gateway.Queue {
        for _, command := range commands[event.Type] {
            go b.InitCommand(command, event)
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
    mu.Lock()
    bufCopy := b.CommandBuffer
    mu.Unlock()
	for _, command := range bufCopy {
        if event.Type != command.Trigger {
            continue
            }
        go b.InitCommand(command, event)
        }
    }
}
```
Use a dynamic event processor if you initialize commands at run time.\
It consumes more memory and may be slower, but each iteration causes a copy of the current array.
## Functions helper
```go
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

```