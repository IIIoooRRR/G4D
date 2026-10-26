package types

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/schema"
)

var features = map[string]func() any{
	_const.EventMessageCreate:     func() any { return &schema.GetMessage{} },
	_const.EventMessageUpdate:     func() any { return &schema.MessageEdit{} },
	_const.EventMessageDelete:     func() any { return &schema.MessageDelete{} },
	_const.EventMessageDeleteBulk: func() any { return &schema.MessageDeleteBulk{} },

	_const.EventInteractionCreate: func() any { return &schema.Interaction{} },
	_const.EventInteractionUpdate: func() any { return &schema.Interaction{} },
	_const.EventInteractionDelete: func() any { return &schema.Interaction{} },

	_const.EventGuildCreate:       func() any { return &schema.Guild{} },
	_const.EventGuildMemberAdd:    func() any { return &schema.GuildMemberAdd{} },
	_const.EventGuildMemberRemove: func() any { return &schema.GuildMemberRemove{} },
	_const.EventGuildBanAdd:       func() any { return &schema.GuildBanAdd{} },
	_const.EventGuildBanRemove:    func() any { return &schema.GuildBanRemove{} },
	_const.EventGuildEmojisUpdate: func() any { return &schema.GuildEmojisUpdate{} },
	_const.EventGuildUpdate:       func() any { return &schema.Guild{} },

	_const.EventMessageReactionAdd:         func() any { return &schema.MessageReactionAdd{} },
	_const.EventMessageReactionRemove:      func() any { return &schema.MessageReactionRemove{} },
	_const.EventMessageReactionRemoveAll:   func() any { return &schema.MessageReactionRemoveAll{} },
	_const.EventMessageReactionRemoveEmoji: func() any { return &schema.MessageReactionRemoveEmoji{} },

	_const.EventChannelCreate: func() any { return &schema.Channel{} },
	_const.EventChannelUpdate: func() any { return &schema.Channel{} },
	_const.EventChannelDelete: func() any { return &schema.ChannelDelete{} },
}

// return a factory
func Get(s string) func() any {
	return features[s]
}

// adds a factory to the missing K
func Add(s string, t func() any) {
	r := features[s]
	if r != nil {

		return
	}
	features[s] = t
}

// replaces the fabric of the existing K
func Change(s string, t func() any) {
	r := features[s]
	if r == nil {
		return
	}
	features[s] = t
}
