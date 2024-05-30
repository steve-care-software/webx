package cryptography

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys"
)

type builder struct {
	hashAdapter hash.Adapter
	encrypt     encrypts.Encrypt
	decrypt     decrypts.Decrypt
	key         keys.Key
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		encrypt:     nil,
		decrypt:     nil,
		key:         nil,
	}

	return &out
}

// Create intiializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEncrypt adds an encrypt to the builder
func (app *builder) WithEncrypt(encrypt encrypts.Encrypt) Builder {
	app.encrypt = encrypt
	return app
}

// WithDecrypt adds a decrypt to the builder
func (app *builder) WithDecrypt(decrypt decrypts.Decrypt) Builder {
	app.decrypt = decrypt
	return app
}

// WithKey adds a key to the builder
func (app *builder) WithKey(key keys.Key) Builder {
	app.key = key
	return app
}

// Now builds a new Cryptography instance
func (app *builder) Now() (Cryptography, error) {
	data := [][]byte{}
	if app.encrypt != nil {
		data = append(data, []byte("encrypt"))
		data = append(data, app.encrypt.Hash().Bytes())
	}

	if app.decrypt != nil {
		data = append(data, []byte("decrypt"))
		data = append(data, app.decrypt.Hash().Bytes())
	}

	if app.key != nil {
		data = append(data, []byte("key"))
		data = append(data, app.key.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Cryptography is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.encrypt != nil {
		return createCryptographyWithEncrypt(*pHash, app.encrypt), nil
	}

	if app.decrypt != nil {
		return createCryptographyWithDecrypt(*pHash, app.decrypt), nil
	}

	return createCryptographyWithKey(*pHash, app.key), nil
}
