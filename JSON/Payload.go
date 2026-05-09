package JSON

import (
	"encoding/json"
)

type Payload struct {
	Op int             `json:"op"`          // Код операции (0, 1, 2, 7, 8, 9, 10, 11)
	D  json.RawMessage `json:"d"`           // Сами данные (сырой JSON)
	S  int             `json:"s,omitempty"` // Порядковый номер (нужен для Resume!)
	T  string          `json:"t,omitempty"` // Название события (напр. MESSAGE_CREATE)
}
