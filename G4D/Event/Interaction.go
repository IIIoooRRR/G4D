package Event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IIIoooRRR/G4D/G4D"
	"github.com/IIIoooRRR/G4D/JSON/Dependencies"
)

type Interaction struct {
	ID            string            `json:"id"` // ID самого взаимодействия
	ApplicationID string            `json:"application_id"`
	Type          int               `json:"type"` // 2 для слэш-команд
	Data          InteractionData   `json:"data"` // данные команды
	GuildID       string            `json:"guild_id"`
	ChannelID     string            `json:"channel_id"`
	Member        Dependencies.User `json:"member"`
	Token         string            `json:"token"` // токен для ответа (в корне!)
	Version       int               `json:"version"`
}

type InteractionData struct {
	ID      string              `json:"id"`      // ID самой команды
	Name    string              `json:"name"`    // Имя команды (напр. "ping")
	Type    int                 `json:"type"`    // тип команды
	Options []InteractionOption `json:"options"` // параметры
}

type InteractionOption struct {
	Name  string      `json:"name"`
	Type  int         `json:"type"`
	Value interface{} `json:"value"`
}

type InteractionResponse struct {
	Type int                     `json:"type"`
	Data InteractionResponseData `json:"data"`
}
type InteractionResponseData struct {
	Content string               `json:"content"`
	Flags   int                  `json:"flags"`
	Embeds  []Dependencies.Embed `json:"embeds,omitempty"`
}

func (msg *InteractionResponse) SendInteractionMessage(event Interaction) error {
	url := fmt.Sprintf("https://discord.com/api/v10/interactions/%s/%s/callback", event.ID, event.Token)
	jsonData, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Println("[SEND-MESSAGE] Discord API ERROR " + resp.Status)
	}
	return nil
}
