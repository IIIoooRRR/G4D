package uiv2

import "github.com/IIIoooRRR/G4D/model/_const"

type File struct {
	Type    int    `json:"type"` // 13
	URL     string `json:"url"`
	Spoiler bool   `json:"spoiler,omitempty"`
}

func (f *File) IsComponent() {}

func NewFile(url string) *File {
	return &File{Type: _const.ComponentTypeFile, URL: url}
}

func (f *File) SetSpoiler(spoiler bool) *File {
	f.Spoiler = spoiler
	return f
}
