package gateway

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Receiver struct {
	token     *string
	Intents   _const.Intents
	connectWS *websocket.Conn
	sessionID string
	lastSeq   atomic.Int64
	interval  time.Duration
	resumeURL string
	Queue     chan *parse.RawEvent
	cancel    context.CancelFunc
	connMutex sync.Mutex
	ctx       context.Context
	Presence  *customize.PresenceUpdate
	logger    *zap.Logger
	dLogger   *zap.Logger // dispatch.go logger
	initOnce  sync.Once
}
