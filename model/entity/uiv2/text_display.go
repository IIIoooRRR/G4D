package uiv2

import "github.com/IIIoooRRR/G4D/model/_const"

type TextDisplay struct {
	Type    int    `json:"type"` // 10
	Content string `json:"content"`
}

func (t *TextDisplay) IsComponent() {}

func NewTextDisplay(content string) *TextDisplay {
	return &TextDisplay{Type: _const.ComponentTypeTextDisplay, Content: content}
}

func (t *TextDisplay) SetContent(content string) *TextDisplay {
	t.Content = content
	return t
}
