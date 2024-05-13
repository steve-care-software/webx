package updates

import "github.com/steve-care-software/datastencil/domain/hash"

// Update represents a database update
type Update interface {
	Hash() hash.Hash
	Origin() string
	Update() string
}
