package retrieves

import "github.com/steve-care-software/datastencil/domain/hash"

// Retrieve represents a retrieve
type Retrieve interface {
	Hash() hash.Hash
	IsList() bool
	IsExists() bool
	Exists() string
	IsRetrieve() bool
	Retrieve() string
}
