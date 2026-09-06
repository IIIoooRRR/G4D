package customize

import "github.com/IIIoooRRR/G4D/model/_const"

type PresenceUpdate struct {
	Since      int              `json:"since,omitempty"`
	Activities []Activity       `json:"activities,omitempty" yaml:"activities,omitempty"`
	Status     _const.NetStatus `json:"status,omitempty" yaml:"status,omitempty"`
	Afk        bool             `json:"afk,omitempty" yaml:"afk,omitempty"`
}

type Activity struct {
	Name          string          `json:"name" yaml:"name"`
	Type          _const.Activity `json:"type" yaml:"type"`                           //_const.customize
	URL           string          `json:"url,omitempty" yaml:"url,omitempty"`         // for streaming
	Details       string          `json:"details,omitempty" yaml:"details,omitempty"` // "Playing Minecraft"
	State         string          `json:"state,omitempty" yaml:"state,omitempty"`     // "On Hypixel"
	ApplicationID *string         `json:"application_id,omitempty"`
	Flags         *int            `json:"flags,omitempty"`
	Instance      *bool           `json:"instance,omitempty"`
}
