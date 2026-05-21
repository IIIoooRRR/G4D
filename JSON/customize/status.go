package customize

type PresenceUpdate struct {
	Since      int        `json:"since,omitempty"`
	Activities []Activity `json:"activities,omitempty"`
	Status     string     `json:"status,omitempty"`
	Afk        bool       `json:"afk,omitempty"`
}

type Activity struct {
	Name          string  `json:"name"`
	Type          int     `json:"type"`              //type.Activity
	URL           *string `json:"url,omitempty"`     // for streaming
	Details       *string `json:"details,omitempty"` // "Playing Minecraft"
	State         *string `json:"state,omitempty"`   // "On Hypixel"
	ApplicationID *string `json:"application_id,omitempty"`
	Flags         *int    `json:"flags,omitempty"`
	Instance      *bool   `json:"instance,omitempty"`
}
