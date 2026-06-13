package shema

import (
	"log"

	Type2 "github.com/IIIoooRRR/G4D/model/_const"
	Dependencies2 "github.com/IIIoooRRR/G4D/model/dependencies"
	"github.com/IIIoooRRR/G4D/model/dependencies/ui"
)

type GetMessage struct {
	ID              Type2.MessageId            `json:"id,omitempty"`
	GuildID         Type2.GuildId              `json:"guild_id,omitempty"` // Отсутствует в личке
	ChannelID       Type2.ChannelId            `json:"channel_id,omitempty"`
	Thread          *Channel                   `json:"thread,omitempty"`
	Author          Dependencies2.User         `json:"author,omitempty"`
	Member          Dependencies2.GuildMember  `json:"member,omitempty"`
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
	// ReferencedMessage - reply
	//put the shema result here.ToMessage(), if you use the MessageCreate type
	ReferencedMessage *MessageReference `json:"message_reference,omitempty"`
}
type SendMessage struct {
	Content     string                     `json:"content"`
	Flags       int                        `json:"flags,omitempty"`
	Attachments []Dependencies2.Attachment `json:"attachments,omitempty"`
	Embeds      []Dependencies2.Embed      `json:"embeds,omitempty"`
	Type        int                        `json:"type,omitempty"`
	Components  []ui.ActionRow             `json:"components,omitempty"`
	// ReferencedMessage - replu
	//put the shema result here.ToMessage(), if you use the MessageCreate type
	ReferencedMessage *MessageReference `json:"message_reference,omitempty"`
}

type MessageDelete struct {
	ID        Type2.MessageId `json:"id"`
	ChannelID Type2.ChannelId `json:"channel_id"`
	GuildID   Type2.GuildId   `json:"guild_id,omitempty"`
}

type MessageEdit struct {
	Content         string                `json:"content"`
	Embeds          []Dependencies2.Embed `json:"embeds,omitempty"`
	Flags           []string              `json:"flags,omitempty"`
	AllowedMentions string                `json:"allowed_mentions,omitempty"`
	Components      []string              `json:"components,omitempty"`
	Files           []string              `json:"files,omitempty"`
	PayloadJson     string                `json:"payload_json,omitempty"`
	Attachments     []string              `json:"attachments,omitempty"`
}

type MessageReference struct {
	ChannelId Type2.ChannelId `json:"channel_id,omitempty"`
	MessageId Type2.MessageId `json:"message_id,omitempty"`
}

func NewMessage() *SendMessage {
	return &SendMessage{}
}
func (m *SendMessage) AddContent(content string) *SendMessage {
	m.Content = content
	return m
}
func (m *SendMessage) AddFlags(flags int) *SendMessage {
	m.Flags = flags
	return m
}
func (m *SendMessage) AddEmbed(embeds ...Dependencies2.Embed) *SendMessage {
	if len(m.Embeds)+len(embeds) <= 10 {
		m.Embeds = append(m.Embeds, embeds...)
	} else {
		log.Println("[MESSAGE CREATE] max 10 embeds")
	}
	return m
}
func (m *SendMessage) AddAttachment(attachments ...Dependencies2.Attachment) *SendMessage {
	if len(m.Attachments)+len(attachments) <= 10 {
		m.Attachments = append(m.Attachments, attachments...)
	} else {
		log.Println("[MESSAGE CREATE]  max 10 attachments")
	}
	return m

}
func (m *SendMessage) AddReferencedMessage(message *GetMessage) *SendMessage {
	m.ReferencedMessage = &MessageReference{
		ChannelId: message.ChannelID,
		MessageId: message.ID,
	}
	return m
}
func (m *SendMessage) AddReference(channelId Type2.ChannelId, messageId Type2.MessageId) *SendMessage {
	m.ReferencedMessage = &MessageReference{
		ChannelId: channelId,
		MessageId: messageId,
	}
	return m
}
func (m *SendMessage) AddActionRow(actionRow ...ui.ActionRow) *SendMessage {
	m.Components = append(m.Components, actionRow...)
	return m
}
