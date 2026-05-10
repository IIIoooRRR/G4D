package Parse

import (
	"errors"

	Dependencies2 "github.com/IIIoooRRR/G4D/JSON/Dependencies"
)

type Message struct {
	ID              string                     `json:"id,omitempty"`
	ChannelID       string                     `json:"channel_id,omitempty"`
	GuildID         string                     `json:"guild_id,omitempty"` // Отсутствует в личке
	Author          Dependencies2.User         `json:"author,omitempty"`
	Content         string                     `json:"content,omitempty"`
	Flags           int                        `json:"flags,omitempty"`
	Timestamp       string                     `json:"timestamp,omitempty"`
	EditedTimestamp string                     `json:"edited_timestamp,omitempty"`
	MentionEveryone bool                       `json:"mention_everyone,omitempty"`
	Mentions        []Dependencies2.User       `json:"mentions,omitempty"`
	MentionRoles    []string                   `json:"mention_roles,omitempty"`
	Attachments     []Dependencies2.Attachment `json:"attachments,omitempty"`
	Embeds          []Dependencies2.Embed      `json:"embeds,omitempty"`
	Type            int                        `json:"type,omitempty"`
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
