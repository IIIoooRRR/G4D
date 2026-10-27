package uiv2

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity/ui"
)

type Section struct {
	Type       int            `json:"type"`                // 9
	Components []ui.Component `json:"components"`          // text display
	Accessory  ui.Component   `json:"accessory,omitempty"` // button or thumbnail
}

func (s *Section) IsComponent() {}

func NewSection(components ...ui.Component) *Section {
	return &Section{Type: _const.ComponentTypeSection, Components: components}
}

func (s *Section) AddText(text string) *Section {
	s.Components = append(s.Components, NewTextDisplay(text))
	return s
}

func (s *Section) SetAccessory(accessory ui.Component) *Section {
	s.Accessory = accessory
	return s
}
