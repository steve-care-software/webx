package modifications

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	signature   []byte
	encryption  []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		signature:   nil,
		encryption:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithSignature adds a sigPK to the builder
func (app *builder) WithSignature(signature []byte) Builder {
	app.signature = signature
	return app
}

// WithEncryption adds an encPK to the builder
func (app *builder) WithEncryption(encryption []byte) Builder {
	app.encryption = encryption
	return app
}

// Now builds a new Modification instance
func (app *builder) Now() (Modification, error) {

	data := [][]byte{}
	if app.name != "" {
		data = append(data, []byte(app.name))
	}

	if app.signature != nil {
		data = append(data, app.signature)
	}

	if app.encryption != nil {
		data = append(data, app.encryption)
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.name != "" && app.signature != nil && app.encryption != nil {
		content := createContentWithNameAndSignaturePKAndEncryptionPK(app.name, app.signature, app.encryption)
		return createModification(*pHash, content), nil
	}

	if app.name != "" && app.signature != nil {
		content := createContentWithNameAndSignaturePK(app.name, app.signature)
		return createModification(*pHash, content), nil
	}

	if app.name != "" && app.encryption != nil {
		content := createContentWithNameAndEncryptionPK(app.name, app.encryption)
		return createModification(*pHash, content), nil
	}

	if app.signature != nil && app.encryption != nil {
		content := createContentWithSignaturePKAndEncryptionPK(app.signature, app.encryption)
		return createModification(*pHash, content), nil
	}

	if app.name != "" {
		content := createContentWithName(app.name)
		return createModification(*pHash, content), nil
	}

	if app.signature != nil {
		content := createContentWithSignaturePK(app.signature)
		return createModification(*pHash, content), nil
	}

	if app.encryption != nil {
		content := createContentWithEncryptionPK(app.encryption)
		return createModification(*pHash, content), nil
	}

	return nil, errors.New("the Modification is invalid")
}
