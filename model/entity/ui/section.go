package ui

import (
	"github.com/IIIoooRRR/G4D/model/_const"
)

type section struct {
	Type       int            `json:"type"`                 // 9
	Components []*textDisplay `json:"components,omitempty"` // only textDisp
	Acs        Accessory      `json:"accessory,omitempty"`  // button or thumbnail
}

func (s *section) isComponent() {}

func Section(components ...*textDisplay) *section {
	return &section{Type: _const.ComponentTypeSection, Components: components}
}

func (s *section) Text(text string) *section {
	s.Components = append(s.Components, TextDisplay(text))
	return s
}

func (s *section) Accessory(accessory Accessory) *section {
	s.Acs = accessory
	return s
}
