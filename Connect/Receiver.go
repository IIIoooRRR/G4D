package Connect

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Receiver struct {
	token     string
	Intents   int
	connectWS *websocket.Conn
	sessionID string
	lastSeq   int
	interval  time.Duration
	resumeURL string
	Queue     chan *RawEvent
	cancel    context.CancelFunc
	QueueSize int
	connMutex sync.Mutex
	ctx       context.Context
}

type RawEvent struct {
	Type string          `json:"t"`
	Data json.RawMessage `json:"d"`
}
