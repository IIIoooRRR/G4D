package ui

import (
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

type Button struct {
	Type     int                 `json:"type"` // 2
	Label    string              `json:"label,omitempty"`
	CustomID string              `json:"custom_id,omitempty"`
	Style    int                 `json:"style,omitempty"`
	URL      string              `json:"url,omitempty"`
	Disabled bool                `json:"disabled,omitempty"`
	Emoji    *dependencies.Emoji `json:"emoji,omitempty"`
}

func (b Button) IsComponent() {}

func NewButton(customId string) *Button {
	return &Button{
		Type:     2,
		CustomID: customId,
	}
}
func (b Button) SetStyle(style int) Button {
	b.Style = style
	return b
}
func (b Button) SetURL(url string) Button {
	b.URL = url
	return b
}
func (b Button) SetDisabled(disabled bool) Button {
	b.Disabled = disabled
	return b
}
func (b Button) SetLabel(label string) Button {
	b.Label = label
	return b
}
func (b Button) SetEmoji(emoji *dependencies.Emoji) Button {
	b.Emoji = emoji
	return b
}
