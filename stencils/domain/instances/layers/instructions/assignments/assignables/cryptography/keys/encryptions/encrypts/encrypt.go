package encrypts

import "github.com/steve-care-software/datastencil/states/domain/hash"

type encrypt struct {
	hash    hash.Hash
	message string
	pubKey  string
}

func createEncrypt(
	hash hash.Hash,
	message string,
	pubKey string,
) Encrypt {
	out := encrypt{
		hash:    hash,
		message: message,
		pubKey:  pubKey,
	}

	return &out
}

// Hash returns the hash
func (obj *encrypt) Hash() hash.Hash {
	return obj.hash
}

// Message returns the message
func (obj *encrypt) Message() string {
	return obj.message
}

// PublicKey returns the publicKey
func (obj *encrypt) PublicKey() string {
	return obj.pubKey
}
