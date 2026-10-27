package uiv2

import "github.com/IIIoooRRR/G4D/model/_const"

type Thumbnail struct {
	Type        int    `json:"type"` // 11
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	Spoiler     bool   `json:"spoiler,omitempty"`
}

func (t *Thumbnail) IsComponent() {}

func NewThumbnail(url string) *Thumbnail {
	return &Thumbnail{Type: _const.ComponentTypeThumbnail, URL: url}
}

func (t *Thumbnail) SetDescription(desc string) *Thumbnail {
	t.Description = desc
	return t
}

func (t *Thumbnail) SetSpoiler(spoiler bool) *Thumbnail {
	t.Spoiler = spoiler
	return t
}
