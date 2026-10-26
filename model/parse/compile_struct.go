package parse

import (
	"encoding/json"
	"unsafe"
)

type RawEvent struct {
	Type  string          `json:"t"`
	Data  json.RawMessage `json:"d"`
	idx   int
	cache []EventEntry
}

func (ev *RawEvent) Reset() *RawEvent {
	ev.cache = nil
	return ev
}

type eface struct {
	Type unsafe.Pointer
	Data unsafe.Pointer
}
