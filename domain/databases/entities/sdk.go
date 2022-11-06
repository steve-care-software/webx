package entities

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"go.dedis.ch/kyber/v3"
)

// Builder represents an entity builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier Identifier) Builder
	WithReference(reference Reference) Builder
	WithSignature(signature Signature) Builder
	Now() (Entity, error)
}

// Entity represents an entity
type Entity interface {
	Identifier() Identifier
	Reference() Reference
	Signature() Signature
}

// ReferenceBuilder represents the reference builder
type ReferenceBuilder interface {
	Create() ReferenceBuilder
	WithIdentifier(identifier Identifier) ReferenceBuilder
	WithTransaction(trx hash.Hash) ReferenceBuilder
	WithBlock(block hash.Hash) ReferenceBuilder
	WithChain(chain hash.Hash) ReferenceBuilder
	Now() (Reference, error)
}

// Reference represents the reference
type Reference interface {
	Identifier() Identifier
	Transaction() hash.Hash
	Block() hash.Hash
	Chain() hash.Hash
}

// IdentifiersBuilder represents identifiers builder
type IdentifiersBuilder interface {
	Create() IdentifiersBuilder
	WithList(list []Identifier) IdentifiersBuilder
	Now() (Identifiers, error)
}

// Identifiers represents identifiers
type Identifiers interface {
	List() []Identifier
}

// IdentifierBuilder represents an identifier builder
type IdentifierBuilder interface {
	Create() IdentifierBuilder
	WithSection(section uint) IdentifierBuilder
	WithElement(element uint) IdentifierBuilder
	Now() (Identifier, error)
}

// Identifier represents an identifier
type Identifier interface {
	Section() uint
	Element() uint
}

// SignatureBuilder represents a signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithRing(ring []hash.Hash) SignatureBuilder
	WithS(s kyber.Scalar) SignatureBuilder
	WithE(e kyber.Scalar) SignatureBuilder
	Now() (Signature, error)
}

// Signature represents a signature
type Signature interface {
	Ring() []hash.Hash
	S() kyber.Scalar
	E() kyber.Scalar
}
