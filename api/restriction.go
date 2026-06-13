package api

import (
	"fmt"
	"time"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func BanUser(guildId _const.GuildId, userId _const.UserId, reason *string, timeMessDelete int) error {
	var uri = fmt.Sprintf("/guilds/%s/bans/%s", guildId, userId)
	body := Ban{
		DeleteMessageSeconds: timeMessDelete,
		Reason:               reason,
	}
	jsonBody, err := parse.Marshal(body)
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("PATCH", uri, jsonBody)
	return err
}
func MuteUser(guildId _const.GuildId, userId _const.UserId, dur time.Duration) error {
	until := time.Now().Add(dur).Format(time.RFC3339)
	var url = fmt.Sprintf("/guilds/%s/members/%s", guildId, userId)
	body := map[string]interface{}{
		"communication_disabled_until": until,
	}
	jsonBody, err := parse.Marshal(body)
	if err != nil {

	}
	_, err = DoDiscordRequest("PATCH", url, jsonBody)
	return err
}

type Ban struct {
	DeleteMessageSeconds int     `json:"delete_message_seconds"`
	Reason               *string `json:"reason,omitempty"`
}
