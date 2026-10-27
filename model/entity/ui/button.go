package ui

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
)

type Button struct {
	Type     int           `json:"type"` // 2
	Label    string        `json:"label,omitempty"`
	CustomID string        `json:"custom_id,omitempty"`
	Style    int           `json:"style,omitempty"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
	Emoji    *entity.Emoji `json:"emoji,omitempty"`
}

func (b Button) IsComponent() {}

func NewButton(customId string) *Button {
	return &Button{
		Type:     2,
		CustomID: customId,
	}
}
func (b *Button) SetStyle(style int) *Button {
	b.Style = style
	return b
}
func (b *Button) SetURL(url string) *Button {
	b.URL = url
	return b
}
func (b *Button) SetDisabled(disabled bool) *Button {
	b.Disabled = disabled
	return b
}
func (b *Button) SetLabel(label string) *Button {
	b.Label = label
	return b
}
func (b *Button) SetEmoji(emoji *entity.Emoji) *Button {
	b.Emoji = emoji
	return b
}
func (b *Button) IsLink() bool {
	return b.Style == 5
}

func (b *Button) IsInteractive() bool {
	return b.Style != 5
}

func NewLinkButton(label, url string) *Button {
	return &Button{
		Type:  2,
		Style: _const.ButtonStyleLink,
		Label: label,
		URL:   url,
	}
}

func NewDangerButton(customID, label string) *Button {
	return &Button{
		Type:     2,
		Style:    _const.ButtonStyleDanger,
		CustomID: customID,
		Label:    label,
	}
}

func NewSuccessButton(customID, label string) *Button {
	return &Button{
		Type:     2,
		Style:    _const.ButtonStyleSuccess,
		CustomID: customID,
		Label:    label,
	}
}

func NewPrimaryButton(customID, label string) *Button {
	return &Button{
		Type:     2,
		Style:    _const.ButtonStylePrimary,
		CustomID: customID,
		Label:    label,
	}
}

func NewSecondaryButton(customID, label string) *Button {
	return &Button{
		Type:     2,
		Style:    _const.ButtonStyleSecondary,
		CustomID: customID,
		Label:    label,
	}
}
