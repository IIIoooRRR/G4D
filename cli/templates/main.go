package templates

var Main = `package main

import (
	"context"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/g4d/cfg"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/dependencies/ui"
	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/shema"
	"go.uber.org/zap"
)

var processor = g4d.StaticEventProcessor // or g4d.DynamicEventProcessor

func main() {
	cfg := cfg.MustLoadCfg("g4d.yaml")
	logger, err := zap.NewDevelopment()

	if err != nil {
		logger.Error("Error initializing logger", zap.Error(err))
	}

	bot, _ := cfg.NewBot(logger, context.Background(), PanicHandler{Logger: logger})

	go func() {
		bot.AddCommands([]g4d.CommandTemplate{
			{_const.EventMessageCreate, "", Button},
			{_const.EventInteractionCreate, "", ButtonReaction},
		}).InitProcessors(processor, 3, 34)
	}()

	if err != nil {
		logger.Error("Error initializing bot", zap.Error(err))
	}
	if err := bot.Run(); err != nil {
		logger.Error("Error bot run", zap.Error(err))
	}
}

func Button(event *gateway.RawEvent) error {
	d := parse.GetEvent[shema.GetMessage](event)
	if d.Content != "button" {
		return nil
	}

	button := ui.NewButton("button").SetLabel("button").SetStyle(1)
	row := ui.NewActionRow().AddComponents(button)
	msg := shema.NewMessage().AddActionRow(row).AddContent("buttooonn!!!!")

	return api.SendMessage(d.ChannelID, msg)
}

func ButtonReaction(event *gateway.RawEvent) error {
	d := parse.GetEvent[shema.Interaction](event)
	if d.Data.ComponentType != _const.ButtonType && d.Data.CustomID == "button" {
		return nil
	}

	data := shema.NewInteractionResponseData("hihi!")
	response := shema.NewInteractionResponse(_const.ResponseChannelMessageWithSource).SetData(*data)

	return api.SendInteractionMessage(&d, response)
}

type PanicHandler struct {
	Logger *zap.Logger
}

func (p PanicHandler) OnPanic(event *gateway.RawEvent, cmd *g4d.CommandTemplate, r any, stack []byte) {
	p.Logger.Panic("cmd:", zap.Any("cmd", cmd), zap.String("stack", string(stack)), zap.Any("r:", r))
}
`
