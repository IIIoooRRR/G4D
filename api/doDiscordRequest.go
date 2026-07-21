package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (c *DiscordClient) DoDiscordRequest(method, uri string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("https://discord.com/api/v10%s", uri)
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
		c.logger.Info("response status", zap.String("", resp.Status))
		return nil, errors.New("bad Request")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("response body read error")
	}
	return respBody, nil
}

func (c *DiscordClient) DoDiscordLimitRequest(ctx context.Context, method, uri string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("https://discord.com/api/v10%s", uri)
	limiter, ok := c.getBucket(uri)
	if !ok {
		limiter = c.newBucket(uri)
	}
	if err := limiter.Wait(ctx); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
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
		c.logger.Warn("response error", zap.String("uri:", uri), zap.String("status:", resp.Status))
		return nil, errors.New("response error: " + resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("response body read error")
	}
	return respBody, nil
}
