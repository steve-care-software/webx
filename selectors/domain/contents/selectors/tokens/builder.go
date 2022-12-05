package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity        entities.Entity
	reverse       entities.Identifier
	element       entities.Identifier
	pElementIndex *uint
	pContentIndex *uint
}

func createBuilder() Builder {
	out := builder{
		entity:        nil,
		reverse:       nil,
		element:       nil,
		pElementIndex: nil,
		pContentIndex: nil,
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

// WithReverse adds a reverse to the builder
func (app *builder) WithReverse(reverse entities.Identifier) Builder {
	app.reverse = reverse
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element entities.Identifier) Builder {
	app.element = element
	return app
}

// WithElementIndex adds an element index to the builder
func (app *builder) WithElementIndex(elementIndex uint) Builder {
	app.pElementIndex = &elementIndex
	return app
}

// WithContentIndex adds a content index to the builder
func (app *builder) WithContentIndex(contentIndex uint) Builder {
	app.pContentIndex = &contentIndex
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Token instance")
	}

	if app.reverse == nil {
		return nil, errors.New("the reverse is mandatory in order to build a Token instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.pElementIndex == nil {
		return nil, errors.New("the element index is mandatory in order to build a Token instance")
	}

	element := createElement(app.element, *app.pElementIndex)
	if app.pContentIndex != nil {
		return createTokenWithContentIndex(app.entity, app.reverse, element, app.pContentIndex), nil
	}

	return createToken(app.entity, app.reverse, element), nil
}
