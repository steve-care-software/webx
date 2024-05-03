package encryptions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
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
	if app.isGenPK {
		return createEncryptionWithGeneratePrivateKey(), nil
	}

	if app.fetchPubKey != "" {
		return createEncryptionWithFetchPublicKey(app.fetchPubKey), nil
	}

	if app.encrypt != nil {
		return createEncryptionWithEncrypt(app.encrypt), nil
	}

	if app.decrypt != nil {
		return createEncryptionWithDecrypt(app.decrypt), nil
	}

	return nil, errors.New("the Encryption is invalid")
}
