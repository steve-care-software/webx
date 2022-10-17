package keys

import (
	"crypto/rsa"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

// NewFactory returns a new encryption's privatekey factory
func NewFactory(bitrate int) Factory {
	builder := NewBuilder()
	return createFactory(builder, bitrate)
}

// NewAdapter returns a new encryption's privatekey adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder returns a new encryption's privatekey builder
func NewBuilder() Builder {
	pubKeyBuilder := NewPublicKeyBuilder()
	return createBuilder(pubKeyBuilder)
}

// NewPublicKeyAdapter creates a new publicKey adapter
func NewPublicKeyAdapter() PublicKeyAdapter {
	hashAdapter := hash.NewAdapter()
	publicKeyBuilder := NewPublicKeyBuilder()
	return createPublicKeyAdapter(hashAdapter, publicKeyBuilder)
}

// NewPublicKeyBuilder creates a new public key builder
func NewPublicKeyBuilder() PublicKeyBuilder {
	return createPublicKeyBuilder()
}

// Factory represents a privateKey factory
type Factory interface {
	Create() (PrivateKey, error)
}

// Adapter represents a privateKey adapter
type Adapter interface {
	FromBytes(bytes []byte) (PrivateKey, error)
	FromEncoded(encoded string) (PrivateKey, error)
	ToBytes(pk PrivateKey) []byte
	ToEncoded(pk PrivateKey) string
}

// Builder represents a privateKey builder
type Builder interface {
	Create() Builder
	WithPK(pk rsa.PrivateKey) Builder
	Now() (PrivateKey, error)
}

// PrivateKey represents an encryption private key
type PrivateKey interface {
	Key() rsa.PrivateKey
	Public() PublicKey
	Decrypt(cipher []byte) ([]byte, error)
}

// PublicKeyAdapter represents a public key adapter
type PublicKeyAdapter interface {
	FromBytes(input []byte) (PublicKey, error)
	FromEncoded(encoded string) (PublicKey, error)
	ToBytes(key PublicKey) []byte
	ToEncoded(key PublicKey) string
	ToHash(key PublicKey) (*hash.Hash, error)
}

// PublicKeyBuilder represents a publicKey builder
type PublicKeyBuilder interface {
	Create() PublicKeyBuilder
	WithKey(key rsa.PublicKey) PublicKeyBuilder
	Now() (PublicKey, error)
}

// PublicKey represents an encryption public key
type PublicKey interface {
	Key() rsa.PublicKey
	Encrypt(msg []byte) ([]byte, error)
}
