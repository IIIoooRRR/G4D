package schema

import (
	"encoding/json"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
	"github.com/IIIoooRRR/G4D/model/entity/ui"
)

type GetMessage struct {
	ID              _const.MessageId    `json:"id,omitempty"`
	GuildID         _const.GuildId      `json:"guild_id,omitempty"` // Отсутствует в личке
	ChannelID       _const.ChannelId    `json:"channel_id,omitempty"`
	Thread          *Channel            `json:"thread,omitempty"`
	Author          entity.User         `json:"author,omitempty"`
	Member          entity.GuildMember  `json:"member,omitempty"`
	Content         string              `json:"content,omitempty"`
	Flags           int                 `json:"flags,omitempty"`
	Timestamp       string              `json:"timestamp,omitempty"`
	EditedTimestamp string              `json:"edited_timestamp,omitempty"`
	MentionEveryone bool                `json:"mention_everyone,omitempty"`
	Mentions        []entity.User       `json:"mentions,omitempty"`
	MentionRoles    []string            `json:"mention_roles,omitempty"`
	Attachments     []entity.Attachment `json:"attachments,omitempty"`
	Type            int                 `json:"type,omitempty"`
	// ReferencedMessage - reply
	//put the schema result here.ToMessage(), if you use the MessageCreate type
	ReferencedMessage *MessageReference `json:"message_reference,omitempty"`
	Components        json.RawMessage   `json:"components,omitempty"`
}
type Message struct {
	Cnt         string               `json:"content"`
	Flgs        int                  `json:"flags,omitempty"`
	Attachments []entity.Attachment  `json:"attachments,omitempty"`
	Embeds      []entity.Embed       `json:"embeds,omitempty"`
	Type        int                  `json:"type,omitempty"`
	Cmpnts      []ui.TopLvlComponent `json:"components,omitempty"`
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
	Content         string            `json:"content,omitempty"`
	Embeds          []entity.Embed    `json:"embeds,omitempty"`
	Flags           int               `json:"flags,omitempty"`
	AllowedMentions string            `json:"allowed_mentions,omitempty"`
	Components      []json.RawMessage `json:"components,omitempty"`
	Files           []string          `json:"files,omitempty"`
	PayloadJson     string            `json:"payload_json,omitempty"`
}

type MessageReference struct {
	ChannelId _const.ChannelId `json:"channel_id,omitempty"`
	MessageId _const.MessageId `json:"message_id,omitempty"`
}

func NewMessage() *Message {
	return &Message{}
}
func (m *Message) Content(content string) *Message {
	m.Cnt = content
	return m
}
func (m *Message) Flags(flags int) *Message {
	m.Flgs = flags
	return m
}
func (m *Message) AddEmbed(embeds ...entity.Embed) *Message {
	m.Embeds = append(m.Embeds, embeds...)
	return m
}
func (m *Message) Reply(message *GetMessage) *Message {
	m.ReferencedMessage = &MessageReference{
		ChannelId: message.ChannelID,
		MessageId: message.ID,
	}
	return m
}
func (m *Message) CustomReply(channelId _const.ChannelId, messageId _const.MessageId) *Message {
	m.ReferencedMessage = &MessageReference{
		ChannelId: channelId,
		MessageId: messageId,
	}
	return m
}
func (m *Message) Components(cmps ...ui.TopLvlComponent) *Message {
	m.Cmpnts = append(m.Cmpnts, cmps...)
	return m
}
