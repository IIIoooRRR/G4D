package Event

import (
	"errors"

	Dependencies2 "github.com/IIIoooRRR/G4D/JSON/Dependencies"
)

type Message struct {
	ID              string                     `json:"id"`
	ChannelID       string                     `json:"channel_id"`
	GuildID         string                     `json:"guild_id,omitempty"` // Отсутствует в личке
	Author          Dependencies2.User         `json:"author"`
	Content         string                     `json:"content"`
	Flags           int                        `json:"flags,omitempty"`
	Timestamp       string                     `json:"timestamp"`
	EditedTimestamp string                     `json:"edited_timestamp,omitempty"`
	MentionEveryone bool                       `json:"mention_everyone"`
	Mentions        []Dependencies2.User       `json:"mentions"`
	MentionRoles    []string                   `json:"mention_roles"`
	Attachments     []Dependencies2.Attachment `json:"attachments"`
	Embeds          []Dependencies2.Embed      `json:"embeds"`
	Type            int                        `json:"type"`
	// ReferencedMessage нужна для обработки ответов (reply)
	ReferencedMessage *Message `json:"referenced_message,omitempty"`
}

type MessageDelete struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
}

func (m *Message) AddEmbed(embeds ...Dependencies2.Embed) error {
	if len(m.Embeds)+len(embeds) > 10 {
		return errors.New("[MESSAGE CREATE] Max 10 Embeds")
	}
	m.Embeds = append(m.Embeds, embeds...)
	return nil
}
