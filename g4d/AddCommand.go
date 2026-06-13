package g4d

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/IIIoooRRR/G4D/model/parse"
)

func (b *Bot) AddCommand(cmd CommandTemplate) {
	b.commandMu.Lock()
	b.CommandBuffer = append(b.CommandBuffer, cmd)
	b.commandMu.Unlock()
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
		log.Println("[PARSER] parse slash command error:", err)
		return err
	}

	url := fmt.Sprintf("https://discord.com/api/v10/applications/%s/commands", b.appId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "g4d "+b.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("[DISCORD INTERACTION] slash-command response:" + string(body))
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
