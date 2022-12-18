package hashtrees

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type compactBuilder struct {
	pHead  *hash.Hash
	leaves Leaves
}

func createCompactBuilder() CompactBuilder {
	out := compactBuilder{
		pHead:  nil,
		leaves: nil,
	}

	return &out
}

// Create initializes the builder
func (app *compactBuilder) Create() CompactBuilder {
	return createCompactBuilder()
}

// WithHead adds an head to the builder
func (app *compactBuilder) WithHead(head hash.Hash) CompactBuilder {
	app.pHead = &head
	return app
}

// WithLeaves add leaves to the builder
func (app *compactBuilder) WithLeaves(leaves Leaves) CompactBuilder {
	app.leaves = leaves
	return app
}

// Now builds a new Compact instance
func (app *compactBuilder) Now() (Compact, error) {
	if app.pHead == nil {
		return nil, errors.New("the head is mandatory in order to build  a Compact instance")
	}

	if app.leaves == nil {
		return nil, errors.New("the leaves is mandatory in order to build  a Compact instance")
	}

	return createCompact(*app.pHead, app.leaves), nil
}
