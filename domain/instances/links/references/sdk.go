package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewReferenceBuilder creates a new reference builder
func NewReferenceBuilder() ReferenceBuilder {
	hashAdapter := hash.NewAdapter()
	return createReferenceBuilder(
		hashAdapter,
	)
}

// Adapter represents the references adapter
type Adapter interface {
	ToBytes(ins References) ([]byte, error)
	ToInstance(bytes []byte) (References, error)
}

// Builder represents an references builder
type Builder interface {
	Create() Builder
	WithList(list []Reference) Builder
	Now() (References, error)
}

// References represents references
type References interface {
	Hash() hash.Hash
	List() []Reference
	Fetch(name string) (Reference, error)
}

// ReferenceBuilder represents an reference builder
type ReferenceBuilder interface {
	Create() ReferenceBuilder
	WithVariable(variable string) ReferenceBuilder
	WithIdentifier(identifier hash.Hash) ReferenceBuilder
	Now() (Reference, error)
}

// Reference represents an reference
type Reference interface {
	Hash() hash.Hash
	Variable() string
	Identifier() hash.Hash
}
