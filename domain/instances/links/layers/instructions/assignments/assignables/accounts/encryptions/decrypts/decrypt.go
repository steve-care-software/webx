package decrypts

import "github.com/steve-care-software/datastencil/domain/hash"

type decrypt struct {
	hash    hash.Hash
	cipher  string
	account string
}

func createDecrypt(
	hash hash.Hash,
	cipher string,
	account string,
) Decrypt {
	out := decrypt{
		hash:    hash,
		cipher:  cipher,
		account: account,
	}

	return &out
}

// Hash returns the hash
func (obj *decrypt) Hash() hash.Hash {
	return obj.hash
}

// Cipher returns the cipher
func (obj *decrypt) Cipher() string {
	return obj.cipher
}

// Account returns the account
func (obj *decrypt) Account() string {
	return obj.account
}
