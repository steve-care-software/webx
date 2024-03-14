package votes

import "github.com/steve-care-software/datastencil/domain/hash"

type vote struct {
	hash    hash.Hash
	message string
	ring    string
	account string
}

func createVote(
	hash hash.Hash,
	message string,
	ring string,
	account string,
) Vote {
	out := vote{
		hash:    hash,
		message: message,
		ring:    ring,
		account: account,
	}

	return &out
}

// Hash returns the hash
func (obj *vote) Hash() hash.Hash {
	return obj.hash
}

// Message returns the message
func (obj *vote) Message() string {
	return obj.message
}

// Ring returns the ring
func (obj *vote) Ring() string {
	return obj.ring
}

// Account returns the account
func (obj *vote) Account() string {
	return obj.account
}
