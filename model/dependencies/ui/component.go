package ui

type Component interface {
	IsComponent() //use it for custom components that I haven't implemented.
}
type ActionRow struct {
	Type       int         `json:"type,omitempty"`
	Components []Component `json:"components,omitempty"`
}

func (ar *ActionRow) AddComponents(components ...Component) ActionRow {
	ar.Components = append(ar.Components, components...)
	return *ar
}
func NewActionRow() *ActionRow {
	return &ActionRow{
		Components: make([]Component, 0),
		Type:       1,
	}
}
