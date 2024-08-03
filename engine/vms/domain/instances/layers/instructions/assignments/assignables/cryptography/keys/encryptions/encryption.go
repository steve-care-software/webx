package encryptions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

type encryption struct {
	hash        hash.Hash
	isGenPK     bool
	fetchPubKey string
	encrypt     encrypts.Encrypt
	decrypt     decrypts.Decrypt
}

func createEncryptionWithGeneratePrivateKey(
	hash hash.Hash,
) Encryption {
	return createEncryptionInternally(hash, true, "", nil, nil)
}

func createEncryptionWithFetchPublicKey(
	hash hash.Hash,
	fetchPubKey string,
) Encryption {
	return createEncryptionInternally(hash, false, fetchPubKey, nil, nil)
}

func createEncryptionWithEncrypt(
	hash hash.Hash,
	encrypt encrypts.Encrypt,
) Encryption {
	return createEncryptionInternally(hash, false, "", encrypt, nil)
}

func createEncryptionWithDecrypt(
	hash hash.Hash,
	decrypt decrypts.Decrypt,
) Encryption {
	return createEncryptionInternally(hash, false, "", nil, decrypt)
}

func createEncryptionInternally(
	hash hash.Hash,
	isGenPK bool,
	fetchPubKey string,
	encrypt encrypts.Encrypt,
	decrypt decrypts.Decrypt,
) Encryption {
	out := encryption{
		hash:        hash,
		isGenPK:     isGenPK,
		fetchPubKey: fetchPubKey,
		encrypt:     encrypt,
		decrypt:     decrypt,
	}

	return &out
}

// Hash returns the hash
func (obj *encryption) Hash() hash.Hash {
	return obj.hash
}

// IsGeneratePrivateKey returns true if generatePK, false otherwise
func (obj *encryption) IsGeneratePrivateKey() bool {
	return obj.isGenPK
}

// IsFetchPublicKey returns true if fetchPublicKey, false otherwise
func (obj *encryption) IsFetchPublicKey() bool {
	return obj.fetchPubKey != ""
}

// FetchPublicKey returns the fetchPublicKey, if any
func (obj *encryption) FetchPublicKey() string {
	return obj.fetchPubKey
}

// IsEncrypt returns true if encrypt, false otherwise
func (obj *encryption) IsEncrypt() bool {
	return obj.encrypt != nil
}

// Encrypt returns the encrypt, if any
func (obj *encryption) Encrypt() encrypts.Encrypt {
	return obj.encrypt
}

// IsDecrypt returns true if decrypt, false otherwise
func (obj *encryption) IsDecrypt() bool {
	return obj.decrypt != nil
}

// Decrypt returns the decrypt, if any
func (obj *encryption) Decrypt() decrypts.Decrypt {
	return obj.decrypt
}
