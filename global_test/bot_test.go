package global_test

import (
	"os"
	"testing"
	"time" // добавить

	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
	gateway2 "github.com/IIIoooRRR/G4D/model/gateway"
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

	gw := gateway.NewGateway(23).
		WithNetStatus(_const.NetStatusIDLE).
		WithIntents(34307)

	bot := &g4d.Bot{
		Token:   token,
		Gateway: gw,
		Logger:  zap.Must(zap.NewProduction()).Named("bot"),
	}

	bot.AddCommands([]g4d.CommandTemplate{
		{Trigger: _const.EventMessageCreate, Name: "hello", Execute: BotHello},
	})
	go bot.InitProcessors(g4d.StaticEventProcessor, 2, 25)
	go func() {
		err := bot.Run()
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(30 * time.Second)

	t.Log("Bot ran for 30 seconds, test passed")
}

func BotHello(event *gateway2.RawEvent, ctx *ctx.Context) error {
	d := parse.GetEvent[schema.GetMessage](event)
	if d.Content != "!hello" {
		return nil
	}
	return ctx.SendMessage(d.ChannelID, schema.NewMessage().AddContent("hello world"))
}
