package keys

import (
	"crypto/x509"
	"encoding/base64"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type publicKeyAdapter struct {
	hashAdapter      hash.Adapter
	publicKeyBuilder PublicKeyBuilder
}

func createPublicKeyAdapter(hashAdapter hash.Adapter, publicKeyBuilder PublicKeyBuilder) PublicKeyAdapter {
	out := publicKeyAdapter{
		hashAdapter:      hashAdapter,
		publicKeyBuilder: publicKeyBuilder,
	}

	return &out
}

// FromBytes converts []byte to Key
func (app *publicKeyAdapter) FromBytes(input []byte) (PublicKey, error) {
	pubKey, err := x509.ParsePKCS1PublicKey(input)
	if err != nil {
		return nil, err
	}

	return app.publicKeyBuilder.Create().WithKey(*pubKey).Now()
}

// FromEncoded converts an encoded string to Key
func (app *publicKeyAdapter) FromEncoded(encoded string) (PublicKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return app.FromBytes(decoded)
}

// ToBytes converts Key to []byte
func (app *publicKeyAdapter) ToBytes(key PublicKey) []byte {
	pubKey := key.Key()
	return x509.MarshalPKCS1PublicKey(&pubKey)
}

// ToEncoded converts a Key to an encoded string
func (app *publicKeyAdapter) ToEncoded(key PublicKey) string {
	bytes := app.ToBytes(key)
	return base64.StdEncoding.EncodeToString(bytes)
}

// ToHash converts a Key to an hash
func (app *publicKeyAdapter) ToHash(key PublicKey) (*hash.Hash, error) {
	bytes := app.ToBytes(key)
	return app.hashAdapter.FromBytes(bytes)
}
