package modifications

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity     entities.Entity
	name       string
	signature  []byte
	encryption []byte
}

func createBuilder() Builder {
	out := builder{
		entity:     nil,
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

// WithEntity adds an entity to the builder
func (app *builder) WithEntity(entity entities.Entity) Builder {
	app.entity = entity
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
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Modification instance")
	}

	if app.name != "" && app.signature != nil && app.encryption != nil {
		content := createContentWithNameAndSignaturePKAndEncryptionPK(app.name, app.signature, app.encryption)
		return createModification(app.entity, content), nil
	}

	if app.name != "" && app.signature != nil {
		content := createContentWithNameAndSignaturePK(app.name, app.signature)
		return createModification(app.entity, content), nil
	}

	if app.name != "" && app.encryption != nil {
		content := createContentWithNameAndEncryptionPK(app.name, app.encryption)
		return createModification(app.entity, content), nil
	}

	if app.signature != nil && app.encryption != nil {
		content := createContentWithSignaturePKAndEncryptionPK(app.signature, app.encryption)
		return createModification(app.entity, content), nil
	}

	if app.name != "" {
		content := createContentWithName(app.name)
		return createModification(app.entity, content), nil
	}

	if app.signature != nil {
		content := createContentWithSignaturePK(app.signature)
		return createModification(app.entity, content), nil
	}

	if app.encryption != nil {
		content := createContentWithEncryptionPK(app.encryption)
		return createModification(app.entity, content), nil
	}

	return nil, errors.New("the Modification is invalid")
}
