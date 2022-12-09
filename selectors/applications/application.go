package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/trees"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// New creates a new selector database
func (app *application) New(name string) error {
	return nil
}

// Retrieve retrieves a selector
func (app *application) Retrieve(context uint, hash hash.Hash) (selectors.Selector, error) {
	return nil, nil
}

// Scan scans the database for the best selector
func (app *application) Scan(context uint, input trees.Tree, output interface{}) (selectors.Selector, error) {
	return nil, nil
}

// Insert inserts a selector
func (app *application) Insert(context uint, selector selectors.Selector) error {
	return nil
}

// InsertAll inserts a list of selectors
func (app *application) InsertAll(context uint, selectors []selectors.Selector) error {
	return nil
}

// Execute executes a selector on a data tree
func (app *application) Execute(selector selectors.Selector, tree trees.Tree) (interface{}, bool, error) {
	return nil, false, nil
}
