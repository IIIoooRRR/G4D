/*
DoDiscordLimitRequest_Old was removed because moving the structure is a very tedious task.
I considered it an unnecessary waste of time and memory, as only 3 benchmarks were shown. Please,
if you handle this neatly, send me your solution via email/Telegram.
*/
package api_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/_const"
)

var c = api.NewClient(new(""), 5*time.Second)
var old = cl{client: http.DefaultClient, token: new("")}

func BenchmarkDoDiscordLimitRequest(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = c.DoDiscordLimitRequest(ctx, _const.Get, "/users/@me", nil)
	}
}
func BenchmarkDoDiscordRequest(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = c.DoDiscordRequest(_const.Get, "/users/@me", nil)
	}
}
func BenchmarkDoDiscordRequest_Old(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = old.DoDiscordRequest("GET", "/users/@me", nil)
	}
}

// garbage(old req-methods

type cl struct {
	client *http.Client
	token  *string
}

func (c *cl) DoDiscordRequest(method, uri string, body []byte) ([]byte, error) {
	url := api.GetURL("https://discord.com/api/v10", uri)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+*c.token)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		return nil, errors.New("bad Request")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("response body read error")
	}
	return respBody, nil
}
