package cfg

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/g4d"
	way "github.com/IIIoooRRR/G4D/gateway"
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
func loadConfig() (*Config, error) {
	cfg := defaultConfig()
	root := os.DirFS(".")
	data, err := fs.ReadFile(root, _const.ConfigPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func mustLoadCfg() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadBot(logger *zap.Logger, panicHandler g4d.PanicHandler) *g4d.Bot {
	cfg := mustLoadCfg()
	var activity []customize.Activity
	if cfg.GatewayConfig.PresenceUpdate.Activities != nil {
		activity = cfg.GatewayConfig.PresenceUpdate.Activities
	}

	token, err := resolveToken(cfg.BotConfig.Token)
	if err != nil {
		logger.Fatal("resolve token", zap.Error(err))
	}
	cfg.BotConfig.Token = token

	gateway := way.NewGateway(
		way.BufferSize(cfg.GatewayConfig.QueueSize),
		way.Intents(cfg.GatewayConfig.Intents),
		way.Activity(activity...),
		way.NetStatus(cfg.GatewayConfig.PresenceUpdate.Status),
	)
	bot := g4d.NewBot(
		cfg.BotConfig.Token,
		panicHandler,
		gateway,
		logger,
		api.NewClient(&cfg.BotConfig.Token, 10*time.Second),
	)
	bot.SetBotBio(cfg.BotConfig.Description)
	return bot
}

func resolveToken(raw string) (string, error) {
	if !strings.HasPrefix(raw, "${") || !strings.HasSuffix(raw, "}") {
		return raw, nil
	}
	envName := raw[2 : len(raw)-1]
	resolved := os.Getenv(envName)
	if resolved == "" {
		return "", fmt.Errorf("environment variable %q not found", envName)
	}
	return resolved, nil
}
