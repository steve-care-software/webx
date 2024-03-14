package encrypts

import "github.com/steve-care-software/datastencil/domain/hash"

type encrypt struct {
	hash    hash.Hash
	message string
	account string
}

func createEncrypt(
	hash hash.Hash,
	message string,
	account string,
) Encrypt {
	out := encrypt{
		hash:    hash,
		message: message,
		account: account,
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

// Account returns the account
func (obj *encrypt) Account() string {
	return obj.account
}
