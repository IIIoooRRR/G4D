package templates

var Main = `
package main

import (
	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/g4d"
	gateway2 "github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
	"github.com/IIIoooRRR/G4D/model/dependencies/ui"
	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/parse"
	shema "github.com/IIIoooRRR/G4D/model/schema"
	"go.uber.org/zap"
)

var processor = g4d.StaticEventProcessor // or g4d.DynamicEventProcessor

func main() {
	var token string
	gateway := gateway2.NewGateway(10).WithIntents(34307).WithNetStatus(_const.NetStatusIDLE).WithDescription("hello!")
	logger := zap.Must(zap.NewProduction()).Named("bot")
	bot := &g4d.Bot{
		PanicHandler: PanicHandler{},
		Token:        token,
		Gateway:      gateway,
		Prefix:       "r.",
		Logger:       logger,
		Client:       api.NewClient(&token, 10, logger.Named("http client")),
	}
	go func() {
		bot.SetBotBio("test bot").AddCommands([]g4d.CommandTemplate{
			{_const.EventMessageCreate, "", Button},
			{_const.EventInteractionCreate, "", ButtonReaction},
		}).InitProcessors(processor, 2, 15)
	}()
	bot.Run()
}

func Button(event *gateway.RawEvent, ctx *ctx.Context) error {
	d := parse.GetEvent[shema.GetMessage](event)
	if d.Content != ctx.Prefix+"button" {
		return nil
	}
	defer ctx.Info("button cmd has finished") // ctx has api.Logger and api.DiscordClient

	button := ui.NewButton("button").SetLabel("button").SetStyle(1)
	row := ui.NewActionRow().AddComponents(button)
	msg := shema.NewMessage().AddActionRow(row).AddContent("buttooonn!!!!")
	return ctx.SendMessageWithLimit(d.ChannelID, msg)
}

func ButtonReaction(event *gateway.RawEvent, ctx *ctx.Context) error {
	d := parse.GetEvent[shema.Interaction](event)
	if d.Data.ComponentType != _const.ButtonType && d.Data.CustomID == "button" {
		return nil
	}

	data := shema.NewInteractionResponseData("hihi!")
	response := shema.NewInteractionResponse(_const.ResponseChannelMessageWithSource).SetData(*data)

	return ctx.SendInteractionMessage(&d, response)
}

type PanicHandler struct {
	g4d.PanicHandler
	Logger *zap.Logger
}

func (p PanicHandler) OnPanic(event *gateway.RawEvent, cmd *g4d.CommandTemplate, r any, stack []byte) {
	p.Logger.Panic("cmd:", zap.Any("cmd", cmd), zap.String("stack", string(stack)), zap.Any("r:", r))
}

`
