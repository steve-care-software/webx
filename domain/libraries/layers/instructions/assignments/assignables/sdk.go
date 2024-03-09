package assignables

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/cryptography"
)

// NewBuilder creates a new assignable builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an assignable builder
type Builder interface {
	Create() Builder
	WithBytes(bytes bytes.Bytes) Builder
	WithConsant(constant constants.Constant) Builder
	WithAccount(account accounts.Account) Builder
	WithCryptography(cryptography cryptography.Cryptography) Builder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() bytes.Bytes
	IsConstant() bool
	Constant() constants.Constant
	IsAccount() bool
	Account() accounts.Account
	IsCryptography() bool
	Cryptography() cryptography.Cryptography
}
