package cfg

import "github.com/IIIoooRRR/G4D/model/customize"

/*
documentation on the structure .The yaml file is located in the documentation/
this folder contains usage examples, bot configurations, and some files.
*/
type Config struct {
	BotConfig     BotConfig     `yaml:"bot"`
	GatewayConfig GatewayConfig `yaml:"gateway"`
}
type BotConfig struct {
	Token       string `yaml:"token"`
	Prefix      string `yaml:"prefix"`
	Description string `yaml:"description"`
}
type GatewayConfig struct {
	Intents        int                       `yaml:"intents"`
	QueueSize      int                       `yaml:"queue_size"`
	PresenceUpdate *customize.PresenceUpdate `yaml:"presence_update,omitempty"`
}
