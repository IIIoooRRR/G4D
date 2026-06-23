package shema

import (
	"github.com/IIIoooRRR/G4D/model/_const"
)

type Channel struct {
	Id               _const.ChannelId `json:"id"`
	Name             string           `json:"name"`
	Type             int              `json:"type"`
	Topic            string           `json:"topic,omitempty"`
	ThreadMetadata   *ThreadMetadata  `json:"thread_metadata,omitempty"`
	ThreadMember     *ThreadMember    `json:"thread_member,omitempty"`
	Nsfw             bool             `json:"nsfw,omitempty"`
	ParentID         *string          `json:"parent_id,omitempty"`
	BitRate          int              `json:"bit_rate,omitempty"`
	UserLimit        int              `json:"user_limit,omitempty"`
	RateLimitPerUser int              `json:"rate_limit_per_user,omitempty"`
}
type ThreadMetadata struct {
	Archived            bool    `json:"archived"`
	AutoArchiveDuration int     `json:"auto_archive_duration"`
	ArchiveTimestamp    *string `json:"archive_timestamp,omitempty"`
	Locked              bool    `json:"locked"`
	Invitable           bool    `json:"invitable,omitempty"`
}

type ThreadMember struct {
	ID            _const.ThreadId `json:"id,omitempty"`
	UserID        _const.UserId   `json:"user_id,omitempty"`
	JoinTimestamp string          `json:"join_timestamp"`
	Flags         int             `json:"flags"`
}
type ChannelDelete struct {
	ID       _const.ChannelId `json:"id"`
	GuildID  *_const.GuildId  `json:"guild_id,omitempty"`
	Type     int              `json:"type"`
	ParentID *string          `json:"parent_id,omitempty"`
}
