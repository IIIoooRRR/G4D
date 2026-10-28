package ui

import "github.com/IIIoooRRR/G4D/model/_const"

type thumbnail struct {
	Type  int    `json:"type"` // 11
	URL   string `json:"url"`
	Desc  string `json:"description,omitempty"`
	Spoil bool   `json:"spoiler,omitempty"`
}

func (t *thumbnail) isComponent() {}
func (t *thumbnail) isAccessory() {}
func Thumbnail(url string) *thumbnail {
	return &thumbnail{Type: _const.ComponentTypeThumbnail, URL: url}
}

func (t *thumbnail) Description(desc string) *thumbnail {
	t.Desc = desc
	return t
}

func (t *thumbnail) Spoiler(spoiler bool) *thumbnail {
	t.Spoil = spoiler
	return t
}
