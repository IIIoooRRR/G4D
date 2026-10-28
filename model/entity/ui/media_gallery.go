package ui

import "github.com/IIIoooRRR/G4D/model/_const"

type mediaGallery struct {
	Type  int                 `json:"type"` // 12
	Items []*mediaGalleryItem `json:"items"`
}

type mediaGalleryItem struct {
	Med   media  `json:"media"`
	Desc  string `json:"description,omitempty"`
	Spoil bool   `json:"spoiler,omitempty"`
}

type media struct {
	URL string `json:"url"`
}

func (m *mediaGallery) isComponent() {}

func MediaGallery(items ...*mediaGalleryItem) *mediaGallery {
	return &mediaGallery{Type: _const.ComponentTypeMediaGallery, Items: items}
}
func (m *mediaGallery) Add(items ...*mediaGalleryItem) *mediaGallery {
	m.Items = append(m.Items, items...)
	return m
}

func MediaItem(url string) *mediaGalleryItem {
	return &mediaGalleryItem{Med: media{URL: url}}
}
func (mi *mediaGalleryItem) Description(description string) *mediaGalleryItem {
	mi.Desc = description
	return mi
}
func (mi *mediaGalleryItem) Spoiler(spoiler bool) *mediaGalleryItem {
	mi.Spoil = spoiler
	return mi
}
