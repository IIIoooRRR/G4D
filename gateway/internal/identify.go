package json

import (
	"github.com/IIIoooRRR/G4D/model/customize"
)

type Identify struct {
	Token      string                    `json:"token"`
	Properties IdentifyProperties        `json:"properties"`
	Intents    int                       `json:"intents"`
	Shard      []int                     `json:"shard,omitempty"`
	Presence   *customize.PresenceUpdate `json:"presence,omitempty"`
}

type IdentifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}
