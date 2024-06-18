package signatures

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
)

// NewSignatureWithGeneratePrivateKeyForTests creates a new signature with generate private key for tests
func NewSignatureWithGeneratePrivateKeyForTests() Signature {
	ins, err := NewBuilder().Create().IsGeneratePrivateKey().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignatureWithFetchPublicKeyForTests creates a new signature with fetchPublicKey for tests
func NewSignatureWithFetchPublicKeyForTests(fetchPublicKey string) Signature {
	ins, err := NewBuilder().Create().WithFetchPublicKey(fetchPublicKey).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignatureWithSignForTests creates a new signature with sign for tests
func NewSignatureWithSignForTests(sign signs.Sign) Signature {
	ins, err := NewBuilder().Create().WithSign(sign).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignatureWithVoteForTests creates a new signature with vote for tests
func NewSignatureWithVoteForTests(vote votes.Vote) Signature {
	ins, err := NewBuilder().Create().WithVote(vote).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
