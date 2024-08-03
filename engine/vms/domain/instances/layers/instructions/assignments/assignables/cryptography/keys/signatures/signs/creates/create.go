package creates

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

type create struct {
	hash       hash.Hash
	message    string
	privateKey string
}

func createCreate(
	hash hash.Hash,
	message string,
	privateKey string,
) Create {
	out := create{
		hash:       hash,
		message:    message,
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

// PrivateKey returns the privateKey
func (obj *create) PrivateKey() string {
	return obj.privateKey
}
