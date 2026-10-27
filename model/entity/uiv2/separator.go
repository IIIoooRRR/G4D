package uiv2

import "github.com/IIIoooRRR/G4D/model/_const"

type Separator struct {
	Type    int  `json:"type"`              // 14
	Spacing int  `json:"spacing,omitempty"` // 1 = small, 2 = large
	Divider bool `json:"divider,omitempty"`
}

func (s *Separator) IsComponent() {}

func NewSeparator() *Separator {
	return &Separator{Type: _const.ComponentTypeSeparator, Spacing: 1}
}

func (s *Separator) SetSpacing(spacing int) *Separator {
	s.Spacing = spacing
	return s
}

func (s *Separator) SetDivider(divider bool) *Separator {
	s.Divider = divider
	return s
}
