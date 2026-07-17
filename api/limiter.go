package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/IIIoooRRR/G4D/g4d"
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

var GlobalClient = newLimiter(5)

type clientWithLimiter struct {
	*http.Client
	buckets atomic.Pointer[map[string]*limiter]
}
type limiter struct {
	rate.Limiter
	TTL time.Time
}

func newLimiter(contextTimeout time.Duration) *clientWithLimiter {
	mp := make(map[string]*limiter)
	l := &clientWithLimiter{}
	l.Client = &http.Client{}
	l.buckets.Store(&mp)
	l.Timeout = contextTimeout
	go l.DeleteBucket()
	return l
}

func (l *clientWithLimiter) NewBucket(uri string) *limiter {

	bucket := &limiter{
		Limiter: *rate.NewLimiter(rate.Limit(5), 1),
		TTL:     time.Now().Add(10 * time.Minute),
	}
	for {
		oldptr := l.buckets.Load()
		newptr := maps.Clone(*oldptr)
		newptr[uri] = bucket
		if l.buckets.CompareAndSwap(oldptr, &newptr) {
			break
		}
	}
	return bucket
}
func (l *clientWithLimiter) GetBucket(uri string) (*limiter, bool) {
	lim, ok := (*l.buckets.Load())[uri]
	return lim, ok
}
func (l *clientWithLimiter) DeleteBucket() {
	for {
		time.Sleep(10 * time.Minute)
		for {
			ptr := l.buckets.Load()
			newMap := maps.Clone(*ptr)
			for uri, lim := range newMap {
				if time.Now().After(lim.TTL) {
					delete(newMap, uri)
				}
			}
			if l.buckets.CompareAndSwap(ptr, &newMap) {
				break
			}
		}
	}
}
func (l *clientWithLimiter) DoRequest(method, uri string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("https://discord.com/api/v10%s", uri)
	limiter, ok := l.GetBucket(uri)
	if !ok {
		limiter = l.NewBucket(uri)
	}

	ctx, cancel := context.WithTimeout(context.Background(), l.Timeout)
	defer cancel()
	if err := limiter.Wait(ctx); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+g4d.CurrentBot().Token)
	resp, err := l.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		g4d.CurrentBot().Logger.Warn("response error", zap.String("uri:", uri), zap.String("status:", resp.Status))
		return nil, errors.New("response error: " + resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("response body read error")
	}
	return respBody, nil
}
func (l *clientWithLimiter) SetTimeout(timeout time.Duration) {
	l.Timeout = timeout
}
