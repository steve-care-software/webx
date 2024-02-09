package layers

import "github.com/steve-care-software/identity/domain/hash"

type encryptor struct {
	hash        hash.Hash
	decrypt     string
	encrypt     string
	isPublicKey bool
}

func createEncryptorWithDecrypt(
	hash hash.Hash,
	decrypt string,
) Encryptor {
	return createEncryptorInternally(hash, decrypt, "", false)
}

func createEncryptorWithEncrypt(
	hash hash.Hash,
	encrypt string,
) Encryptor {
	return createEncryptorInternally(hash, "", encrypt, false)
}

func createEncryptorWithIsPublicKey(
	hash hash.Hash,
) Encryptor {
	return createEncryptorInternally(hash, "", "", true)
}

func createEncryptorInternally(
	hash hash.Hash,
	decrypt string,
	encrypt string,
	isPublicKey bool,
) Encryptor {
	out := encryptor{
		hash:        hash,
		decrypt:     decrypt,
		encrypt:     encrypt,
		isPublicKey: isPublicKey,
	}

	return &out
}

// Hash returns the hash
func (obj *encryptor) Hash() hash.Hash {
	return obj.hash
}

// IsDecrypt returns true if there is a decrypt, false otherwise
func (obj *encryptor) IsDecrypt() bool {
	return obj.decrypt != ""
}

// Decrypt returns the decrypt, if any
func (obj *encryptor) Decrypt() string {
	return obj.decrypt
}

// IsEncrypt returns true if there is an encrypt, false otherwise
func (obj *encryptor) IsEncrypt() bool {
	return obj.encrypt != ""
}

// Encrypt returns the encrypt, if any
func (obj *encryptor) Encrypt() string {
	return obj.encrypt
}

// IsPublicKey returns true if isPublicKey, false otherwise
func (obj *encryptor) IsPublicKey() bool {
	return obj.isPublicKey
}
