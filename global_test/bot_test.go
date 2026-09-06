package global_test

import (
	"os"
	"testing"
	"time"

	"github.com/IIIoooRRR/G4D/g4d"
	way "github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/schema"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func TestBotCreate(t *testing.T) {
	err := godotenv.Load(".env.ci")
	if err != nil {
		t.Log(".env.ci not found")
	}

	token := os.Getenv("CI_TOKEN")
	if token == "" {
		t.Skip("CI_TOKEN not set, skipping integration test")
	}

	gw := way.NewGateway(way.BufferSize(15),
		way.Intents(34307),
		way.Activity(customize.Activity{
			Name: "With G4D",
			Type: _const.ActivityStreaming,
		}),
		way.NetStatus(_const.NetStatusIDLE)).WithDescription("hello!")

	bot := &g4d.Bot{
		Token:   token,
		Gateway: gw,
		Logger:  zap.Must(zap.NewProduction()).Named("bot"),
	}

	bot.AddCommands([]g4d.CommandTemplate{
		{Trigger: _const.EventMessageCreate, Execute: BotHello},
	})
	go func() {
		err := bot.Run(g4d.WithDispQuantity(4), g4d.WithSemaphoreLimit(120), _const.StaticEventProcessor)
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(30 * time.Second)

	t.Log("Bot ran for 30 seconds, test passed")
}

func BotHello(event *parse.RawEvent, ctx *ctx.Context) error {
	d := parse.GetEvent[schema.GetMessage](event)
	if d.Content != "!hello" {
		return nil
	}
	return ctx.SendMessage(d.ChannelID, schema.NewMessage().AddContent("hello world"))
}
