package ui

type Component interface {
	isComponent() //use it for custom components that I haven't implemented.
}
type Accessory interface {
	Component
	isAccessory()
}
type UiV1 interface {
	Component
	isUiV1()
}
type Modal interface {
	isModal()
}
type TopLvlComponent interface {
	Component
	isComponentTopLvl()
}
type actionRow struct {
	Type   int    `json:"type,omitempty"`
	Cmpnts []UiV1 `json:"components,omitempty"`
}

func (ar *actionRow) Components(components ...UiV1) *actionRow {
	ar.Cmpnts = append(ar.Cmpnts, components...)
	return ar
}
func ActionRow(cmpnts ...UiV1) *actionRow {
	return &actionRow{
		Cmpnts: cmpnts,
		Type:   1,
	}
}
func (ar *actionRow) isComponent()       {}
func (ar *actionRow) isComponentTopLvl() {}
