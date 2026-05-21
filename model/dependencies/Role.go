package dependencies

type Role struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Color        int    `json:"color"`
	Hoist        bool   `json:"hoist"` // Отображается ли отдельно в списке
	Position     int    `json:"position"`
	Permissions  string `json:"permissions"`
	Managed      bool   `json:"managed"` // Создана ли интеграцией
	Mentionable  bool   `json:"mentionable"`
	UnicodeEmoji string `json:"unicode_emoji,omitempty"`
}
