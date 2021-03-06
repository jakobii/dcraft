package whitelist

import (
	"context"
	"encoding/json"

	"github.com/jakobii/dcraft/internal/app"
	"github.com/jakobii/dcraft/internal/fileproxy"
	"github.com/jakobii/dcraft/internal/minecraft"
)

// Whitelist manages minecraft server whitelists.
type Whitelist struct {
	fileproxy.Conenter
}

// New create a new whitelist.
// path to whitelist.json
func New(path string) app.Whitelister {
	return &Whitelist{
		Conenter: fileproxy.NewConenter(path),
	}
}

// Get whitelisted users.
func (w *Whitelist) Get(ctx context.Context) ([]minecraft.User, error) {
	b, err := w.Contents()
	if err != nil {
		return nil, err
	}
	var users []minecraft.User
	return users, json.Unmarshal(b, &users)
}
