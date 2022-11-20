package entries

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type relationBuilder struct {
	new      Entries
	existing []hash.Hash
}

func createRelationBuilder() RelationBuilder {
	out := relationBuilder{
		new:      nil,
		existing: nil,
	}

	return &out
}

// Create initializes the builder
func (app *relationBuilder) Create() RelationBuilder {
	return createRelationBuilder()
}

// WithNew add new entries to the builder
func (app *relationBuilder) WithNew(new Entries) RelationBuilder {
	app.new = new
	return app
}

// WithExisting add existing hashes to the builder
func (app *relationBuilder) WithExisting(existing []hash.Hash) RelationBuilder {
	app.existing = existing
	return app
}

// Now builds a new Relation instance
func (app *relationBuilder) Now() (Relation, error) {
	if app.new != nil {
		return createRelationWithNew(app.new), nil
	}

	if app.existing != nil && len(app.existing) <= 0 {
		app.existing = nil
	}

	if app.existing != nil {
		return createRelationWithExisting(app.existing), nil
	}

	return nil, errors.New("the Relation is invalid")
}
