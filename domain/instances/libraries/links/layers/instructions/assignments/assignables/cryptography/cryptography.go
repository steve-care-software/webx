package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/cryptography/encrypts"
)

type cryptography struct {
	hash    hash.Hash
	encrypt encrypts.Encrypt
	decrypt decrypts.Decrypt
}

func createCryptographyWithEncrypt(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
) Cryptography {
	return createCryptographyInternally(hash, encrypt, nil)
}

func createCryptographyWithDecrypt(
	hash hash.Hash,
	decrypt decrypts.Decrypt,
) Cryptography {
	return createCryptographyInternally(hash, nil, decrypt)
}

func createCryptographyInternally(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
	decrypt decrypts.Decrypt,
) Cryptography {
	out := cryptography{
		hash:    hash,
		encrypt: encrypt,
		decrypt: decrypt,
	}

	return &out
}

// Hash returns the hash
func (obj *cryptography) Hash() hash.Hash {
	return obj.hash
}

// IsEncrypt returns true if there is an encrypt, false otherwise
func (obj *cryptography) IsEncrypt() bool {
	return obj.encrypt != nil
}

// Encrypt returns the encrypt, if any
func (obj *cryptography) Encrypt() encrypts.Encrypt {
	return obj.encrypt
}

// IsDecrypt returns true if there is a decrypt, false otherwise
func (obj *cryptography) IsDecrypt() bool {
	return obj.decrypt != nil
}

// Decrypt returns the decrypt, if any
func (obj *cryptography) Decrypt() decrypts.Decrypt {
	return obj.decrypt
}
