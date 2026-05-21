package ui

type TextInput struct {
	Type        int    `json:"type"` // 4
	CustomID    string `json:"custom_id"`
	Label       string `json:"label"`
	Style       int    `json:"style"` // 1=short, 2=paragraph
	Placeholder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Value       string `json:"value,omitempty"`
	MinLength   int    `json:"min_length,omitempty"`
	MaxLength   int    `json:"max_length,omitempty"`
}

func (t TextInput) IsComponent() {}

func NewTextInput(customId, label string, style int) TextInput {
	return TextInput{
		Type:     4,
		CustomID: customId,
		Label:    label,
		Style:    style,
	}
}
func (t TextInput) SetPlaceholder(placeholder string) TextInput {
	t.Placeholder = placeholder
	return t
}
func (t TextInput) SetRequired(required bool) TextInput {
	t.Required = required
	return t
}
func (t TextInput) SetValue(value string) TextInput {
	t.Value = value
	return t
}
func (t TextInput) SetMinLength(minLength int) TextInput {
	t.MinLength = minLength
	return t
}
func (t TextInput) SetMaxLength(maxLength int) TextInput {
	t.MaxLength = maxLength
	return t
}
