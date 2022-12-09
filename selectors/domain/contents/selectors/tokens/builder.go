package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash         *hash.Hash
	pReverse      *hash.Hash
	pElement      *hash.Hash
	pElementIndex *uint
	pContentIndex *uint
}

func createBuilder() Builder {
	out := builder{
		pHash:         nil,
		pReverse:      nil,
		pElement:      nil,
		pElementIndex: nil,
		pContentIndex: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.pHash = &hash
	return app
}

// WithReverse adds a reverse to the builder
func (app *builder) WithReverse(reverse hash.Hash) Builder {
	app.pReverse = &reverse
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element hash.Hash) Builder {
	app.pElement = &element
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
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Token instance")
	}

	if app.pReverse == nil {
		return nil, errors.New("the reverse is mandatory in order to build a Token instance")
	}

	if app.pElement == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.pElementIndex == nil {
		return nil, errors.New("the element index is mandatory in order to build a Token instance")
	}

	element := createElement(*app.pElement, *app.pElementIndex)
	if app.pContentIndex != nil {
		return createTokenWithContentIndex(*app.pHash, *app.pReverse, element, app.pContentIndex), nil
	}

	return createToken(*app.pHash, *app.pReverse, element), nil
}
