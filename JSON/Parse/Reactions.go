package Parse

import "github.com/IIIoooRRR/G4D/JSON/Dependencies"

type MessageReactionAdd struct {
	UserID          string                    `json:"user_id"`
	ChannelID       string                    `json:"channel_id"`
	MessageID       string                    `json:"message_id"`
	GuildID         string                    `json:"guild_id,omitempty"`
	Member          *Dependencies.GuildMember `json:"member,omitempty"`
	Emoji           PartialEmoji              `json:"emoji"`
	MessageAuthorID string                    `json:"message_author_id,omitempty"`
	Burst           bool                      `json:"burst"`
	BurstColors     []string                  `json:"burst_colors,omitempty"`
	Type            int                       `json:"type"`
}

type PartialEmoji struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Animated bool   `json:"animated,omitempty"`
}
