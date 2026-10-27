package schema

import (
	_const "github.com/IIIoooRRR/G4D/model/_const"
	Dependencies2 "github.com/IIIoooRRR/G4D/model/entity"
	"github.com/IIIoooRRR/G4D/model/entity/ui"
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
type Message struct {
	Content     string                     `json:"content"`
	Flags       int                        `json:"flags,omitempty"`
	Attachments []Dependencies2.Attachment `json:"attachments,omitempty"`
	Embeds      []Dependencies2.Embed      `json:"embeds,omitempty"`
	Type        int                        `json:"type,omitempty"`
	Components  []ui.Component             `json:"components,omitempty"`
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
}

type MessageReference struct {
	ChannelId _const.ChannelId `json:"channel_id,omitempty"`
	MessageId _const.MessageId `json:"message_id,omitempty"`
}

func NewMessage() *Message {
	return &Message{}
}
func (m *Message) AddContent(content string) *Message {
	m.Content = content
	return m
}
func (m *Message) AddFlags(flags int) *Message {
	m.Flags = flags
	return m
}
func (m *Message) AddEmbed(embeds ...Dependencies2.Embed) *Message {
	m.Embeds = append(m.Embeds, embeds...)
	return m
}
func (m *Message) AddReferencedMessage(message *GetMessage) *Message {
	m.ReferencedMessage = &MessageReference{
		ChannelId: message.ChannelID,
		MessageId: message.ID,
	}
	return m
}
func (m *Message) AddReference(channelId _const.ChannelId, messageId _const.MessageId) *Message {
	m.ReferencedMessage = &MessageReference{
		ChannelId: channelId,
		MessageId: messageId,
	}
	return m
}
func (m *Message) AddComponents(cmps ...ui.Component) *Message {
	m.Components = append(m.Components, cmps...)
	return m
}
