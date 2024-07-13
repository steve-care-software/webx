package encrypts

import "github.com/steve-care-software/historydb/domain/hash"

type encrypt struct {
	hash     hash.Hash
	message  string
	password string
}

func createEncrypt(
	hash hash.Hash,
	message string,
	password string,
) Encrypt {
	out := encrypt{
		hash:     hash,
		message:  message,
		password: password,
	}

	return &out
}

// Hash returns the message
func (obj *encrypt) Hash() hash.Hash {
	return obj.hash
}

// Message returns the message
func (obj *encrypt) Message() string {
	return obj.message
}

// Password returns the password
func (obj *encrypt) Password() string {
	return obj.password
}
