package Event

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/IIIoooRRR/G4D/G4D"
)

type Channel struct {
	Name             string `json:"name"`
	Type             int    `json:"type"`
	Topic            string `json:"topic,omitempty"`
	Nsfw             bool   `json:"nsfw,omitempty"`
	ParentID         int    `json:"parent_id,omitempty"`
	BitRate          int    `json:"bit_rate,omitempty"`
	UserLimit        int    `json:"user_limit,omitempty"`
	RateLimitPerUser int    `json:"rate_limit_per_user,omitempty"`
}

const (
	ChannelTypeGuildText          = 0  // Текстовый канал на сервере
	ChannelTypeGuildVoice         = 2  // Голосовой канал
	ChannelTypeGuildAnnouncement  = 5  // Канал объявлений
	ChannelTypeAnnouncementThread = 10 // Тред объявления (сообщение)
	ChannelTypePublicThread       = 11 // Публичный тред
	ChannelTypePrivateThread      = 12 // Приватный тред
	ChannelTypeGuildStageVoice    = 13 // Сцена (Stage Channel)
	ChannelTypeGuildDirectory     = 14 // Каталог (для связки каналов)
	ChannelTypeGuildMedia         = 16 // Медиа-канал
)

func (channel *Channel) CreateChannel(guildId string) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/guilds/%s/channels", guildId)
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("[DISCORD] Failed to create channel")
	}
	return nil
}

func (channel *Channel) DeleteChannel(channelId string) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s", channelId)
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("[DISCORD] Failed to delete channel")
	}
	return nil
}

func (channel *Channel) ChangeChannels(channelID string) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s", channelID)
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("[DISCORD] Failed to update channels")
	}
	return nil
}

func GetChannel(channelId string) (*Channel, error) {
	if G4D.CurrentBot().Cache != nil {
		cache, err := G4D.CurrentBot().Cache.GetChannel(channelId)
		if err == nil {
			return &cache, err
		}

	}
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s", channelId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &Channel{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.New("[DISCORD] Failed to get channel")
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var channel Channel
	err = json.Unmarshal(bodyBytes, &channel)
	if err != nil {
		return &Channel{}, err
	}

	return &channel, nil
}
