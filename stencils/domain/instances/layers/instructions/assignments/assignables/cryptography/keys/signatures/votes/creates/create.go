package creates

import "github.com/steve-care-software/datastencil/states/domain/hash"

type create struct {
	hash       hash.Hash
	message    string
	ring       string
	privateKey string
}

func createCreate(
	hash hash.Hash,
	message string,
	ring string,
	privateKey string,
) Create {
	out := create{
		hash:       hash,
		message:    message,
		ring:       ring,
		privateKey: privateKey,
	}

	return &out
}

// Hash returns the hash
func (obj *create) Hash() hash.Hash {
	return obj.hash
}

// Message returns the message
func (obj *create) Message() string {
	return obj.message
}

// Ring returns the ring
func (obj *create) Ring() string {
	return obj.ring
}

// PrivateKey returns the privateKey
func (obj *create) PrivateKey() string {
	return obj.privateKey
}
