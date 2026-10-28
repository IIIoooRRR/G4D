package ui

type label struct {
	Type  int    `json:"type"`                  // 18
	Label string `json:"label"`                 // header, max 45smbl
	Desc  string `json:"description,omitempty"` // max 100 smbl
	Cmpnt Modal  `json:"component"`             // Text Input or String Select
}

func (l *label) isComponent()       {}
func (l *label) isComponentTopLvl() {}
func Label(lbl string, child Modal) *label {
	return &label{
		Type:  18,
		Label: lbl,
		Cmpnt: child,
	}
}

func (l *label) Description(desc string) *label {
	l.Desc = desc
	return l
}
