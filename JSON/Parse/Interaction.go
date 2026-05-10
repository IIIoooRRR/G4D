package Parse

import (
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
