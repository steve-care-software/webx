package decrypts

import "github.com/steve-care-software/historydb/domain/hash"

type decrypt struct {
	hash       hash.Hash
	cipher     string
	privateKey string
}

func createDecrypt(
	hash hash.Hash,
	cipher string,
	privateKey string,
) Decrypt {
	out := decrypt{
		hash:       hash,
		cipher:     cipher,
		privateKey: privateKey,
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

// PrivateKey returns the privateKey
func (obj *decrypt) PrivateKey() string {
	return obj.privateKey
}
