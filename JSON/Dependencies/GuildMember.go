package Dependencies

type GuildMember struct {
	User     User     `json:"user"`
	Roles    []string `json:"roles"`
	JoinedAt string   `json:"joined_at"`
	Deaf     bool     `json:"deaf"`
	Mute     bool     `json:"mute"`
}
