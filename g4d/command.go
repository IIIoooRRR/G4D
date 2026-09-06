package g4d

import (
	"runtime/debug"

	"github.com/IIIoooRRR/G4D/model/ctx"
	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
)

type Execute func(event *parse.RawEvent, ctx *ctx.Context) error

type CommandTemplate struct {
	Trigger string
	Name    string
	Execute
}

type SlashCommandTemplate struct {
	CommandTemplate
	Form SlashCreateCommand
}

func (b *Bot) initCommand(command CommandTemplate, event *parse.RawEvent, ctx *ctx.Context) {
	defer func() {
		if r := recover(); r != nil {
			b.OnPanic(event, &command, r, debug.Stack())
		}
	}()
	err := command.Execute(event, ctx)
	if err != nil {
		b.Logger.Error("cmd execution error ", zap.Error(err))
	}
}
