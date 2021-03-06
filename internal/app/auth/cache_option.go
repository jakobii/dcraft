package auth

// CacheOption is an option parameter
type CacheOption func(*Cache)

// WithSuper secret
func WithSuper(secret string) CacheOption {
	return func(c *Cache) {
		c.super = secret
	}
}
