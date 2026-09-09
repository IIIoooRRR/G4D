package g4d

import (
	"errors"
	"strings"
	"sync"

	"github.com/IIIoooRRR/G4D/api"
	gw "github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
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
	processorsOnce sync.Once
	eventCache     *parse.Cache
}
type PanicHandler interface {
	OnPanic(event *parse.RawEvent, cmd *CommandTemplate, r any, stack []byte)
}

func (b *Bot) Run(qnt _const.Quantity, limit _const.SemaphoreLimit, processorType _const.ProcessorType) error {
	if err := b.validate(); err != nil {
		return err
	}

	b.initLogger()
	b.initCache(qnt)
	b.initClient()

	if err := b.getBotInfo(); err != nil {
		return err
	}

	b.initProcessors(processorType, qnt, limit)
	return b.Gateway.InitGateway(b.Logger.Named("gateway"), &b.Token)
}

func (b *Bot) validate() error {
	if b.Logger == nil {
		return errors.New("logger is nil")
	}
	if b.Client == nil {
		return errors.New("http client not initialized")
	}
	if b.PanicHandler == nil {
		return errors.New("panic handler not initialized")
	}
	return nil
}

func (b *Bot) initClient() {
	b.Client.SetLogger(b.Logger.Named("http-client"))
}
func (b *Bot) initLogger() {
	name := b.Logger.Name()
	if !strings.Contains(name, "bot") {
		b.Logger = b.Logger.Named("bot")
	}
	b.cmdLogger = b.Logger.Named("command")
}
func (b *Bot) initCache(qnt _const.Quantity) {
	b.eventCache = parse.InitCache(qnt, b.Logger.Named("cache"))
}
