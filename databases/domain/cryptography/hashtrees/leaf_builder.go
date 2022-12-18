package hashtrees

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type leafBuilder struct {
	pHead  *hash.Hash
	parent ParentLeaf
}

func createLeafBuilder() LeafBuilder {
	out := leafBuilder{
		pHead:  nil,
		parent: nil,
	}

	return &out
}

// Create initializes the builder
func (app *leafBuilder) Create() LeafBuilder {
	return createLeafBuilder()
}

// WithHead adds an head to the builder
func (app *leafBuilder) WithHead(head hash.Hash) LeafBuilder {
	app.pHead = &head
	return app
}

// WithParent adds a parent to the builder
func (app *leafBuilder) WithParent(parent ParentLeaf) LeafBuilder {
	app.parent = parent
	return app
}

// Now builds a new Leaf instance
func (app *leafBuilder) Now() (Leaf, error) {
	if app.pHead == nil {
		return nil, errors.New("the head is mandatory in order to build a Leaf instance")
	}

	if app.parent != nil {
		return createLeafWithParent(*app.pHead, app.parent), nil
	}

	return createLeaf(*app.pHead), nil
}
