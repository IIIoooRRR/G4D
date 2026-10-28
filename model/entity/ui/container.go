package ui

import (
	"strconv"
	"strings"

	"github.com/IIIoooRRR/G4D/model/_const"
)

type container struct {
	Type      int         `json:"type"` // 17
	Cmpnts    []Component `json:"components"`
	AcntColor int         `json:"accent_color,omitempty"`
	Spoil     bool        `json:"spoiler,omitempty"`
}

func (c *container) isComponent()       {}
func (c *container) isComponentTopLvl() {}
func Container(components ...Component) *container {
	return &container{Type: _const.ComponentTypeContainer, Cmpnts: components}
}

func (c *container) Components(components ...Component) *container {
	c.Cmpnts = append(c.Cmpnts, components...)
	return c
}

func (c *container) AccentColor(hex string) *container {
	hex = strings.ReplaceAll(hex, "#", "")
	if color, err := strconv.ParseInt(hex, 16, 32); err == nil {
		c.AcntColor = int(color)
	}
	return c
}

func (c *container) Spoiler(spoiler bool) *container {
	c.Spoil = spoiler
	return c
}
