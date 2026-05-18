package G4D

import (
	"context"
	"sync"

	"github.com/IIIoooRRR/G4D/connect"
)

type Bot struct {
	Token         string
	Gateway       *connect.Receiver
	Prefix        string
	CommandBuffer []Command
	appId         string
	Context       context.Context
	Cache         connect.Cacher
	PanicHandler  PanicHandler
}
type PanicHandler interface {
	OnPanic(event *connect.RawEvent, cmd *Command, r any, stack []byte)
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
