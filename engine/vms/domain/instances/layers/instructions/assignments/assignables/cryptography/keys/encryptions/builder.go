package encryptions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

type builder struct {
	hashAdapter hash.Adapter
	isGenPK     bool
	fetchPubKey string
	encrypt     encrypts.Encrypt
	decrypt     decrypts.Decrypt
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isGenPK:     false,
		fetchPubKey: "",
		encrypt:     nil,
		decrypt:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// IsGeneratePrivateKey flags the builder as generatePK
func (app *builder) IsGeneratePrivateKey() Builder {
	app.isGenPK = true
	return app
}

// WithFetchPublicKey adds a fetchPublicKey to the builder
func (app *builder) WithFetchPublicKey(fetchPublicKey string) Builder {
	app.fetchPubKey = fetchPublicKey
	return app
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

// Now builds a new Encryption instance
func (app *builder) Now() (Encryption, error) {
	data := [][]byte{}
	if app.isGenPK {
		data = append(data, []byte("generatePK"))
	}

	if app.fetchPubKey != "" {
		data = append(data, []byte("fetchPubKey"))
		data = append(data, []byte(app.fetchPubKey))
	}

	if app.encrypt != nil {
		data = append(data, []byte("encrypt"))
		data = append(data, app.encrypt.Hash().Bytes())
	}

	if app.decrypt != nil {
		data = append(data, []byte("decrypt"))
		data = append(data, app.decrypt.Hash().Bytes())
	}

	length := len(data)
	if length != 1 && length != 2 {
		return nil, errors.New("the Encryption is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isGenPK {
		return createEncryptionWithGeneratePrivateKey(*pHash), nil
	}

	if app.fetchPubKey != "" {
		return createEncryptionWithFetchPublicKey(*pHash, app.fetchPubKey), nil
	}

	if app.encrypt != nil {
		return createEncryptionWithEncrypt(*pHash, app.encrypt), nil
	}

	return createEncryptionWithDecrypt(*pHash, app.decrypt), nil
}
