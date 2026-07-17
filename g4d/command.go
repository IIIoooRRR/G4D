package g4d

import (
	"runtime/debug"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"
)

type CommandTemplate struct {
	Trigger string
	Name    string
	Action  ToCommand
}

type SlashCommandTemplate struct {
	Form            SlashCreateCommand
	CommandTemplate CommandTemplate
}

func (b *Bot) initCommand(command CommandTemplate, event *gateway.RawEvent, logger *zap.Logger) {
	defer func() {
		if r := recover(); r != nil {
			b.OnPanic(event, &command, r, debug.Stack())
		}
	}()
	err := command.Action(event)
	if err != nil {
		logger.Error("cmd execution error ", zap.Error(err))
	}
}
