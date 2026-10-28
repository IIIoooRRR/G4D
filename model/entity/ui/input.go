package ui

type textInput struct {
	Type     int    `json:"type"` // 4
	CustomID string `json:"custom_id"`
	Label    string `json:"label"`
	Style    int    `json:"style"` // 1=short, 2=paragraph
	Placehld string `json:"placeholder,omitempty"`
	Rqr      bool   `json:"required,omitempty"`
	Val      string `json:"value,omitempty"`
	MinLen   int    `json:"min_length,omitempty"`
	MaxLen   int    `json:"max_length,omitempty"`
}

func (t textInput) isModal() {}
func TextInput(customId, label string, style int) *textInput {
	return &textInput{
		Type:     4,
		CustomID: customId,
		Label:    label,
		Style:    style,
	}
}
func (t *textInput) Placeholder(placeholder string) *textInput {
	t.Placehld = placeholder
	return t
}
func (t *textInput) Required(required bool) *textInput {
	t.Rqr = required
	return t
}
func (t *textInput) Value(value string) *textInput {
	t.Val = value
	return t
}
func (t *textInput) MinLength(minLength int) *textInput {
	t.MinLen = minLength
	return t
}
func (t *textInput) MaxLength(maxLength int) *textInput {
	t.MaxLen = maxLength
	return t
}
