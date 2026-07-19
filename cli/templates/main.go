package templates

var Main = `
import (
	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/g4d/cfg"
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
	logger, err := zap.NewDevelopment()

	bot := cfg.LoadBot("g4d.yaml", logger, PanicHandler{})
	if err != nil {
		logger.Error("Error initializing logger", zap.Error(err))
	}
	go func() {
		bot.AddCommands([]g4d.CommandTemplate{
			{_const.EventMessageCreate, "", Button},
			{_const.EventInteractionCreate, "", ButtonReaction},
		}).InitProcessors(processor, 3, 34)
	}()
	if err := bot.Run(); err != nil {
		logger.Error("Error bot run", zap.Error(err))
	}
}

func Button(event *gateway.RawEvent, ctx *ctx.Context) error {
	d := parse.GetEvent[shema.GetMessage](event)
	if d.Content != "button" {
		return nil
	}

	button := ui.NewButton("button").SetLabel("button").SetStyle(1)
	row := ui.NewActionRow().AddComponents(button)
	msg := shema.NewMessage().AddActionRow(row).AddContent("buttooonn!!!!")

	return ctx.SendMessage(d.ChannelID, msg)
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
