package g4d

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"

	gw "github.com/IIIoooRRR/G4D/gateway"
)

type Bot struct {
	PanicHandler
	Token         string
	Gateway       *gw.Receiver // Don't you dare delete it.
	Prefix        string       // this is for beauty and so that you don't forget which prefix
	CommandBuffer []CommandTemplate
	appId         string
	CommandMu     sync.Mutex // you need this field if you want to rewrite the processor and work with commandBuffer in runtime, before the bot is fully initialized.
	Logger        *zap.Logger
	cmdLogger     *zap.Logger
	Client        *api.DiscordClient // I didn't make the client field private so that you wouldn't have to write crutches to change http.Client
	// without crutches. It may not be safe, but.. as it turned out. excuse me
	CtxTimeout time.Duration // if you need a timeout for commands, redefine this field.
}
type PanicHandler interface {
	OnPanic(event *gateway.RawEvent, cmd *CommandTemplate, r any, stack []byte)
}

func (b *Bot) Run() error {
	if b.CtxTimeout == 0 {
		b.CtxTimeout = time.Second * 10
	}
	if b.Client == nil {
		b.Logger.Panic("Discord http client not initialized. set bot.Client = api.NewClient(*bot.token, 10)")
	}
	if b.PanicHandler == nil {
		b.Logger.Panic("No panic handler. Initialize b.PanicHandler")
		return nil
	}
	name := b.Logger.Name()
	if name == "" || name == "root" { // дефолтное имя
		b.Logger = b.Logger.Named("bot")
	} else if !strings.Contains(name, "bot") {
		b.Logger = b.Logger.Named("bot")
	}
	b.cmdLogger = b.Logger.Named("command")
	err := b.Gateway.InitGateway(context.Background(), b.Logger.Named("gateway"), &b.Token)
	if err != nil {
		return err
	}
	return nil
}
