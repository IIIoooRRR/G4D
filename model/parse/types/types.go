package types

import (
	"reflect"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/shema"
)

var features = map[string]reflect.Type{
	_const.EventMessageCreate:     reflect.TypeOf(shema.GetMessage{}),
	_const.EventMessageUpdate:     reflect.TypeOf(shema.MessageEdit{}),
	_const.EventMessageDelete:     reflect.TypeOf(shema.MessageDelete{}),
	_const.EventMessageDeleteBulk: reflect.TypeOf(shema.MessageDeleteBulk{}),

	_const.EventInteractionCreate: reflect.TypeOf(shema.Interaction{}),
	_const.EventInteractionUpdate: reflect.TypeOf(shema.Interaction{}),
	_const.EventInteractionDelete: reflect.TypeOf(shema.Interaction{}),

	_const.EventGuildCreate:       reflect.TypeOf(shema.Guild{}),
	_const.EventGuildMemberAdd:    reflect.TypeOf(shema.GuildMemberAdd{}),
	_const.EventGuildMemberRemove: reflect.TypeOf(shema.GuildMemberRemove{}),
	_const.EventGuildBanAdd:       reflect.TypeOf(shema.GuildBanAdd{}),
	_const.EventGuildBanRemove:    reflect.TypeOf(shema.GuildBanRemove{}),
	_const.EventGuildEmojisUpdate: reflect.TypeOf(shema.GuildEmojisUpdate{}),
	_const.EventGuildUpdate:       reflect.TypeOf(shema.Guild{}),

	_const.EventMessageReactionAdd:         reflect.TypeOf(shema.MessageReactionAdd{}),
	_const.EventMessageReactionRemove:      reflect.TypeOf(shema.MessageReactionRemove{}),
	_const.EventMessageReactionRemoveAll:   reflect.TypeOf(shema.MessageReactionRemoveAll{}),
	_const.EventMessageReactionRemoveEmoji: reflect.TypeOf(shema.MessageReactionRemoveEmoji{}),

	_const.EventChannelCreate: reflect.TypeOf(shema.Channel{}),
	_const.EventChannelUpdate: reflect.TypeOf(shema.Channel{}),
	_const.EventChannelDelete: reflect.TypeOf(shema.ChannelDelete{}),
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
