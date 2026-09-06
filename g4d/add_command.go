package g4d

import (
	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
)

func (b *Bot) AddCommand(cmd CommandTemplate) {
	b.CommandMu.Lock()
	b.CommandBuffer = append(b.CommandBuffer, cmd)
	b.CommandMu.Unlock()
}
func (b *Bot) AddCommands(cmds []CommandTemplate) *Bot {
	for _, cmd := range cmds {
		b.AddCommand(cmd)
	}
	return b
}
func (b *Bot) AddSlashCommand(cmd SlashCommandTemplate) error {

	jsonData, err := parse.Marshal(cmd.Form)
	if err != nil {

	}
	body, err := b.Client.DoDiscordRequest("POST", "/api/v10/applications/%s/commands", jsonData)
	if err != nil {
		return err
	}
	b.Logger.Info("add slash-command response:",
		zap.ByteString("body:", body))
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
