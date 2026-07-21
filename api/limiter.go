package api

import (
	"net/http"
	"sync"
	"time"

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
	client  *http.Client
	buckets map[string]*limiter
	rwmu    sync.RWMutex
	token   *string
	logger  *zap.Logger
	Timeout time.Duration
}
type limiter struct {
	rate.Limiter
	TTL time.Time
}

func NewClient(token *string, clientTimeout int, logger *zap.Logger) *DiscordClient {
	client := &DiscordClient{
		token:   token,
		client:  &http.Client{Timeout: time.Duration(clientTimeout) * time.Second},
		buckets: make(map[string]*limiter),
		logger:  logger,
		Timeout: time.Duration(clientTimeout) * time.Second,
	}
	go client.deleteBucket()
	return client
}

func (c *DiscordClient) newBucket(uri string) *limiter {
	if lim, ok := c.getBucket(uri); ok {
		return lim
	}
	bucket := &limiter{
		Limiter: *rate.NewLimiter(rate.Limit(5), 1),
		TTL:     time.Now().Add(10 * time.Minute),
	}
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
func (c *DiscordClient) deleteBucket() {
	for {
		time.Sleep(10 * time.Minute)
		c.rwmu.Lock()
		for uri, lim := range c.buckets {
			if time.Now().After(lim.TTL) {
				delete(c.buckets, uri)
			}
		}
		c.rwmu.Unlock()
	}
}

func (c *DiscordClient) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}
