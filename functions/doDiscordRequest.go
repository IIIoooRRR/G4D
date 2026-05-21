package functions

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IIIoooRRR/G4D/G4D"
)

func doDiscordRequest(method, uri string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("https://discord.com/api/v10%s", uri)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		log.Print("[REQUEST] ", resp.Status)
		return nil, errors.New("Bad Request")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[RESPONSE READER] %s", err)
	}
	return respBody, nil
}
