package dependencies

import "time"

type RetryResponse struct {
	After time.Duration `json:"retry_after"`
}
