package g4d

import (
	"context"
	"sync"

	"github.com/IIIoooRRR/G4D/gateway"
)

type Bot struct {
	Token         string
	Gateway       *gateway.Receiver
	Prefix        string
	CommandBuffer []CommandTemplate
	appId         string
	Context       context.Context
	Cache         gateway.Cacher
	PanicHandler  PanicHandler
	commandMu     sync.Mutex
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
	b.Gateway.Cache = b.Cache
	botMu.Lock()
	bot = b
	botMu.Unlock()
	b.Gateway.CreateBot(b.Context, &b.Token)
	return nil
}
