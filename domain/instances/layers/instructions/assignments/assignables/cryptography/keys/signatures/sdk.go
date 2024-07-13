package signatures

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the signature adapter
type Adapter interface {
	ToBytes(ins Signature) ([]byte, error)
	ToInstance(bytes []byte) (Signature, error)
}

// Builder represents the signature builder
type Builder interface {
	Create() Builder
	WithFetchPublicKey(fetchPubKey string) Builder
	WithSign(sign signs.Sign) Builder
	WithVote(vote votes.Vote) Builder
	IsGeneratePrivateKey() Builder
	Now() (Signature, error)
}

// Signature represents signature
type Signature interface {
	Hash() hash.Hash
	IsGeneratePrivateKey() bool
	IsFetchPublicKey() bool
	FetchPublicKey() string
	IsSign() bool
	Sign() signs.Sign
	IsVote() bool
	Vote() votes.Vote
}
