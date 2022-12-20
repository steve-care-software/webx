package peers

import "net/url"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// Adapter represents a peers adapter
type Adapter interface {
	ToContent(peers []*url.URL) ([]byte, error)
	ToPeers(content []byte) ([]*url.URL, error)
}
