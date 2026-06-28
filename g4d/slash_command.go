package g4d

type SlashCreateCommand struct {
	Name        string               `json:"name"` // только маленькие буквы, без пробелов
	Description string               `json:"description"`
	Type        int                  `json:"type"`
	Options     []SlashCommandOption `json:"options,omitempty"`
}

type SlashCommandOption struct {
	Type        int    `json:"type"` // 3 - строка, 4 - число и т.д.
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}
