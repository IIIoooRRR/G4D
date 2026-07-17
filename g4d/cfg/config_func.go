package cfg

import (
	"context"
	"io/fs"
	"os"

	"github.com/IIIoooRRR/G4D/g4d"
	gateway2 "github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/customize"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func defaultConfig() *Config {
	return &Config{
		BotConfig: BotConfig{
			Token:       "",
			Prefix:      "",
			Description: "",
		},
		GatewayConfig: GatewayConfig{
			Intents: 34307,
			PresenceUpdate: &customize.PresenceUpdate{
				Activities: nil,
				Status:     _const.NetStatusOnline,
			}},
	}
}
func LoadConfig(path string) (*Config, error) {
	cfg := defaultConfig()
	root := os.DirFS(".")
	data, err := fs.ReadFile(root, path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func MustLoadCfg(path string) *Config {
	cfg, err := LoadConfig(path)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (cfg *Config) NewBot(logger *zap.Logger, ctx context.Context, panicHandler g4d.PanicHandler) (*g4d.Bot, error) {
	gateway := gateway2.NewGateway(cfg.GatewayConfig.QueueSize).
		WithNetStatus(cfg.GatewayConfig.PresenceUpdate.Status).
		WithIntents(cfg.GatewayConfig.Intents)
	if cfg.GatewayConfig.PresenceUpdate.Activities != nil {
		gateway = gateway.WithActivity(cfg.GatewayConfig.PresenceUpdate.Activities...)
	}
	bot := &g4d.Bot{
		Token:        cfg.BotConfig.Token,
		Gateway:      gateway,
		Prefix:       cfg.BotConfig.Prefix,
		Logger:       logger,
		Context:      ctx,
		PanicHandler: panicHandler,
	}
	bot.SetBotDescription(cfg.BotConfig.Description)
	return bot, nil
}
