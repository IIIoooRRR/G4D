package shema

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

type GuildMemberAdd struct {
	GuildID      _const.GuildId    `json:"guild_id"`
	User         dependencies.User `json:"user"`
	Roles        []string          `json:"roles"`
	JoinedAt     string            `json:"joined_at"`
	Nick         string            `json:"nick,omitempty"`
	Avatar       string            `json:"avatar,omitempty"`
	Pending      bool              `json:"pending,omitempty"` // Прошел ли верификацию
	PremiumSince string            `json:"premium_since,omitempty"`
}
type GuildMemberRemove struct {
	GuildID _const.GuildId    `json:"guild_id"`
	User    dependencies.User `json:"user"`
}

type Guild struct {
	ID                          _const.GuildId       `json:"id"`
	Name                        string               `json:"name"`
	Icon                        string               `json:"icon"`
	IconHash                    string               `json:"icon_hash,omitempty"`
	Splash                      string               `json:"splash,omitempty"`
	DiscoverySplash             string               `json:"discovery_splash,omitempty"`
	OwnerID                     string               `json:"owner_id"`
	Permissions                 string               `json:"permissions,omitempty"`
	Region                      string               `json:"region,omitempty"`
	WidgetEnabled               bool                 `json:"widget_enabled,omitempty"`
	WidgetChannelID             string               `json:"widget_channel_id,omitempty"`
	AfkChannelID                string               `json:"afk_channel_id"`
	AfkTimeout                  int                  `json:"afk_timeout"`
	VerificationLevel           int                  `json:"verification_level"`
	DefaultMessageNotifications int                  `json:"default_message_notifications"`
	ExplicitContentFilter       int                  `json:"explicit_content_filter"`
	MFALevel                    int                  `json:"mfa_level"`
	NSFWLevel                   int                  `json:"nsfw_level"`
	Roles                       []dependencies.Role  `json:"roles"`
	Emojis                      []dependencies.Emoji `json:"emojis"`
	Features                    []string             `json:"features"`
	Banner                      string               `json:"banner,omitempty"`
	Description                 string               `json:"description,omitempty"`
	VanityURLCode               interface{}          `json:"vanity_url_code,omitempty"`
	PremiumTier                 int                  `json:"premium_tier"`
	PremiumProgressBarEnabled   bool                 `json:"premium_progress_bar_enabled,omitempty"`
	PremiumSubscriptionCount    int                  `json:"premium_subscription_count,omitempty"`
	SystemChannelID             string               `json:"system_channel_id"`
	SystemChannelFlags          int                  `json:"system_channel_flags"`
	RulesChannelID              string               `json:"rules_channel_id"`
	PublicUpdatesChannelID      string               `json:"public_updates_channel_id"`
	PreferredLocale             string               `json:"preferred_locale"`
	ApplicationID               interface{}          `json:"application_id,omitempty"`
	MaxPresences                interface{}          `json:"max_presences,omitempty"`
	MaxMembers                  int                  `json:"max_members,omitempty"`
	MaxVideoChannelUsers        int                  `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   int                  `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      int                  `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    int                  `json:"approximate_presence_count,omitempty"`

	Channels    []Channel                  `json:"channels,omitempty"`
	Members     []dependencies.GuildMember `json:"members,omitempty"`
	Presences   []interface{}              `json:"presences,omitempty"`
	VoiceStates []dependencies.VoiceState  `json:"voice_states,omitempty"`
	Threads     []Channel                  `json:"threads,omitempty"`
}
type GuildBanAdd struct {
	GuildId _const.GuildId    `json:"guild_id"`
	User    dependencies.User `json:"user"`
}
type GuildBanRemove struct {
	GuildId _const.GuildId    `json:"guild_id"`
	User    dependencies.User `json:"user"`
}
type GuildEmojisUpdate struct {
	GuildId _const.GuildId       `json:"guild_id"`
	Emoji   []dependencies.Emoji `json:"emoji"`
}
