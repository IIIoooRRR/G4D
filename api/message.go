package api

import (
	"encoding/json"
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/_struct"
)

func SendInteractionMessage(event *_struct.Interaction, msg _struct.InteractionResponse) error {
	uri := fmt.Sprintf("/interactions/%s/%s/callback", event.ID, event.Token)
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("POST", uri, jsonData)
	return err
}

func SendMessage(ToChannel _const.ChannelId, msg *_struct.SendMessage) error {
	uri := fmt.Sprintf("/channels/%s/messages", ToChannel)
	body := msg
	jsonBody, err := json.Marshal(body) //delaem жсон из message
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("POST", uri, jsonBody)
	return err
}
func EditMessage(ToChannel _const.ChannelId, msgId _const.MessageId, msg *_struct.MessageEdit) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("PATCH", uri, jsonBody)
	return err
}
func DeleteMessage(ToChannel _const.ChannelId, msgId _const.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	_, err := DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
