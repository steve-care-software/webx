package encryptors

import (
	"crypto/x509"

	"github.com/steve-care-software/identity/domain/hash"
)

type publicKeyAdapter struct {
	hashAdapter hash.Adapter
	builder     PublicKeyBuilder
}

func createPublicKeyAdapter(
	hashAdapter hash.Adapter,
	builder PublicKeyBuilder,
) PublicKeyAdapter {
	out := publicKeyAdapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}
	return &out
}

// ToPublicKey converts bytes to publicKey
func (app *publicKeyAdapter) ToPublicKey(input []byte) (PublicKey, error) {
	pubKey, err := x509.ParsePKCS1PublicKey(input)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithKey(*pubKey).
		Now()
}

// ToBytes converts publicKey to bytes
func (app *publicKeyAdapter) ToBytes(key PublicKey) []byte {
	pubKey := key.Key()
	return x509.MarshalPKCS1PublicKey(&pubKey)
}

// ToHash converts publicKey to hash
func (app *publicKeyAdapter) ToHash(key PublicKey) (*hash.Hash, error) {
	bytes := app.ToBytes(key)
	return app.hashAdapter.FromBytes(bytes)
}
