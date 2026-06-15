package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/IIIoooRRR/G4D/g4d"
	"go.uber.org/zap"
)

func DoDiscordRequest(method, uri string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("https://discord.com/api/v10%s", uri)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+g4d.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		g4d.CurrentBot().Logger.Info("response status", zap.String("", resp.Status))
		return nil, errors.New("bad Request")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("response body read error")
	}
	return respBody, nil
}
