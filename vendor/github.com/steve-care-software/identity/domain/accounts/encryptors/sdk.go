package encryptors

import (
	"crypto/rsa"

	"github.com/steve-care-software/identity/domain/hash"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(
		builder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	pubKeyBuilder := NewPublicKeyBuilder()
	return createBuilder(
		pubKeyBuilder,
	)
}

// NewPublicKeyAdapter creates a new public key adapter
func NewPublicKeyAdapter() PublicKeyAdapter {
	hashAdapter := hash.NewAdapter()
	builder := NewPublicKeyBuilder()
	return createPublicKeyAdapter(
		hashAdapter,
		builder,
	)
}

// NewPublicKeyBuilder creates a new public key builder
func NewPublicKeyBuilder() PublicKeyBuilder {
	return createPublicKeyBuilder()
}

// Adapter represents a privateKey adapter
type Adapter interface {
	ToEncryptor(bytes []byte) (Encryptor, error)
	ToBytes(pk Encryptor) []byte
}

// Builder represents a privateKey builder
type Builder interface {
	Create() Builder
	WithPK(pk rsa.PrivateKey) Builder
	WithBitRate(bitRate int) Builder
	Now() (Encryptor, error)
}

// Encryptor represents an encryptor
type Encryptor interface {
	Decrypt(cipher []byte) ([]byte, error)
	Public() PublicKey
	Key() rsa.PrivateKey
}

// PublicKeyAdapter represents a public key adapter
type PublicKeyAdapter interface {
	ToPublicKey(input []byte) (PublicKey, error)
	ToBytes(key PublicKey) []byte
	ToHash(key PublicKey) (*hash.Hash, error)
}

// PublicKeyBuilder represents a public key builder
type PublicKeyBuilder interface {
	Create() PublicKeyBuilder
	WithKey(key rsa.PublicKey) PublicKeyBuilder
	Now() (PublicKey, error)
}

// PublicKey represents a public key
type PublicKey interface {
	Encrypt(msg []byte) ([]byte, error)
	Key() rsa.PublicKey
}
