package ui

import (
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

type SelectMenu struct {
	Type        int            `json:"type"` // 3
	CustomID    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   int            `json:"min_values,omitempty"`
	MaxValues   int            `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`
}

type SelectOption struct {
	Label       string              `json:"label"`
	Value       string              `json:"value"`
	Description string              `json:"description,omitempty"`
	Emoji       *dependencies.Emoji `json:"emoji,omitempty"`
	Default     bool                `json:"default,omitempty"`
}

func (s *SelectMenu) IsComponent() {}

func NewMenu(customId string) *SelectMenu {
	return &SelectMenu{
		Type:     3,
		CustomID: customId,
	}
}
func NewSelectOption(label, value string) *SelectOption {
	return &SelectOption{
		Label: label,
		Value: value,
	}
}
func (op SelectOption) SetDescription(description string) SelectOption {
	op.Description = description
	return op
}
func (op SelectOption) SetEmoji(emoji *dependencies.Emoji) SelectOption {
	op.Emoji = emoji
	return op
}
func (op SelectOption) SetDefault(defaultValue bool) SelectOption {
	op.Default = defaultValue
	return op
}
