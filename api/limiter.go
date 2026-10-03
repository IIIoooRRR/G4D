package api

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

/*
Limiter provides only a basic set of functions to work properly with the discord API (safely, without 429 Too many requests)
However, please note: I did not do error protection or add mutexes to the set, etc.
Therefore, you want to perform all configuration operations at the initialization stage
without touching the client after initialization to avoid panicking during execution.
I abandoned cas blocks and operations, because after diagnostics on my own bot,
I saw how many lines were copied and translated for nothing several times.
I also switched the structure to a more gentle read-write mutex, which improves performance,
and also makes reading/writing atomic at the method level, but not at the map level.
*/
type DiscordClient struct {
	client  *fasthttp.Client
	buckets map[string]*limiter
	rwmu    sync.RWMutex
	token   *string
	logger  *zap.Logger
	timeout time.Duration
	AppId   *string
	ctx     context.Context
}
type limiter struct {
	rate.Limiter
	TTL atomic.Int64
}

func NewClient(token *string, clientTimeout time.Duration) *DiscordClient {
	client := &DiscordClient{
		token: token,
		client: &fasthttp.Client{
			TLSConfig: &tls.Config{
				ClientSessionCache: tls.NewLRUClientSessionCache(100),
				MinVersion:         tls.VersionTLS12,
			},
			Dial: func(addr string) (net.Conn, error) {
				return fasthttp.DialTimeout(addr, 10*time.Second)
			},
			MaxIdleConnDuration: 30 * time.Second,
			ReadTimeout:         clientTimeout,
			WriteTimeout:        clientTimeout * 3,
		},
		buckets: make(map[string]*limiter),
		timeout: clientTimeout,
	}
	return client
}

func (c *DiscordClient) GetOrNewBucket(uri string) *limiter {
	if lim, ok := c.getBucket(uri); ok {
		return lim
	}
	bucket := &limiter{
		Limiter: *rate.NewLimiter(rate.Limit(5), 1),
		TTL:     atomic.Int64{},
	}

	bucket.TTL.Store(time.Now().Add(10 * time.Minute).UnixNano())

	c.rwmu.Lock()
	c.buckets[uri] = bucket
	c.rwmu.Unlock()

	return bucket
}
func (c *DiscordClient) getBucket(uri string) (*limiter, bool) {
	c.rwmu.RLock()
	defer c.rwmu.RUnlock()
	bucket, ok := c.buckets[uri]
	return bucket, ok
}
func (c *DiscordClient) RunDelete() {
	select {
	case <-c.ctx.Done():
		return
	case <-time.After(c.timeout):
		for {
			c.rwmu.Lock()
			for uri, lim := range c.buckets {
				if time.Now().UnixNano() > lim.TTL.Load()+int64(time.Minute*10) {
					delete(c.buckets, uri)
				}
			}
			c.rwmu.Unlock()
		}
	}
}

func (c *DiscordClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}
func (c *DiscordClient) SetLogger(logger *zap.Logger) {
	c.logger = logger
}
func (c *DiscordClient) SetContext(ctx context.Context) {
	c.ctx = ctx
}
