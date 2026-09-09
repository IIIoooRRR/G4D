package g4d

import (
	"context"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
)

func (b *Bot) AddCommand(cmd CommandTemplate) {
	b.CommandMu.Lock()
	b.CommandBuffer = append(b.CommandBuffer, cmd)
	b.CommandMu.Unlock()
}
func (b *Bot) AddCommands(cmds []CommandTemplate) *Bot {
	b.CommandMu.Lock()
	b.CommandBuffer = append(b.CommandBuffer, cmds...)
	b.CommandMu.Unlock()
	return b
}
func (b *Bot) AddSlashCommand(cmd SlashCommandTemplate) error {
	data, err := parse.Marshal(cmd.Form)
	if err != nil {
		return err
	}
	resp, err := b.Client.DoDiscordLimitRequest(context.Background(), "POST", api.GetURI("/applications/", b.appId, "/commands"), data)
	if err != nil {
		return err
	}
	b.Logger.Info("add slash-command response:",
		zap.ByteString("body:", resp))
	b.AddCommand(cmd.CommandTemplate)
	return nil
}

func (b *Bot) AddSlashCommands(cmds []SlashCommandTemplate) []error {
	var errs []error
	for _, cmd := range cmds {
		err := b.AddSlashCommand(cmd)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
