package gateway

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Receiver struct {
	token     *string
	Intents   int
	connectWS *websocket.Conn
	sessionID string
	lastSeq   atomic.Int64
	interval  time.Duration
	resumeURL string
	Queue     chan *gateway.RawEvent
	cancel    context.CancelFunc
	QueueSize int
	connMutex sync.Mutex
	ctx       context.Context
	Presence  *customize.PresenceUpdate
	logger    *zap.Logger
	dLogger   *zap.Logger // dispatch.go logger
}
