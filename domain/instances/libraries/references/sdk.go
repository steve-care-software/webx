package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

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
	Retrieve(name string) (Reference, error)
}

// ReferenceBuilder represents an reference builder
type ReferenceBuilder interface {
	Create() ReferenceBuilder
	WithVariable(variale string) ReferenceBuilder
	WithIdentifier(identifier hash.Hash) ReferenceBuilder
	Now() (Reference, error)
}

// Reference represents an reference
type Reference interface {
	Hash() hash.Hash
	Variable() string
	Identifier() hash.Hash
}
