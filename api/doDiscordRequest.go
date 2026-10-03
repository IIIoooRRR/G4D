package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

func (c *DiscordClient) DoDiscordRequest(method _const.Method, uri string, body []byte) ([]byte, error) {
	url := GetURL("https://discord.com/api/v10", uri)
	code, resp, err := c.doRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if code < 200 || code >= 300 {
		return nil, &discordError{Code: code, Message: resp}
	}
	return resp, nil
}

func (c *DiscordClient) DoDiscordLimitRequest(ctx context.Context, method _const.Method, uri string, body []byte) ([]byte, error) {
	url := GetURL("https://discord.com/api/v10", uri)
	limiter := c.GetOrNewBucket(uri)
	for at := 0; at < 2; at++ {
		if err := limiter.Wait(ctx); err != nil {
			return nil, err
		}

		code, resp, err := c.doRequest(method, url, body)
		if err != nil {
			return nil, err
		}

		if code < 200 || code >= 300 {
			if code == http.StatusTooManyRequests {
				var retry dependencies.RetryResponse
				if err := json.Unmarshal(resp, &retry); err == nil {
					timer := time.NewTimer(retry.After * time.Second)

					select {
					case <-ctx.Done():
						timer.Stop()
						return nil, ctx.Err()
					case <-timer.C:
						continue
					}
				} else {
					return nil, err
				}

			}
			return nil, &discordError{Code: code, Message: resp}
		}
		return resp, nil
	}
	return nil, nil
}
