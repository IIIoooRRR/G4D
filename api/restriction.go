package api

import (
	"context"
	"time"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func (c *DiscordClient) BanUser(guildId _const.GuildId, userId _const.UserId, reason *string, timeMessDelete int) error {
	uri := GetURI("/guilds/", string(guildId), "/bans/", string(userId))
	body := Ban{
		DeleteMessageSeconds: timeMessDelete,
		Reason:               reason,
	}
	jsonBody, err := parse.Marshal(body)
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest(_const.Patch, uri, jsonBody)
	return err
}
func (c *DiscordClient) MuteUser(guildId _const.GuildId, userId _const.UserId, dur time.Duration) error {
	until := time.Now().Add(dur).Format(time.RFC3339)
	uri := GetURI("/guilds/", string(guildId), "/members/", string(userId))
	body := Mute{Duration: until}
	jsonBody, err := parse.Marshal(body)
	if err != nil {

	}
	_, err = c.DoDiscordRequest(_const.Patch, uri, jsonBody)
	return err
}

func (c *DiscordClient) BanUserWithLimit(guildId _const.GuildId, userId _const.UserId, reason *string, timeMessDelete int) error {
	var uri = GetURI("/guilds/", string(guildId), "/bans/", string(userId))
	body := Ban{
		DeleteMessageSeconds: timeMessDelete,
		Reason:               reason,
	}
	jsonBody, err := parse.Marshal(body)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, _const.Patch, uri, jsonBody)
	return err
}
func (c *DiscordClient) MuteUserWithLimit(guildId _const.GuildId, userId _const.UserId, dur time.Duration) error {
	until := time.Now().Add(dur).Format(time.RFC3339)
	uri := GetURI("/guilds/", string(guildId), "/members/", string(userId))
	body := Mute{Duration: until}
	jsonBody, err := parse.Marshal(body)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, _const.Patch, uri, jsonBody)
	return err
}

type Ban struct {
	DeleteMessageSeconds int     `json:"delete_message_seconds"`
	Reason               *string `json:"reason,omitempty"`
}
type Mute struct {
	Duration string `json:"communication_disabled_until"`
}
