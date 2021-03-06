package app

import (
	"context"

	"github.com/jakobii/dcraft/internal/minecraft"
)

// Whitelister is a wrapper interface for managing whitelist
type Whitelister interface {
	WhitelistGetter
}

// WhitelistGetter gets users in a whitelist
type WhitelistGetter interface {
	Get(ctx context.Context) ([]minecraft.User, error)
}
