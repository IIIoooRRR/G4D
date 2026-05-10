package G4D

import (
	"context"
	"sync"

	"github.com/IIIoooRRR/G4D/Connect"
	"github.com/IIIoooRRR/G4D/JSON/Dependencies"
	"github.com/IIIoooRRR/G4D/JSON/Parse"
)

type Bot struct {
	Token         string
	Gateway       *Connect.Receiver
	Prefix        string
	CommandBuffer []Command
	appId         string
	Context       context.Context
	Cache         Cacher
}

type Cacher interface {
	GetUser(Id string) (Dependencies.User, error)
	GetGuild(Id string) (Parse.Guild, error)
	GetMessage(Id string) (Parse.Message, error)
	GetInteraction(Id string) (Parse.Interaction, error)
	GetChannel(Id string) (Parse.Channel, error)
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
	err := b.Gateway.CreateBot(b.Context, &b.Token)
	if err != nil {
		return err
	}
	return nil
}
