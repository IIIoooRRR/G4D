package g4d

import (
	"context"
	"strings"
	"sync"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"

	gw "github.com/IIIoooRRR/G4D/gateway"
)

type Bot struct {
	PanicHandler
	Token         string
	Gateway       *gw.Receiver
	Prefix        string
	CommandBuffer []CommandTemplate
	appId         string
	Context       context.Context
	CommandMu     sync.Mutex
	Logger        *zap.Logger
}
type PanicHandler interface {
	OnPanic(event *gateway.RawEvent, cmd *CommandTemplate, r any, stack []byte)
}

var bot *Bot
var botMu sync.RWMutex

func CurrentBot() *Bot {
	botMu.RLock()
	defer botMu.RUnlock()
	return bot
}
func (b *Bot) Run() error {
	botMu.Lock()
	bot = b
	botMu.Unlock()
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
	err := b.Gateway.CreateGateway(b.Context, b.Logger.Named("gateway"), &b.Token)
	if err != nil {
		return err
	}
	return nil
}
