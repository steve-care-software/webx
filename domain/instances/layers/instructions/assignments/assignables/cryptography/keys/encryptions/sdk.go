package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the encryption adapter
type Adapter interface {
	ToBytes(ins Encryption) ([]byte, error)
	ToInstance(bytes []byte) (Encryption, error)
}

// Builder represents an encryption builder
type Builder interface {
	Create() Builder
	IsGeneratePrivateKey() Builder
	WithFetchPublicKey(fetchPublicKey string) Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	Now() (Encryption, error)
}

// Encryption represents encryption
type Encryption interface {
	Hash() hash.Hash
	IsGeneratePrivateKey() bool
	IsFetchPublicKey() bool
	FetchPublicKey() string
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
}
