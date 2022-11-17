package entries

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type linkBuilder struct {
	new       Entry
	pExisting *hash.Hash
}

func createLinkBuilder() LinkBuilder {
	out := linkBuilder{
		new:       nil,
		pExisting: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder()
}

// WithNew adds a new entry to the builder
func (app *linkBuilder) WithNew(new Entry) LinkBuilder {
	app.new = new
	return app
}

// WithExisting adds an existing hash to the builder
func (app *linkBuilder) WithExisting(existing hash.Hash) LinkBuilder {
	app.pExisting = &existing
	return app
}

// WithExisting adds an existing hash to the builder
func (app *linkBuilder) Now() (Link, error) {
	if app.new != nil {
		return createLinkWithNew(app.new), nil
	}

	if app.pExisting != nil {
		return createLinkWithExisting(app.pExisting), nil
	}

	return nil, errors.New("the Link is invalid")
}
