package ui

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
)

type SelectMenu struct {
	Type        int            `json:"type"` // 3, 5, 6, 7, 8
	CustomID    string         `json:"custom_id"`
	Options     []SelectOption `json:"options,omitempty"` // only for TypeString
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   int            `json:"min_values,omitempty"`
	MaxValues   int            `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`

	// onlu for Channel Select (type 8)
	ChannelTypes []int `json:"channel_types,omitempty"`

	//for self‑filling forms only (type 5, 6, 7, 8)
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
}

func (s *SelectMenu) IsComponent() {}

type SelectOption struct {
	Label       string        `json:"label"`
	Value       string        `json:"value"`
	Description string        `json:"description,omitempty"`
	Emoji       *entity.Emoji `json:"emoji,omitempty"`
	Default     bool          `json:"default,omitempty"`
}

type SelectDefaultValue struct {
	ID   string `json:"id"`
	Type string `json:"type"` // "user", "role", "channel"
}

func NewMenu(customId string) *SelectMenu {
	return &SelectMenu{
		Type:     _const.ComponentSelectTypeString,
		CustomID: customId,
	}
}

func NewUserSelect(customId string) *SelectMenu {
	return &SelectMenu{
		Type:     _const.ComponentSelectTypeUser,
		CustomID: customId,
	}
}

func NewRoleSelect(customId string) *SelectMenu {
	return &SelectMenu{
		Type:     _const.ComponentSelectTypeRole,
		CustomID: customId,
	}
}

func NewMentionableSelect(customId string) *SelectMenu {
	return &SelectMenu{
		Type:     _const.ComponentSelectTypeMentionable,
		CustomID: customId,
	}
}

func NewChannelSelect(customId string) *SelectMenu {
	return &SelectMenu{
		Type:     _const.ComponentSelectTypeChannel,
		CustomID: customId,
	}
}

func NewSelectOption(label, value string) *SelectOption {
	return &SelectOption{
		Label: label,
		Value: value,
	}
}

/* options */

func (op *SelectOption) SetDescription(description string) *SelectOption {
	op.Description = description
	return op
}

func (op *SelectOption) SetEmoji(emoji *entity.Emoji) *SelectOption {
	op.Emoji = emoji
	return op
}

func (op *SelectOption) SetDefault(defaultValue bool) *SelectOption {
	op.Default = defaultValue
	return op
}

/* menu */

func (s *SelectMenu) AddOptions(options ...*SelectOption) *SelectMenu {
	for _, opt := range options {
		s.Options = append(s.Options, *opt)
	}
	return s
}

func (s *SelectMenu) SetPlaceholder(placeholder string) *SelectMenu {
	s.Placeholder = placeholder
	return s
}

func (s *SelectMenu) SetMinValues(minValues int) *SelectMenu {
	s.MinValues = minValues
	return s
}

func (s *SelectMenu) SetMaxValues(maxValues int) *SelectMenu {
	s.MaxValues = maxValues
	return s
}

func (s *SelectMenu) SetDisabled(disabled bool) *SelectMenu {
	s.Disabled = disabled
	return s
}

// only for Channel Select
func (s *SelectMenu) SetChannelTypes(types ...int) *SelectMenu {
	s.ChannelTypes = types
	return s
}

//for self‑filling forms only

func (s *SelectMenu) AddDefaultUser(userId string) *SelectMenu {
	s.DefaultValues = append(s.DefaultValues, SelectDefaultValue{
		ID:   userId,
		Type: "user",
	})
	return s
}

func (s *SelectMenu) AddDefaultRole(roleId string) *SelectMenu {
	s.DefaultValues = append(s.DefaultValues, SelectDefaultValue{
		ID:   roleId,
		Type: "role",
	})
	return s
}

func (s *SelectMenu) AddDefaultChannel(channelId string) *SelectMenu {
	s.DefaultValues = append(s.DefaultValues, SelectDefaultValue{
		ID:   channelId,
		Type: "channel",
	})
	return s
}
