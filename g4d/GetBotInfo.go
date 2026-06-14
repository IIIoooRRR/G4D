package g4d

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IIIoooRRR/G4D/model/parse"
)

func (b *Bot) GetBotInfo() error {
	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		log.Println(fmt.Errorf("error creating request: %v", err))
		return err
	}
	req.Header.Set("Authorization", "g4d "+b.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req) //отправка инфо
	if err != nil {
		log.Println("[FATAL ERROR] Error getting bot info: ", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("[FATAL ERROR] Error closing body: ", err)
		}
	}(resp.Body)

	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[FATAL ERROR] Error reading body: ", err)
	}
	var body map[string]interface{}
	err = parse.Unmarshal(byteValue, &body)
	if err != nil {
		log.Println("[FATAL ERROR] Error parsing body: ", err)
	}
	log.Println("[DISCORD] gateway is correct") // получаем id бота
	id, ok := body["id"].(string)
	if !ok {
		return errors.New("[BOT CONNECT] Pasre failed")
	}
	b.appId = id
	return nil
}
