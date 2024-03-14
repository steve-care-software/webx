package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

type encryption struct {
	hash    hash.Hash
	encrypt encrypts.Encrypt
	decrypt decrypts.Decrypt
}

func createEncryptionWithEncrypt(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
) Encryption {
	return createEncryptionInternally(hash, encrypt, nil)
}

func createEncryptionWithDecrypt(
	hash hash.Hash,
	decrypt decrypts.Decrypt,
) Encryption {
	return createEncryptionInternally(hash, nil, decrypt)
}

func createEncryptionInternally(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
	decrypt decrypts.Decrypt,
) Encryption {
	out := encryption{
		hash:    hash,
		encrypt: encrypt,
		decrypt: decrypt,
	}

	return &out
}

// Hash returns the hash
func (obj *encryption) Hash() hash.Hash {
	return obj.hash
}

// IsEncrypt returns true if encrypt, false otherwse
func (obj *encryption) IsEncrypt() bool {
	return obj.encrypt != nil
}

// Encrypt returns encrypt, if any
func (obj *encryption) Encrypt() encrypts.Encrypt {
	return obj.encrypt
}

// IsDecrypt returns true if decrypt, false otherwse
func (obj *encryption) IsDecrypt() bool {
	return obj.decrypt != nil
}

// Decrypt returns decrypt, if any
func (obj *encryption) Decrypt() decrypts.Decrypt {
	return obj.decrypt
}
