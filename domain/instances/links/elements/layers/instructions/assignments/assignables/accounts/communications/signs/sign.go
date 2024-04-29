package signs

import "github.com/steve-care-software/datastencil/domain/hash"

type sign struct {
	hash    hash.Hash
	message string
	account string
}

func createSign(
	hash hash.Hash,
	message string,
	account string,
) Sign {
	out := sign{
		hash:    hash,
		message: message,
		account: account,
	}

	return &out
}

// Hash returns the hash
func (obj *sign) Hash() hash.Hash {
	return obj.hash
}

// Message returns the message
func (obj *sign) Message() string {
	return obj.message
}

// Account returns the account
func (obj *sign) Account() string {
	return obj.account
}
