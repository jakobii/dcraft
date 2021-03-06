package minecraft

import "github.com/google/uuid"

// User ...
type User struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
