package Parse

import "github.com/IIIoooRRR/G4D/JSON/Type"

type Channel struct {
	Name             string          `json:"name"`
	Type             int             `json:"type"`
	Topic            string          `json:"topic,omitempty"`
	ThreadMetadata   *ThreadMetadata `json:"thread_metadata,omitempty"`
	ThreadMember     *ThreadMember   `json:"thread_member,omitempty"`
	Nsfw             bool            `json:"nsfw,omitempty"`
	ParentID         int             `json:"parent_id,omitempty"`
	BitRate          int             `json:"bit_rate,omitempty"`
	UserLimit        int             `json:"user_limit,omitempty"`
	RateLimitPerUser int             `json:"rate_limit_per_user,omitempty"`
}
type ThreadMetadata struct {
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked"`
	Invitable           bool   `json:"invitable,omitempty"`
}

type ThreadMember struct {
	ID            Type.ThreadId `json:"id,omitempty"`
	UserID        Type.UserId   `json:"user_id,omitempty"`
	JoinTimestamp string        `json:"join_timestamp"`
	Flags         int           `json:"flags"`
}
