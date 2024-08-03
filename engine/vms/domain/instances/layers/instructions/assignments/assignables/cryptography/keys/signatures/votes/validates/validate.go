package validates

import "github.com/steve-care-software/webx/engine/hashes/domain/hash"

type validate struct {
	hash       hash.Hash
	vote       string
	message    string
	hashedRing string
}

func createValidate(
	hash hash.Hash,
	vote string,
	message string,
	hashedRing string,
) Validate {
	out := validate{
		hash:       hash,
		vote:       vote,
		message:    message,
		hashedRing: hashedRing,
	}

	return &out
}

// Hash returns the hash
func (obj *validate) Hash() hash.Hash {
	return obj.hash
}

// Vote returns the vote
func (obj *validate) Vote() string {
	return obj.vote
}

// Message returns the message
func (obj *validate) Message() string {
	return obj.message
}

// HashedRing returns the hashed ring
func (obj *validate) HashedRing() string {
	return obj.hashedRing
}
