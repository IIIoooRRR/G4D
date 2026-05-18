package G4D

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (b *Bot) SetBotDescription(description string) error {
	url := "https://discord.com/api/v10/applications/@me"

	payload := map[string]string{"description": description}
	jsonBody, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+b.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
