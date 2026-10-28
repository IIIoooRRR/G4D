package entity

type Role struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Color        int               `json:"color"`
	Hoist        bool              `json:"hoist"`
	Position     int               `json:"position"`
	Permissions  string            `json:"permissions"`
	Managed      bool              `json:"managed"`
	Mentionable  bool              `json:"mentionable"`
	Icon         *string           `json:"icon,omitempty"`
	UnicodeEmoji *string           `json:"unicode_emoji,omitempty"`
	Flags        int               `json:"flags,omitempty"`
	Version      int               `json:"version,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`
	Colors       *RoleColors       `json:"colors,omitempty"`
}

type RoleColors struct {
	PrimaryColor   int  `json:"primary_color"`
	SecondaryColor *int `json:"secondary_color"`
	TertiaryColor  *int `json:"tertiary_color"`
}
