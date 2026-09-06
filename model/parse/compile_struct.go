package parse

import (
	"encoding/json"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type RawEvent struct {
	Type  string          `json:"t"`
	Data  json.RawMessage `json:"d"`
	idx   int
	cache *[]EventEntry
}
