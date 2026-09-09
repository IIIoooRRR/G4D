package api

import (
	"context"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/schema"
)

func (c *DiscordClient) SendInteractionMessage(event *schema.Interaction, msg schema.InteractionResponse) error {
	uri := GetURI("/interactions/", event.ID, "/", event.Token, "/callback")
	jsonData, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest("POST", uri, jsonData)
	return err
}

func (c *DiscordClient) SendMessage(toChannel _const.ChannelId, msg *schema.SendMessage) error {
	uri := GetURI("/channels/", string(toChannel), "/messages")
	body := msg
	jsonBody, err := parse.Marshal(body)
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest("POST", uri, jsonBody)
	return err
}
func (c *DiscordClient) EditMessage(toChannel _const.ChannelId, msgId _const.MessageId, msg *schema.MessageEdit) error {
	uri := GetURI("/channels/", string(toChannel), "/messages/", string(msgId))
	jsonBody, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest("PATCH", uri, jsonBody)
	return err
}
func (c *DiscordClient) DeleteMessage(toChannel _const.ChannelId, msgId _const.MessageId) error {
	uri := GetURI("/channels/", string(toChannel), "/messages/", string(msgId))
	_, err := c.DoDiscordRequest("DELETE", uri, nil)
	return err
}

func (c *DiscordClient) SendMessageWithLimit(toChannel _const.ChannelId, msg *schema.SendMessage) error {
	uri := GetURI("/channels/", string(toChannel), "/messages")
	body := msg
	jsonBody, err := parse.Marshal(body)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, "POST", uri, jsonBody)
	return err
}
func (c *DiscordClient) EditMessageWithLimit(toChannel _const.ChannelId, msgId _const.MessageId, msg *schema.MessageEdit) error {
	uri := GetURI("/channels/", string(toChannel), "/messages/", string(msgId))
	jsonBody, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	_, err = c.DoDiscordLimitRequest(ctx, "PATCH", uri, jsonBody)
	return err
}
func (c *DiscordClient) DeleteMessageWithLimit(toChannel _const.ChannelId, msgId _const.MessageId) error {
	uri := GetURI("/channels/", string(toChannel), "/messages/", string(msgId))
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	_, err := c.DoDiscordLimitRequest(ctx, "DELETE", uri, nil)
	return err
}
