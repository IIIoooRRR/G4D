package gateway

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func NewGateway(bufferSize _const.BufferSize, intents _const.Intents, activity []customize.Activity, netStatus _const.NetStatus) *Receiver {
	return &Receiver{
		Queue: make(chan *parse.RawEvent, bufferSize),
		Presence: &customize.PresenceUpdate{
			Since:      0,
			Activities: activity,
			Status:     netStatus,
			Afk:        false,
		},
		Intents: intents,
	}
}

func Intents(intents ...int) _const.Intents {
	var intent int
	for _, i := range intents {
		intent = intent | i
	}
	return _const.Intents(intent)
}

func Activity(activity ...customize.Activity) []customize.Activity {
	return activity
}

func NetStatus(netStatus _const.NetStatus) _const.NetStatus {
	return netStatus
}
func (r *Receiver) WithDescription(description string) *Receiver {
	r.Presence.Activities = append(r.Presence.Activities, customize.Activity{
		Name: description,
		Type: 4,
	})
	return r
}
func BufferSize(size uint) _const.BufferSize {
	return _const.BufferSize(size)
}
