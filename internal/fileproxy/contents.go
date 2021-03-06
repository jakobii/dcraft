package fileproxy

import "io/ioutil"

// Conenter is a proxy to a files contents
type Conenter interface {
	Contents() ([]byte, error)
}

// NewConenter create a new file proxy
func NewConenter(path string) Conenter {
	return &contenter{
		Path: path,
	}
}

type contenter struct {
	Path string
}

func (c *contenter) Contents() ([]byte, error) {
	return ioutil.ReadFile(c.Path)
}

type mockContenter struct {
	content []byte
}

// NewMockContenter create a Conenter with a static value.
// useful for tests.
func NewMockContenter(b []byte) Conenter {
	return &mockContenter{b}
}

func (m *mockContenter) Contents() ([]byte, error) {
	b := make([]byte, 0, len(m.content))
	for _, x := range m.content {
		b = append(b, x)
	}
	return b, nil
}
