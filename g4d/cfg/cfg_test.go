package cfg_test

import (
	"testing"

	"github.com/IIIoooRRR/G4D/g4d/cfg"
)

func TestLoadCfg(t *testing.T) {
	bot, err := cfg.LoadConfig("config.exp.yaml")
	if err != nil {
		t.Errorf("Expected error: %v", err)
		return
	}
	t.Log(bot)
}
