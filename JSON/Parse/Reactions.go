package Parse

import (
	"github.com/IIIoooRRR/G4D/JSON/Dependencies"
	Type2 "github.com/IIIoooRRR/G4D/JSON/Type"
)

type MessageReactionAdd struct {
	UserID          Type2.UserId              `json:"user_id"`
	ChannelID       Type2.ChannelId           `json:"channel_id"`
	MessageID       Type2.MessageId           `json:"message_id"`
	GuildID         Type2.GuildId             `json:"guild_id,omitempty"`
	Member          *Dependencies.GuildMember `json:"member,omitempty"`
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
	Emoji     Dependencies.Emoji `json:"emoji,omitempty"`
}
