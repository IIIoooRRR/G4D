package api

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/shema"
)

func CreateChannel(guildId _const.GuildId, channel *shema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", guildId)
	_, err = DoDiscordRequest("POST", endpoint, jsonBody)
	return err
}

func DeleteChannel(channelId _const.ChannelId, channel *shema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = DoDiscordRequest("DELETE", endpoint, jsonBody)
	return err
}

func ChangeChannels(channelId _const.ChannelId, channel *shema.Channel) error {
	jsonBody, err := parse.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = DoDiscordRequest("PATCH", endpoint, jsonBody)
	return err
}

func GetChannel(channelId _const.ChannelId) (*shema.Channel, error) {
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	abstract, err := DoDiscordRequest("GET", endpoint, []byte{})
	if err != nil {
		return nil, err
	}
	return parse.ToChannel(abstract)
}
