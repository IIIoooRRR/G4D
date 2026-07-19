package schema

import (
	_const "github.com/IIIoooRRR/G4D/model/_const"
	Dependencies2 "github.com/IIIoooRRR/G4D/model/dependencies"
	"github.com/IIIoooRRR/G4D/model/dependencies/ui"
)

type GetMessage struct {
	ID              _const.MessageId           `json:"id,omitempty"`
	GuildID         _const.GuildId             `json:"guild_id,omitempty"` // Отсутствует в личке
	ChannelID       _const.ChannelId           `json:"channel_id,omitempty"`
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
	//put the schema result here.ToMessage(), if you use the MessageCreate type
	ReferencedMessage *MessageReference `json:"message_reference,omitempty"`
}
type SendMessage struct {
	Content     string                     `json:"content"`
	Flags       int                        `json:"flags,omitempty"`
	Attachments []Dependencies2.Attachment `json:"attachments,omitempty"`
	Embeds      []Dependencies2.Embed      `json:"embeds,omitempty"`
	Type        int                        `json:"type,omitempty"`
	Components  []ui.ActionRow             `json:"components,omitempty"`
	// ReferencedMessage - reply
	//put the schema result here.ToMessage(), if you use the MessageCreate type
	ReferencedMessage *MessageReference `json:"message_reference,omitempty"`
}

type MessageDelete struct {
	ID        _const.MessageId `json:"id"`
	ChannelID _const.ChannelId `json:"channel_id"`
	GuildID   _const.GuildId   `json:"guild_id,omitempty"`
}
type MessageDeleteBulk struct {
	IDs       []_const.MessageId `json:"ids"`
	ChannelId _const.ChannelId   `json:"channel_id"`
	GuildId   _const.GuildId     `json:"guild_id,omitempty"`
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
	ChannelId _const.ChannelId `json:"channel_id,omitempty"`
	MessageId _const.MessageId `json:"message_id,omitempty"`
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
	m.Embeds = append(m.Embeds, embeds...)
	return m
}
func (m *SendMessage) AddAttachment(attachments ...Dependencies2.Attachment) *SendMessage {
	m.Attachments = append(m.Attachments, attachments...)
	return m

}
func (m *SendMessage) AddReferencedMessage(message *GetMessage) *SendMessage {
	m.ReferencedMessage = &MessageReference{
		ChannelId: message.ChannelID,
		MessageId: message.ID,
	}
	return m
}
func (m *SendMessage) AddReference(channelId _const.ChannelId, messageId _const.MessageId) *SendMessage {
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
