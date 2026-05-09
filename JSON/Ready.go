package JSON

import (
	D_param "github.com/IIIoooRRR/G4D/JSON/Dependencies"
)

// Hello получает бот при подключении
type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type ReadyEvent struct {
	Version   int          `json:"v"`
	SessionID string       `json:"session_id"`
	User      D_param.User `json:"user"`
	Guilds    []struct {
		ID          string `json:"id"`
		Unavailable bool   `json:"unavailable"`
	} `json:"guilds"`
	ResumeGatewayURL string `json:"resume_gateway_url"`
}
