package updates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/updates/criterias"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials string) Builder
	WithCriteria(criteria criterias.Criteria) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Hash() hash.Hash
	Credentials() string
	Criteria() criterias.Criteria
}
