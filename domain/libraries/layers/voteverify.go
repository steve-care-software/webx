package layers

import "github.com/steve-care-software/identity/domain/hash"

type voteVerify struct {
	hash       hash.Hash
	vote       string
	message    string
	hashedRing string
}

func createVoteVerify(
	hash hash.Hash,
	vote string,
	message string,
	hashedRing string,
) VoteVerify {
	out := voteVerify{
		hash:       hash,
		vote:       vote,
		message:    message,
		hashedRing: hashedRing,
	}

	return &out
}

// Hash returns the hash
func (obj *voteVerify) Hash() hash.Hash {
	return obj.hash
}

// Vote returns the vote
func (obj *voteVerify) Vote() string {
	return obj.vote
}

// Message returns the message
func (obj *voteVerify) Message() string {
	return obj.message
}

// HashedRing returns the hashed ring
func (obj *voteVerify) HashedRing() string {
	return obj.hashedRing
}
