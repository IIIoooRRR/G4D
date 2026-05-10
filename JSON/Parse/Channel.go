package Parse

type Channel struct {
	Name             string `json:"name"`
	Type             int    `json:"type"`
	Topic            string `json:"topic,omitempty"`
	Nsfw             bool   `json:"nsfw,omitempty"`
	ParentID         int    `json:"parent_id,omitempty"`
	BitRate          int    `json:"bit_rate,omitempty"`
	UserLimit        int    `json:"user_limit,omitempty"`
	RateLimitPerUser int    `json:"rate_limit_per_user,omitempty"`
}


