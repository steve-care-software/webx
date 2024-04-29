package communications

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/votes"
)

type communication struct {
	hash         hash.Hash
	sign         signs.Sign
	vote         votes.Vote
	generateRing string
}

func createCommunicationWithSign(
	hash hash.Hash,
	sign signs.Sign,
) Communication {
	return createCommunicationInternally(hash, sign, nil, "")
}

func createCommunicationWithVote(
	hash hash.Hash,
	vote votes.Vote,
) Communication {
	return createCommunicationInternally(hash, nil, vote, "")
}

func createCommunicationWithGenerateRing(
	hash hash.Hash,
	generateRing string,
) Communication {
	return createCommunicationInternally(hash, nil, nil, generateRing)
}

func createCommunicationInternally(
	hash hash.Hash,
	sign signs.Sign,
	vote votes.Vote,
	generateRing string,
) Communication {
	out := communication{
		hash:         hash,
		sign:         sign,
		vote:         vote,
		generateRing: generateRing,
	}

	return &out
}

// Hash returns the hash
func (obj *communication) Hash() hash.Hash {
	return obj.hash
}

// IsSign returns true if there is a sign, false otherwise
func (obj *communication) IsSign() bool {
	return obj.sign != nil
}

// Sign returns the sign, if any
func (obj *communication) Sign() signs.Sign {
	return obj.sign
}

// IsVote returns true if there is a vote, false otherwise
func (obj *communication) IsVote() bool {
	return obj.vote != nil
}

// Vote returns the vote, if any
func (obj *communication) Vote() votes.Vote {
	return obj.vote
}

// IsGenerateRing returns true if generateRing, false otherwise
func (obj *communication) IsGenerateRing() bool {
	return obj.generateRing != ""
}

// GenerateRing returns the generateRing, if any
func (obj *communication) GenerateRing() string {
	return obj.generateRing
}
