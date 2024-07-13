package signatures

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/historydb/domain/hash"
)

type signature struct {
	hash         hash.Hash
	isGenPrivKey bool
	fetchPubKey  string
	sign         signs.Sign
	vote         votes.Vote
}

func createSignatureWithGeneratePrivateKey(
	hash hash.Hash,
) Signature {
	return createSignatureInternally(hash, true, "", nil, nil)
}

func createSignatureWithFetchPublicKey(
	hash hash.Hash,
	fetchPubKey string,
) Signature {
	return createSignatureInternally(hash, false, fetchPubKey, nil, nil)
}

func createSignatureWithSign(
	hash hash.Hash,
	sign signs.Sign,
) Signature {
	return createSignatureInternally(hash, false, "", sign, nil)
}

func createSignatureWithVote(
	hash hash.Hash,
	vote votes.Vote,
) Signature {
	return createSignatureInternally(hash, false, "", nil, vote)
}

func createSignatureInternally(
	hash hash.Hash,
	isGenPrivKey bool,
	fetchPubKey string,
	sign signs.Sign,
	vote votes.Vote,
) Signature {
	out := signature{
		hash:         hash,
		isGenPrivKey: isGenPrivKey,
		fetchPubKey:  fetchPubKey,
		sign:         sign,
		vote:         vote,
	}

	return &out
}

// Hash returns the hash
func (obj *signature) Hash() hash.Hash {
	return obj.hash
}

// IsGeneratePrivateKey returns true if generatePrivateKey, false otherwise
func (obj *signature) IsGeneratePrivateKey() bool {
	return obj.isGenPrivKey
}

// IsFetchPublicKey returns true if fetchPubKey, false otherwise
func (obj *signature) IsFetchPublicKey() bool {
	return obj.fetchPubKey != ""
}

// FetchPublicKey returns true if fetchPubKey, false otherwise
func (obj *signature) FetchPublicKey() string {
	return obj.fetchPubKey
}

// IsSign returns true if sign, false otherwise
func (obj *signature) IsSign() bool {
	return obj.sign != nil
}

// Sign returns true if sign, false otherwise
func (obj *signature) Sign() signs.Sign {
	return obj.sign
}

// IsVote returns true if vote, false otherwise
func (obj *signature) IsVote() bool {
	return obj.vote != nil
}

// Vote returns true if vote, false otherwise
func (obj *signature) Vote() votes.Vote {
	return obj.vote
}
