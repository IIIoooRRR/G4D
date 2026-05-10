package Functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IIIoooRRR/G4D/G4D"
)

func BanUser(guildId string, userId string) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/guilds/%s/bans/%s", guildId, userId)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func MuteUser(guildId string, userId string, dur time.Duration) error {
	until := time.Now().Add(dur).Format(time.RFC3339)
	var url = fmt.Sprintf("https://discord.com/api/v10/guilds/%s/members/%s", guildId, userId)
	body := &until
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
