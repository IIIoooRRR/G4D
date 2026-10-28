package ui

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
)

type button struct {
	Type     int           `json:"type"` // 2
	Label    string        `json:"label,omitempty"`
	CustomID string        `json:"custom_id,omitempty"`
	Stl      int           `json:"style,omitempty"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
	Emoji    *entity.Emoji `json:"emoji,omitempty"`
}

func (b *button) isComponent() {}
func (b *button) isAccessory() {}
func (b *button) isUiV1() {

}
func Button(customId string) *button {
	return &button{
		Type:     2,
		CustomID: customId,
	}
}
func (b *button) Style(style int) *button {
	b.Stl = style
	return b
}
func (b *button) SetURL(url string) *button {
	b.URL = url
	return b
}
func (b *button) Disable(disabled bool) *button {
	b.Disabled = disabled
	return b
}
func (b *button) SetLabel(label string) *button {
	b.Label = label
	return b
}
func (b *button) SetEmoji(emoji *entity.Emoji) *button {
	b.Emoji = emoji
	return b
}
func (b *button) IsLink() bool {
	return b.Stl == 5
}

func (b *button) IsInteractive() bool {
	return b.Stl != 5
}

func LinkButton(label, url string) *button {
	return &button{
		Type:  2,
		Stl:   _const.ButtonStyleLink,
		Label: label,
		URL:   url,
	}
}

func DangerButton(customID, label string) *button {
	return &button{
		Type:     2,
		Stl:      _const.ButtonStyleDanger,
		CustomID: customID,
		Label:    label,
	}
}

func SuccessButton(customID, label string) *button {
	return &button{
		Type:     2,
		Stl:      _const.ButtonStyleSuccess,
		CustomID: customID,
		Label:    label,
	}
}

func PrimaryButton(customID, label string) *button {
	return &button{
		Type:     2,
		Stl:      _const.ButtonStylePrimary,
		CustomID: customID,
		Label:    label,
	}
}

func SecondaryButton(customID, label string) *button {
	return &button{
		Type:     2,
		Stl:      _const.ButtonStyleSecondary,
		CustomID: customID,
		Label:    label,
	}
}
