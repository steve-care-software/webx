package layers

import "github.com/steve-care-software/identity/domain/hash"

type vote struct {
	hash    hash.Hash
	ring    string
	message string
}

func createVote(
	hash hash.Hash,
	ring string,
	message string,
) Vote {
	out := vote{
		hash:    hash,
		ring:    ring,
		message: message,
	}

	return &out
}

// Hash returns the hash
func (obj *vote) Hash() hash.Hash {
	return obj.hash
}

// Ring returns the ring
func (obj *vote) Ring() string {
	return obj.ring
}

// Message returns the message
func (obj *vote) Message() string {
	return obj.message
}
