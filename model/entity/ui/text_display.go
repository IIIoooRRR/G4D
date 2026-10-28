package ui

import "github.com/IIIoooRRR/G4D/model/_const"

type textDisplay struct {
	Type int    `json:"type"` // 10
	Cnt  string `json:"content"`
}

func (t *textDisplay) isComponent() {}

func TextDisplay(content string) *textDisplay {
	return &textDisplay{Type: _const.ComponentTypeTextDisplay, Cnt: content}
}

func (t *textDisplay) Content(content string) *textDisplay {
	t.Cnt = content
	return t
}
