package shema

import (
	Type2 "github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

type MessageReactionAdd struct {
	UserID          Type2.UserId              `json:"user_id"`
	ChannelID       Type2.ChannelId           `json:"channel_id"`
	MessageID       Type2.MessageId           `json:"message_id"`
	GuildID         Type2.GuildId             `json:"guild_id,omitempty"`
	Member          *dependencies.GuildMember `json:"member,omitempty"`
	Emoji           PartialEmoji              `json:"emoji"`
	MessageAuthorID Type2.MessageId           `json:"message_author_id,omitempty"`
	Burst           bool                      `json:"burst"`
	BurstColors     []string                  `json:"burst_colors,omitempty"`
	Type            int                       `json:"type"`
}

type PartialEmoji struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Animated bool   `json:"animated,omitempty"`
}
type MessageReactionRemove struct {
	UserID    Type2.UserId       `json:"user_id"`
	ChannelID Type2.ChannelId    `json:"channel_id"`
	MessageID Type2.MessageId    `json:"message_id"`
	GuildID   Type2.GuildId      `json:"guild_id,omitempty"`
	Emoji     dependencies.Emoji `json:"emoji,omitempty"`
}
type MessageReactionRemoveAll struct {
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
	GuildID   string `json:"guild_id"`
}
type MessageReactionRemoveEmoji struct {
	ChannelID string             `json:"channel_id"`
	GuildID   string             `json:"guild_id"`
	Emoji     dependencies.Emoji `json:"emoji"`
}
