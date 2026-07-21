package api

import (
	"context"
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/schema"
)

func (c *DiscordClient) CreateChannel(guildId _const.GuildId, channel *schema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", guildId)
	_, err = c.DoDiscordRequest("POST", endpoint, jsonBody)
	return err
}

func (c *DiscordClient) DeleteChannel(channelId _const.ChannelId, channel *schema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = c.DoDiscordRequest("DELETE", endpoint, jsonBody)
	return err
}

func (c *DiscordClient) ChangeChannels(channelId _const.ChannelId, channel *schema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = c.DoDiscordRequest("PATCH", endpoint, jsonBody)
	return err
}

func (c DiscordClient) GetChannel(channelId _const.ChannelId) (*schema.Channel, error) {
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	abstract, err := c.DoDiscordRequest("GET", endpoint, []byte{})
	if err != nil {
		return nil, err
	}
	return parse.ToChannel(abstract)
}

func (c *DiscordClient) CreateChannelWithLimit(guildId _const.GuildId, channel *schema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", guildId)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, "POST", endpoint, jsonBody)
	return err
}
func (c DiscordClient) GetChannelWithLimit(channelId _const.ChannelId) (*schema.Channel, error) {
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	abstract, err := c.DoDiscordLimitRequest(ctx, "GET", endpoint, []byte{})
	if err != nil {
		return nil, err
	}
	return parse.ToChannel(abstract)
}

func (c *DiscordClient) DeleteChannelWithLimit(channelId _const.ChannelId, channel *schema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, "DELETE", endpoint, jsonBody)
	return err
}

func (c *DiscordClient) ChangeChannelsWithLimit(channelId _const.ChannelId, channel *schema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, "PATCH", endpoint, jsonBody)
	return err
}
