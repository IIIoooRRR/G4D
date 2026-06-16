## Static Event Processor
```go

func (b *Bot) StaticEventProcessor(limitSize int) {
logger := b.Logger.Named("eventProcessor")
limiter := make(chan struct{}, limitSize)
commands := prepareCommand(b.CommandBuffer)
    for event := range b.Gateway.Queue {
        for _, command := range commands[event.Type] {
            limiter <- struct{}{}
            go func(command CommandTemplate, event *gateway.RawEvent) {
                defer func() { <-limiter }()
                b.InitCommand(command, event, logger)
            }(command, event)
        }
    }
}

```
Use a static event handler if you do not have commands that will be initialized while the bot is running under certain conditions.\
It does not check new commands again after launching the bot.CreateBot(...) consumes less RAM
## Dynamic Event Processor
```go
func (b *Bot) DynamicEventProcessor(limitSize int) {
	logger := b.Logger.Named("eventProcessor")
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
				b.InitCommand(localCmd, event, logger)
			}(command, event)
		}
	}
}
```
Use a dynamic event processor if you initialize commands at run time.\
It consumes more memory and may be slower, but each iteration causes a copy of the current array.
## Functions helper
```go

func prepareCommand(cmds []CommandTemplate) map[string][]CommandTemplate {
CmdMap := make(map[string][]CommandTemplate)
    for _, command := range cmds {
        CmdMap[command.Trigger] = append(CmdMap[command.Trigger], command)
    }
return CmdMap
}

func (b *Bot) InitCommand(command CommandTemplate, event *gateway.RawEvent, logger *zap.Logger) {
defer func() {
    if r := recover(); r != nil {
        if b.PanicHandler != nil {
            b.PanicHandler.OnPanic(event, &command, r, debug.Stack())
        }
    }
}()
err := command.Action.Execute(event)
if err != nil {
    logger.Error("cmd execution error ", zap.Error(err))
    }
}

```