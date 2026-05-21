package connect

import (
	"github.com/IIIoooRRR/G4D/JSON/Dependencies"
	"github.com/IIIoooRRR/G4D/JSON/customize"
)

func NewGateway() *Receiver {
	return &Receiver{
		Presence: &customize.PresenceUpdate{
			Since:      0,
			Activities: []customize.Activity{},
			Status:     "online",
			Afk:        false,
		},
	}
}

func (r *Receiver) WithIntents(intents int) *Receiver {
	r.Intents = intents
	return r
}

func (r *Receiver) WithActivity(activity customize.Activity) *Receiver {
	r.Presence.Activities = append(r.Presence.Activities, activity)
	return r
}

func (r *Receiver) WithNetStatus(netStatus string) *Receiver {
	r.Presence.Status = netStatus
	return r
}
func (r *Receiver) WithDescription(description string, emoji Dependencies.Emoji) *Receiver {
	r.Presence.Activities = append(r.Presence.Activities, customize.Activity{
		Name: description,
		Type: 4,
	})
	return r
}

func (r *Receiver) WithQueueSize(size int) *Receiver {
	r.QueueSize = size
	return r
}
