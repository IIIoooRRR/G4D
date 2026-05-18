package customize

const (
	Game      = 0
	Streaming = 1
	Listening = 2
	Watching  = 3
)

type PresenceUpdate struct {
	Since      int        `json:"since,omitempty"`
	Activities []Activity `json:"activities,omitempty"`
	Status     string     `json:"status,omitempty"`
	Afk        bool       `json:"afk,omitempty"`
}

type Activity struct {
	Name string `json:"name,omitempty"`
	Type int    `json:"type,omitempty"` // 0=Game, 1=Streaming, 2=Listening, 3=Watching
}
