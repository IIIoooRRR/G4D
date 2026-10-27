package uiv2

import "github.com/IIIoooRRR/G4D/model/_const"

type MediaGallery struct {
	Type  int                `json:"type"` // 12
	Items []MediaGalleryItem `json:"items"`
}

type MediaGalleryItem struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	Spoiler     bool   `json:"spoiler,omitempty"`
}

func (m *MediaGallery) IsComponent() {}

func NewMediaGallery(items ...MediaGalleryItem) *MediaGallery {
	return &MediaGallery{Type: _const.ComponentTypeMediaGallery, Items: items}
}

func (m *MediaGallery) AddItem(url, description string) *MediaGallery {
	m.Items = append(m.Items, MediaGalleryItem{URL: url, Description: description})
	return m
}
