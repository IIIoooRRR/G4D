package api

import (
	"context"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func (c *DiscordClient) GetUserInfo(guildId _const.GuildId, userId _const.UserId) (*entity.GuildMember, error) {
	uri := GetURI("/guilds/", string(guildId), "/members/", string(userId))
	bytes, err := c.DoDiscordRequest(_const.Get, uri, nil)
	if err != nil {
		return nil, err
	}
	member := new(entity.GuildMember)
	err = parse.Unmarshal(bytes, member)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (c *DiscordClient) GetUserInfoWithLimit(guildId _const.GuildId, userId _const.UserId) (*entity.GuildMember, error) {
	uri := GetURI("/guilds/", string(guildId), "/members/", string(userId))
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	bytes, err := c.DoDiscordLimitRequest(ctx, _const.Get, uri, nil)
	if err != nil {
		return nil, err
	}

	member := new(entity.GuildMember)
	err = parse.Unmarshal(bytes, member)
	if err != nil {
		return nil, err
	}
	return member, nil
}
