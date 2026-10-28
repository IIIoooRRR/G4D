package ui

import "github.com/IIIoooRRR/G4D/model/_const"

type file struct {
	Type int    `json:"type"` // 13
	URL  string `json:"url"`
	Spl  bool   `json:"spoiler,omitempty"`
}

func (f *file) isComponent() {}

func File(url string) *file {
	return &file{Type: _const.ComponentTypeFile, URL: url}
}

func (f *file) Spoiler(spoiler bool) *file {
	f.Spl = spoiler
	return f
}
