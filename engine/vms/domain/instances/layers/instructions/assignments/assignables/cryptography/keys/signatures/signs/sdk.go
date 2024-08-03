package signs

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the sign adapter
type Adapter interface {
	ToBytes(ins Sign) ([]byte, error)
	ToInstance(bytes []byte) (Sign, error)
}

// Builder represents a sign builder
type Builder interface {
	Create() Builder
	WithCreate(create creates.Create) Builder
	WithValidate(validate validates.Validate) Builder
	Now() (Sign, error)
}

// Sign represents a sign
type Sign interface {
	Hash() hash.Hash
	IsCreate() bool
	Create() creates.Create
	IsValidate() bool
	Validate() validates.Validate
}
