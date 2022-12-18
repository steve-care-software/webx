package references

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type commitBuilder struct {
	pHash      *hash.Hash
	pointer    Pointer
	pCreatedOn *time.Time
}

func createCommitBuilder() CommitBuilder {
	out := commitBuilder{
		pHash:      nil,
		pointer:    nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *commitBuilder) Create() CommitBuilder {
	return createCommitBuilder()
}

// WithHash adds an hash to the builder
func (app *commitBuilder) WithHash(hash hash.Hash) CommitBuilder {
	app.pHash = &hash
	return app
}

// WithPointer adds a pointer to the builder
func (app *commitBuilder) WithPointer(pointer Pointer) CommitBuilder {
	app.pointer = pointer
	return app
}

// CreatedOn adds a creation time to the builder
func (app *commitBuilder) CreatedOn(createdOn time.Time) CommitBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Commit instance
func (app *commitBuilder) Now() (Commit, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Commit instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Commit instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Commit instance")
	}

	return createCommit(*app.pHash, app.pointer, *app.pCreatedOn), nil
}
