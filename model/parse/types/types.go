package types

import (
	"reflect"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/schema"
)

var features = map[string]reflect.Type{
	_const.EventMessageCreate:     reflect.TypeOf(schema.GetMessage{}),
	_const.EventMessageUpdate:     reflect.TypeOf(schema.MessageEdit{}),
	_const.EventMessageDelete:     reflect.TypeOf(schema.MessageDelete{}),
	_const.EventMessageDeleteBulk: reflect.TypeOf(schema.MessageDeleteBulk{}),

	_const.EventInteractionCreate: reflect.TypeOf(schema.Interaction{}),
	_const.EventInteractionUpdate: reflect.TypeOf(schema.Interaction{}),
	_const.EventInteractionDelete: reflect.TypeOf(schema.Interaction{}),

	_const.EventGuildCreate:       reflect.TypeOf(schema.Guild{}),
	_const.EventGuildMemberAdd:    reflect.TypeOf(schema.GuildMemberAdd{}),
	_const.EventGuildMemberRemove: reflect.TypeOf(schema.GuildMemberRemove{}),
	_const.EventGuildBanAdd:       reflect.TypeOf(schema.GuildBanAdd{}),
	_const.EventGuildBanRemove:    reflect.TypeOf(schema.GuildBanRemove{}),
	_const.EventGuildEmojisUpdate: reflect.TypeOf(schema.GuildEmojisUpdate{}),
	_const.EventGuildUpdate:       reflect.TypeOf(schema.Guild{}),

	_const.EventMessageReactionAdd:         reflect.TypeOf(schema.MessageReactionAdd{}),
	_const.EventMessageReactionRemove:      reflect.TypeOf(schema.MessageReactionRemove{}),
	_const.EventMessageReactionRemoveAll:   reflect.TypeOf(schema.MessageReactionRemoveAll{}),
	_const.EventMessageReactionRemoveEmoji: reflect.TypeOf(schema.MessageReactionRemoveEmoji{}),

	_const.EventChannelCreate: reflect.TypeOf(schema.Channel{}),
	_const.EventChannelUpdate: reflect.TypeOf(schema.Channel{}),
	_const.EventChannelDelete: reflect.TypeOf(schema.ChannelDelete{}),
}

func Get(s string) reflect.Type {
	return features[s]
}
func Add(s string, t reflect.Type) {
	r := features[s]
	if r != nil {

		return
	}
	features[s] = t
}
func Change(s string, t reflect.Type) {
	r := features[s]
	if r == nil {
		return
	}
	features[s] = t
}
