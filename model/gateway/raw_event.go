package gateway

import "encoding/json"

type RawEvent struct {
	Type string          `json:"t"`
	Data json.RawMessage `json:"d"`
}
