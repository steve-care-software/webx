package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

// Builder represents a references builder
type Builder interface {
	Create() Builder
	WithList(list []Reference) Builder
	Now() (References, error)
}

// References represents references
type References interface {
	Hash() hash.Hash
	List() []Reference
}

// ReferenceBuilder represents a reference builder
type ReferenceBuilder interface {
	Create() ReferenceBuilder
	WithVariable(variable string) ReferenceBuilder
	WithInstance(insatnce instances.Instance) ReferenceBuilder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Hash() hash.Hash
	Variable() string
	Instance() instances.Instance
}
