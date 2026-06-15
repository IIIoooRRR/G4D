package g4d

import (
	"errors"
	"io"
	"net/http"

	"github.com/IIIoooRRR/G4D/model/parse"
	"go.uber.org/zap"
)

func (b *Bot) GetBotInfo() error {
	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		b.Logger.Error("error creating request: ", zap.Error(err))
		return err
	}
	req.Header.Set("Authorization", "Bot "+b.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req) //отправка инфо
	if err != nil {
		b.Logger.Error("getting bot info: ", zap.Error(err))
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			b.Logger.Error("Error closing body: ", zap.Error(err))
		}
	}()

	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		b.Logger.Error("Error reading body: ", zap.Error(err))
	}
	var body map[string]interface{}
	err = parse.Unmarshal(byteValue, &body)
	if err != nil {
		b.Logger.Error("Error parsing body: ", zap.Error(err))
	}
	b.Logger.Info("gateway is correct") // получаем id бота
	id, ok := body["id"].(string)
	if !ok {
		return errors.New("parse failed")
	}
	b.appId = id
	return nil
}
