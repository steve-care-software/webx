package blockchains

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity     entities.Entity
	pReference *hash.Hash
	head       entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:     nil,
		pReference: nil,
		head:       nil,
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

// WithReference adds a reference to the builder
func (app *builder) WithReference(reference hash.Hash) Builder {
	app.pReference = &reference
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head entities.Identifier) Builder {
	app.head = head
	return app
}

// Now builds a new Blockchain instance
func (app *builder) Now() (Blockchain, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Blockchain instance")
	}

	if app.pReference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Blockchain instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Blockchain instance")
	}

	return createBlockchain(app.entity, *app.pReference, app.head), nil
}
