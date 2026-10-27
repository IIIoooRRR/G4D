package uiv2

import (
	"strconv"
	"strings"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity/ui"
)

type Container struct {
	Type        int            `json:"type"` // 17
	Components  []ui.Component `json:"components"`
	AccentColor int            `json:"accent_color,omitempty"`
	Spoiler     bool           `json:"spoiler,omitempty"`
}

func (c *Container) IsComponent() {}

func NewContainer(components ...ui.Component) *Container {
	return &Container{Type: _const.ComponentTypeContainer, Components: components}
}

func (c *Container) AddComponents(components ...ui.Component) *Container {
	c.Components = append(c.Components, components...)
	return c
}

func (c *Container) SetAccentColor(hex string) *Container {
	hex = strings.ReplaceAll(hex, "#", "")
	if color, err := strconv.ParseInt(hex, 16, 32); err == nil {
		c.AccentColor = int(color)
	}
	return c
}

func (c *Container) SetSpoiler(spoiler bool) *Container {
	c.Spoiler = spoiler
	return c
}
