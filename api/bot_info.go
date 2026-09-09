package api

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/model/dependencies"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func (c *DiscordClient) GetBotInfo() (*dependencies.User, error) {
	resp, err := c.DoDiscordRequest("GET", "/users/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return parse.ToUser(resp)
}
