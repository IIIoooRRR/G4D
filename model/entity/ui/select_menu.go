package ui

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
)

type selectMenu struct {
	Type     int             `json:"type"` // 3, 5, 6, 7, 8
	Id       string          `json:"custom_id"`
	Opts     []*selectOption `json:"options,omitempty"` // only for TypeString
	Placehld string          `json:"placeholder,omitempty"`
	MinVal   *int            `json:"min_values,omitempty"`
	MaxVal   *int            `json:"max_values,omitempty"`
	Disabled bool            `json:"disabled,omitempty"`

	// onlu for Channel Select (type 8)
	ChannelTps []int `json:"channel_types,omitempty"`

	//for self‑filling forms only (type 5, 6, 7, 8)
	DefaultValues []selectDefaultValue `json:"default_values,omitempty"`
}

func (s *selectMenu) isComponent() {}
func (s *selectMenu) isUiV1()      {}
func (s *selectMenu) isModal()     {}

type selectOption struct {
	Label string        `json:"label"`
	Val   string        `json:"value"`
	Desc  string        `json:"description,omitempty"`
	Emj   *entity.Emoji `json:"emoji,omitempty"`
	Deflt bool          `json:"default,omitempty"`
}

type selectDefaultValue struct {
	ID   string `json:"id"`
	Type string `json:"type"` // "user", "role", "channel"
}

func SelectMenu(customId string) *selectMenu {
	return &selectMenu{
		Type: _const.ComponentSelectTypeString,
		Id:   customId,
	}
}

func UserSelectMenu(customId string) *selectMenu {
	return &selectMenu{
		Type: _const.ComponentSelectTypeUser,
		Id:   customId,
	}
}

func RoleSelectMenu(customId string) *selectMenu {
	return &selectMenu{
		Type: _const.ComponentSelectTypeRole,
		Id:   customId,
	}
}

func MentionableSelectMenu(customId string) *selectMenu {
	return &selectMenu{
		Type: _const.ComponentSelectTypeMentionable,
		Id:   customId,
	}
}

func ChannelSelectMenu(customId string) *selectMenu {
	return &selectMenu{
		Type: _const.ComponentSelectTypeChannel,
		Id:   customId,
	}
}

func SelectOption(label, value string) *selectOption {
	return &selectOption{
		Label: label,
		Val:   value,
	}
}

/* options */

func (op *selectOption) Description(description string) *selectOption {
	op.Desc = description
	return op
}

func (op *selectOption) Emoji(emoji *entity.Emoji) *selectOption {
	op.Emj = emoji
	return op
}

func (op *selectOption) Default(defaultValue bool) *selectOption {
	op.Deflt = defaultValue
	return op
}

/* menu */

func (s *selectMenu) Options(options ...*selectOption) *selectMenu {
	s.Opts = append(s.Opts, options...)
	return s
}

func (s *selectMenu) Placeholder(placeholder string) *selectMenu {
	s.Placehld = placeholder
	return s
}

func (s *selectMenu) MinValues(minValues int) *selectMenu {
	s.MinVal = &minValues
	return s
}

func (s *selectMenu) MaxValues(maxValues int) *selectMenu {
	s.MaxVal = &maxValues
	return s
}

func (s *selectMenu) Disable(disabled bool) *selectMenu {
	s.Disabled = disabled
	return s
}

// only for Channel Select
func (s *selectMenu) ChannelTypes(types ...int) *selectMenu {
	s.ChannelTps = types
	return s
}

//for self‑filling forms only

func (s *selectMenu) DefaultUser(userId string) *selectMenu {
	s.DefaultValues = append(s.DefaultValues, selectDefaultValue{
		ID:   userId,
		Type: "user",
	})
	return s
}

func (s *selectMenu) DefaultRole(roleId string) *selectMenu {
	s.DefaultValues = append(s.DefaultValues, selectDefaultValue{
		ID:   roleId,
		Type: "role",
	})
	return s
}

func (s *selectMenu) DefaultChannel(channelId string) *selectMenu {
	s.DefaultValues = append(s.DefaultValues, selectDefaultValue{
		ID:   channelId,
		Type: "channel",
	})
	return s
}
