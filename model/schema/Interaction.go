package schema

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

type Interaction struct {
	ID            string            `json:"id"`
	ApplicationID string            `json:"application_id"`
	Type          int               `json:"type"` // 1-5
	Data          InteractionData   `json:"data"`
	GuildID       _const.GuildId    `json:"guild_id"`
	ChannelID     _const.ChannelId  `json:"channel_id"`
	Member        dependencies.User `json:"member"`
	Token         string            `json:"token"`
	Version       int               `json:"version"`

	// Для компонентов (кнопки/селекты)
	Message *GetMessage `json:"message,omitempty"`
}
type InteractionData struct {
	ID              int                 `json:"id,omitempty"`
	CustomID        string              `json:"custom_id,omitempty"`
	Type            int                 `json:"type,omitempty"`
	ComponentType   int                 `json:"component_type,omitempty"`
	SlashName       string              `json:"name,omitempty"`
	SlashOptions    []InteractionOption `json:"options,omitempty"`
	ComponentValues []string            `json:"values,omitempty"`
	Components      []Components        `json:"components,omitempty"`
	Focused         bool                `json:"focused,omitempty"`
}
type InteractionOption struct {
	Name  string      `json:"name"`
	Type  int         `json:"type"`
	Value interface{} `json:"value"`
}

// for modal submit
type Components struct {
	Type       int           `json:"type"`
	Components []interface{} `json:"components"`
}
