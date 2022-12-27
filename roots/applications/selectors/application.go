package selectors

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/trees"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

type application struct {
	grammarSoftware grammar_applications.Application
}

func createApplication(
	grammarSoftware grammar_applications.Application,
) Application {
	out := application{
		grammarSoftware: grammarSoftware,
	}

	return &out
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
