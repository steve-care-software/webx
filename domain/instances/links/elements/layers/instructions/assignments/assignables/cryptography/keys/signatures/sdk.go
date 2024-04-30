package signatures

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
)

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
