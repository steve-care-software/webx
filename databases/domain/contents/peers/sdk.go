package peers

import "net/url"

// Adapter represents a peers adapter
type Adapter interface {
	ToContent(peers []*url.URL) ([]byte, error)
	ToPeers(content []byte) ([]*url.URL, error)
}
