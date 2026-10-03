package json

type Resume struct {
	Op   int   `json:"op"`
	Data RData `json:"d"`
}

type RData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int64  `json:"seq"`
}
