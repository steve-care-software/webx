package decrypts

import "github.com/steve-care-software/datastencil/states/domain/hash"

type decrypt struct {
	hash     hash.Hash
	cipher   string
	password string
}

func createDecrypt(
	hash hash.Hash,
	cipher string,
	password string,
) Decrypt {
	out := decrypt{
		hash:     hash,
		cipher:   cipher,
		password: password,
	}

	return &out
}

// Hash returns the cipher
func (obj *decrypt) Hash() hash.Hash {
	return obj.hash
}

// Cipher returns the cipher
func (obj *decrypt) Cipher() string {
	return obj.cipher
}

// Password returns the password
func (obj *decrypt) Password() string {
	return obj.password
}
