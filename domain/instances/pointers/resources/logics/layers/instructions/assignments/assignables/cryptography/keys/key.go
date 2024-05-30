package keys

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

type key struct {
	hash       hash.Hash
	encryption encryptions.Encryption
	signature  signatures.Signature
}

func createKeyWithEncryption(
	hash hash.Hash,
	encryption encryptions.Encryption,
) Key {
	return createKeyInternally(hash, encryption, nil)
}

func createKeyWithSignature(
	hash hash.Hash,
	signature signatures.Signature,
) Key {
	return createKeyInternally(hash, nil, signature)
}

func createKeyInternally(
	hash hash.Hash,
	encryption encryptions.Encryption,
	signature signatures.Signature,
) Key {
	out := key{
		hash:       hash,
		encryption: encryption,
		signature:  signature,
	}

	return &out
}

// Hash returns the hash
func (obj *key) Hash() hash.Hash {
	return obj.hash
}

// IsEncryption returns true if there is encryption, false otherwise
func (obj *key) IsEncryption() bool {
	return obj.encryption != nil
}

// Encryption returns the encryption, if any
func (obj *key) Encryption() encryptions.Encryption {
	return obj.encryption
}

// IsSignature returns true if there is signature, false otherwise
func (obj *key) IsSignature() bool {
	return obj.signature != nil
}

// Signature returns the signature, if any
func (obj *key) Signature() signatures.Signature {
	return obj.signature
}
