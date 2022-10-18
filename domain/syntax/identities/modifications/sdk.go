package modifications

import (
	"time"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/signatures"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewModificationBuilder creates a new modification builder
func NewModificationBuilder() ModificationBuilder {
	return createModificationBuilder()
}

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithList(list []Modification) Builder
	Now() (Modifications, error)
}

// Modifications represents modifications
type Modifications interface {
	List() []Modification
	First() Modification
}

// ModificationBuilder represents a modification builder
type ModificationBuilder interface {
	Create() ModificationBuilder
	WithName(name string) ModificationBuilder
	WithSignature(signature signatures.PrivateKey) ModificationBuilder
	WithEncryption(encryption keys.PrivateKey) ModificationBuilder
	CreatedOn(createdOn time.Time) ModificationBuilder
	Now() (Modification, error)
}

// Modification represents a modifucation
type Modification interface {
	Content() Content
	CreatedOn() time.Time
}

// Content represents a modification content
type Content interface {
	HasName() bool
	Name() string
	HasSignature() bool
	Signature() signatures.PrivateKey
	HasEncryption() bool
	Encryption() keys.PrivateKey
}
