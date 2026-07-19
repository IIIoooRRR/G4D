package api

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/schema"
)

func (c *DiscordClient) SendInteractionMessage(event *schema.Interaction, msg schema.InteractionResponse) error {
	uri := fmt.Sprintf("/interactions/%v/%s/callback", event.ID, event.Token)
	jsonData, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest("POST", uri, jsonData)
	return err
}

func (c *DiscordClient) SendMessage(ToChannel _const.ChannelId, msg *schema.SendMessage) error {
	uri := fmt.Sprintf("/channels/%s/messages", ToChannel)
	body := msg
	jsonBody, err := parse.Marshal(body) //delaem жсон из message
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest("POST", uri, jsonBody)
	return err
}
func (c *DiscordClient) EditMessage(ToChannel _const.ChannelId, msgId _const.MessageId, msg *schema.MessageEdit) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	jsonBody, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = c.DoDiscordRequest("PATCH", uri, jsonBody)
	return err
}
func (c *DiscordClient) DeleteMessage(ToChannel _const.ChannelId, msgId _const.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	_, err := c.DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
