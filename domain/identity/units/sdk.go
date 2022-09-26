package units

import (
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/units/genesis"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewUnitBuilder creates a new unit builder
func NewUnitBuilder() UnitBuilder {
	hashAdapter := hash.NewAdapter()
	signatureAdapter := signatures.NewRingSignatureAdapter()
	return createUnitBuilder(hashAdapter, signatureAdapter)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Builder represents units builder
type Builder interface {
	Create() Builder
	WithList(list []Unit) Builder
	Now() (Units, error)
}

// Units represents units
type Units interface {
	Hash() hash.Hash
	Amount() uint64
	List() []Unit
}

// UnitBuilder represents a unit builder
type UnitBuilder interface {
	Create() UnitBuilder
	WithContent(content Content) UnitBuilder
	WithSignatures(signatures []signatures.RingSignature) UnitBuilder
	Now() (Unit, error)
}

// Unit represents a unit
type Unit interface {
	Hash() hash.Hash
	Content() Content
	Signatures() []signatures.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithAmount(amount uint64) ContentBuilder
	WithOwner(owner []hash.Hash) ContentBuilder
	WithUnits(units Units) ContentBuilder
	WithGenesis(genesis genesis.Genesis) ContentBuilder
	Now() (Content, error)
}

// Content represents a unit content
type Content interface {
	Hash() hash.Hash
	Amount() uint64
	Owner() []hash.Hash
	Previous() Previous
}

// Previous represents a previous unit
type Previous interface {
	Hash() hash.Hash
	Amount() uint64
	IsUnits() bool
	Units() Units
	IsGenesis() bool
	Genesis() genesis.Genesis
}
