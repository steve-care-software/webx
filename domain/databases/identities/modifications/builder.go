package modifications

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	identifier entities.Identifier
	name       string
	signature  []byte
	encryption []byte
}

func createBuilder() Builder {
	out := builder{
		identifier: nil,
		name:       "",
		signature:  nil,
		encryption: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier entities.Identifier) Builder {
	app.identifier = identifier
	return app
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
	if app.identifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Modification instance")
	}

	if app.name != "" && app.signature != nil && app.encryption != nil {
		content := createContentWithNameAndSignaturePKAndEncryptionPK(app.name, app.signature, app.encryption)
		return createModification(app.identifier, content), nil
	}

	if app.name != "" && app.signature != nil {
		content := createContentWithNameAndSignaturePK(app.name, app.signature)
		return createModification(app.identifier, content), nil
	}

	if app.name != "" && app.encryption != nil {
		content := createContentWithNameAndEncryptionPK(app.name, app.encryption)
		return createModification(app.identifier, content), nil
	}

	if app.signature != nil && app.encryption != nil {
		content := createContentWithSignaturePKAndEncryptionPK(app.signature, app.encryption)
		return createModification(app.identifier, content), nil
	}

	if app.name != "" {
		content := createContentWithName(app.name)
		return createModification(app.identifier, content), nil
	}

	if app.signature != nil {
		content := createContentWithSignaturePK(app.signature)
		return createModification(app.identifier, content), nil
	}

	if app.encryption != nil {
		content := createContentWithEncryptionPK(app.encryption)
		return createModification(app.identifier, content), nil
	}

	return nil, errors.New("the Modification is invalid")
}
