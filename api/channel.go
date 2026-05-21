package api

import (
	"encoding/json"
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/_struct"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func CreateChannel(guildId _const.GuildId, channel *_struct.Channel) error {
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", guildId)
	_, err = DoDiscordRequest("POST", endpoint, jsonBody)
	return err
}

func DeleteChannel(channelId _const.ChannelId, channel *_struct.Channel) error {
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = DoDiscordRequest("DELETE", endpoint, jsonBody)
	return err
}

func ChangeChannels(channelId _const.ChannelId, channel *_struct.Channel) error {
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = DoDiscordRequest("PATCH", endpoint, jsonBody)
	return err
}

func GetChannel(channelId _const.ChannelId) (*_struct.Channel, error) {
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	abstract, err := DoDiscordRequest("GET", endpoint, []byte{})
	if err != nil {
		return nil, err
	}
	return parse.ToChannel(abstract)
}
