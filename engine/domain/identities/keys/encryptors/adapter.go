package encryptors

import "crypto/x509"

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// ToEncryptor converts bytes to encryptor
func (app *adapter) ToEncryptor(bytes []byte) (Encryptor, error) {
	pk, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithPK(*pk).
		Now()
}

// ToBytes converts encryptor to bytes
func (app *adapter) ToBytes(encryptor Encryptor) []byte {
	key := encryptor.Key()
	return x509.MarshalPKCS1PrivateKey(&key)
}
