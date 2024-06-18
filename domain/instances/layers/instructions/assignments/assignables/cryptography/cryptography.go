package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
)

type cryptography struct {
	hash    hash.Hash
	encrypt encrypts.Encrypt
	decrypt decrypts.Decrypt
	key     keys.Key
}

func createCryptographyWithEncrypt(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
) Cryptography {
	return createCryptographyInternally(hash, encrypt, nil, nil)
}

func createCryptographyWithDecrypt(
	hash hash.Hash,
	decrypt decrypts.Decrypt,
) Cryptography {
	return createCryptographyInternally(hash, nil, decrypt, nil)
}

func createCryptographyWithKey(
	hash hash.Hash,
	key keys.Key,
) Cryptography {
	return createCryptographyInternally(hash, nil, nil, key)
}

func createCryptographyInternally(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
	decrypt decrypts.Decrypt,
	key keys.Key,
) Cryptography {
	out := cryptography{
		hash:    hash,
		encrypt: encrypt,
		decrypt: decrypt,
		key:     key,
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

// IsKey returns true if there is a key, false otherwise
func (obj *cryptography) IsKey() bool {
	return obj.key != nil
}

// Key returns the key, if any
func (obj *cryptography) Key() keys.Key {
	return obj.key
}
