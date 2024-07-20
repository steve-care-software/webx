package validates

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type validate struct {
	hash      hash.Hash
	signature string
	message   string
	publicKey string
}

func createValidate(
	hash hash.Hash,
	signature string,
	message string,
	publicKey string,
) Validate {
	out := validate{
		hash:      hash,
		signature: signature,
		message:   message,
		publicKey: publicKey,
	}

	return &out
}

// Hash returns the hash
func (obj *validate) Hash() hash.Hash {
	return obj.hash
}

// Signature returns the signature
func (obj *validate) Signature() string {
	return obj.signature
}

// Message returns the message
func (obj *validate) Message() string {
	return obj.message
}

// PublicKey returns the publicKey
func (obj *validate) PublicKey() string {
	return obj.publicKey
}
