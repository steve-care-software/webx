package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type encryptorBuilder struct {
	hashAdapter hash.Adapter
	decrypt     string
	encrypt     string
	isPublicKey bool
}

func createEncryptorBuilder(
	hashAdapter hash.Adapter,
) EncryptorBuilder {
	out := encryptorBuilder{
		hashAdapter: hashAdapter,
		decrypt:     "",
		encrypt:     "",
		isPublicKey: false,
	}

	return &out
}

// Create initializes the builder
func (app *encryptorBuilder) Create() EncryptorBuilder {
	return createEncryptorBuilder(
		app.hashAdapter,
	)
}

// WithDecrypt adds a decrypt to the builder
func (app *encryptorBuilder) WithDecrypt(decrypt string) EncryptorBuilder {
	app.decrypt = decrypt
	return app
}

// WithEncrypt adds an encrypt to the builder
func (app *encryptorBuilder) WithEncrypt(encrypt string) EncryptorBuilder {
	app.encrypt = encrypt
	return app
}

// IsPublicKey flags the builder as isPublicKey
func (app *encryptorBuilder) IsPublicKey() EncryptorBuilder {
	app.isPublicKey = true
	return app
}

// Now builds a new Encryptor instance
func (app *encryptorBuilder) Now() (Encryptor, error) {
	data := [][]byte{}
	if app.decrypt != "" {
		data = append(data, []byte("decrypt"))
		data = append(data, []byte(app.decrypt))
	}

	if app.encrypt != "" {
		data = append(data, []byte("encrypt"))
		data = append(data, []byte(app.encrypt))
	}

	if app.isPublicKey {
		data = append(data, []byte("isPublicKey"))
	}

	if len(data) <= 0 {
		return nil, errors.New("the Encryptor is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.decrypt != "" {
		return createEncryptorWithDecrypt(*pHash, app.decrypt), nil
	}

	if app.encrypt != "" {
		return createEncryptorWithEncrypt(*pHash, app.encrypt), nil
	}

	return createEncryptorWithIsPublicKey(*pHash), nil
}
