package votes

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// Builder represents a vote builder
type Builder interface {
	Create() Builder
	WithCreate(create creates.Create) Builder
	WithValidate(validate validates.Validate) Builder
	Now() (Vote, error)
}

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	IsCreate() bool
	Create() creates.Create
	IsValidate() bool
	Validate() validates.Validate
}
