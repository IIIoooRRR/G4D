package g4d

import (
	"bytes"
	"net/http"

	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
)

func (b *Bot) SetBotDescription(description string) *Bot {
	url := "https://discord.com/api/v10/applications/@me"

	payload := map[string]string{"description": description}
	jsonBody, _ := parse.Marshal(payload)

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+b.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		b.Logger.Error("set description err:", zap.Error(err))
		return nil
	}
	defer resp.Body.Close()

	return b
}
