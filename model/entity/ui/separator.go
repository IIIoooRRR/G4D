package ui

import "github.com/IIIoooRRR/G4D/model/_const"

type separator struct {
	Type  int  `json:"type"`              // 14
	Spcng int  `json:"spacing,omitempty"` // 1 = small, 2 = large
	Dvdr  bool `json:"divider,omitempty"`
}

func (s *separator) isComponent() {}

func Separator() *separator {
	return &separator{Type: _const.ComponentTypeSeparator, Spcng: 1}
}

func (s *separator) Spacing(spacing int) *separator {
	s.Spcng = spacing
	return s
}

func (s *separator) Divider(divider bool) *separator {
	s.Dvdr = divider
	return s
}
