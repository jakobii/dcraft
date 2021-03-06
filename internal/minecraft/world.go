package minecraft

import "time"

// World details of a minecraft server
type World struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created"`
	Seed      string    `json:"seed"`
	Operators []string  `json:"operators"`
}
