package g4d

type SlashCreateCommand struct {
	Name        string               `json:"name"` // only little litter
	Description string               `json:"description"`
	Type        int                  `json:"type"`
	Options     []SlashCommandOption `json:"options,omitempty"`
}

type SlashCommandOption struct {
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}
