package keys

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

type builder struct {
	hashAdapter hash.Adapter
	encryption  encryptions.Encryption
	signature   signatures.Signature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		encryption:  nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEncryption adds an encryption to the builder
func (app *builder) WithEncryption(enc encryptions.Encryption) Builder {
	app.encryption = enc
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(sig signatures.Signature) Builder {
	app.signature = sig
	return app
}

// Now builds a new Key instance
func (app *builder) Now() (Key, error) {
	data := [][]byte{}
	if app.encryption != nil {
		data = append(data, []byte("encryption"))
		data = append(data, app.encryption.Hash().Bytes())
	}

	if app.signature != nil {
		data = append(data, []byte("signature"))
		data = append(data, app.signature.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Key is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.encryption != nil {
		return createKeyWithEncryption(*pHash, app.encryption), nil
	}

	return createKeyWithSignature(*pHash, app.signature), nil

}
