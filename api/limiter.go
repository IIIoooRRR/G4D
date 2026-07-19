package api

import (
	"maps"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

/*
The limiter provides only a basic set of functions for proper operation with the discord API (safely, without 429 Too Many Requests)
However, please note: I did not do foolproof and did not add mutexes to the Set, etc.0
Consequently, you want to perform all configuration operations at the initialization stage,
without touching the client after it has initialized, in order to avoid panics in runtime.
This module is not built into the main library for better performance
(read the code /model/parse/... for this). to work with it, write your own wrappers yourself.
I've provided a convenient API (method, uri, body), so it won't be difficult.
I’m not sure about the cleanup process; if you have a large bot, frequently adding URIs might interfere with CAS operations.
I’m thinking of replacing CAS with a standard swap to reduce conflicts.
However, this could lead to data loss in the limiters—triggering a panic—and working with copies is extremely costly.
*/
var once sync.Once

type DiscordClient struct {
	*http.Client
	buckets atomic.Pointer[map[string]*limiter]
	token   *string
	Logger  *zap.Logger
}
type limiter struct {
	rate.Limiter
	TTL time.Time
}

func NewClient(token *string, contextTimeoutDuration time.Duration, logger *zap.Logger) *DiscordClient {
	var client *DiscordClient
	once.Do(func() {
		mp := make(map[string]*limiter)
		client = &DiscordClient{
			token:  token,
			Client: &http.Client{Timeout: contextTimeoutDuration},
			Logger: logger,
		}
		client.buckets.Store(&mp)
		go client.deleteBucket()
	})
	return client
}

func (c *DiscordClient) NewBucket(uri string) *limiter {

	bucket := &limiter{
		Limiter: *rate.NewLimiter(rate.Limit(5), 1),
		TTL:     time.Now().Add(10 * time.Minute),
	}
	for {
		oldptr := c.buckets.Load()
		newptr := maps.Clone(*oldptr)
		newptr[uri] = bucket
		if c.buckets.CompareAndSwap(oldptr, &newptr) {
			break
		}
	}
	return bucket
}
func (c *DiscordClient) GetBucket(uri string) (*limiter, bool) {
	lim, ok := (*c.buckets.Load())[uri]
	return lim, ok
}
func (c *DiscordClient) deleteBucket() {
	for {
		time.Sleep(10 * time.Minute)
		for {
			ptr := c.buckets.Load()
			newMap := maps.Clone(*ptr)
			for uri, lim := range newMap {
				if time.Now().After(lim.TTL) {
					delete(newMap, uri)
				}
			}
			if c.buckets.CompareAndSwap(ptr, &newMap) {
				break
			}
		}
	}
}

func (c *DiscordClient) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}
