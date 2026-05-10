package Parse

import (
	"github.com/IIIoooRRR/G4D/JSON/Dependencies"
)

type GuildMemberAdd struct {
	GuildID      string            `json:"guild_id"`
	User         Dependencies.User `json:"user"`
	Roles        []string          `json:"roles"`
	JoinedAt     string            `json:"joined_at"`
	Nick         string            `json:"nick,omitempty"`
	Avatar       string            `json:"avatar,omitempty"`
	Pending      bool              `json:"pending,omitempty"` // Прошел ли верификацию
	PremiumSince string            `json:"premium_since,omitempty"`
}
type GuildMemberRemove struct {
	GuildID string            `json:"guild_id"`
	User    Dependencies.User `json:"user"`
}

type Guild struct {
	ID                          string                     `json:"id"`
	Name                        string                     `json:"name"`
	Icon                        string                     `json:"icon"`
	OwnerID                     string                     `json:"owner_id"`
	Permissions                 string                     `json:"permissions"`
	AfkChannelID                string                     `json:"afk_channel_id"`
	AfkTimeout                  int                        `json:"afk_timeout"`
	VerificationLevel           int                        `json:"verification_level"`
	DefaultMessageNotifications int                        `json:"default_message_notifications"`
	ExplicitContentFilter       int                        `json:"explicit_content_filter"`
	Roles                       []Dependencies.Role        `json:"roles"`
	Emojis                      []Dependencies.Emoji       `json:"emojis"`
	Features                    []string                   `json:"features"`
	MFALevel                    int                        `json:"mfa_level"`
	SystemChannelID             string                     `json:"system_channel_id"`
	SystemChannelFlags          int                        `json:"system_channel_flags"`
	RulesChannelID              string                     `json:"rules_channel_id"`
	VanityURLCode               string                     `json:"vanity_url_code"`
	Description                 string                     `json:"description"`
	Banner                      string                     `json:"banner"`
	PremiumTier                 int                        `json:"premium_tier"`
	PreferredLocale             string                     `json:"preferred_locale"`
	PublicUpdatesChannelID      string                     `json:"public_updates_channel_id"`
	NSFWLevel                   int                        `json:"nsfw_level"`
	PremiumProgressBarEnabled   bool                       `json:"premium_progress_bar_enabled"`
	Channels                    []Channel                  `json:"channels"`
	Members                     []Dependencies.GuildMember `json:"members"`
	Presences                   []interface{}              `json:"presences"`
	VoiceStates                 []Dependencies.VoiceState  `json:"voice_states"`
	Threads                     []Channel                  `json:"threads"`
}
