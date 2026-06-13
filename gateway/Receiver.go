package gateway

import (
	"context"
	"sync"
	"time"

	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/gateway"
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
	Queue     chan *gateway.RawEvent
	cancel    context.CancelFunc
	QueueSize int
	connMutex sync.Mutex
	ctx       context.Context
	Presence  *customize.PresenceUpdate
}
