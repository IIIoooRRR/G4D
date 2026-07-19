package g4d

import (
	"bytes"
	"io"
	"net/http"

	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
)

func (b *Bot) SetBotBio(description string) *Bot {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return b
}
