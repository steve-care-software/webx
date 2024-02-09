package layers

import "github.com/steve-care-software/identity/domain/hash"

type signatureVerify struct {
	hash      hash.Hash
	signature string
	message   string
}

func createSignatureVerify(
	hash hash.Hash,
	signature string,
	message string,
) SignatureVerify {
	out := signatureVerify{
		hash:      hash,
		signature: signature,
		message:   message,
	}

	return &out
}

// Hash returns the hash
func (obj *signatureVerify) Hash() hash.Hash {
	return obj.hash
}

// Signature returns the signature
func (obj *signatureVerify) Signature() string {
	return obj.signature
}

// Message returns the message
func (obj *signatureVerify) Message() string {
	return obj.message
}
