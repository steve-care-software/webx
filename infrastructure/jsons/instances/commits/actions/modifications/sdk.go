package modifications

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithInsert(insert string) Builder
	WithUpdate(update string) Builder
	WithDelete(delete string) Builder
	Now() (Modification, error)
}

// Modification represents a modification
type Modification interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() string
	IsUpdate() bool
	Update() string
	IsDelete() bool
	Delete() string
}
