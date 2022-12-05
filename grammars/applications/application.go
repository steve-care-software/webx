package applications

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/trees"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/coverages"
)

type application struct {
	blockchainApp applications.Application
}

func createApplication(
	blockchainApp applications.Application,
) Application {
	out := application{
		blockchainApp: blockchainApp,
	}

	return &out
}

// Open opens the context
func (app *application) Open(name string, height int) (*uint, error) {
	return nil, nil
}

// Retrieve retrieves a grammar by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error) {
	return nil, nil
}

// Search searches a grammar by suite
func (app *application) Search(context uint, suites grammars.Suites) (grammars.Grammar, error) {
	return nil, nil
}

// Insert inserts a grammar
func (app *application) Insert(context uint, grammar grammars.Grammar) error {
	return nil
}

// InsertAll inserts a list of grammars
func (app *application) InsertAll(context uint, grammars []grammars.Grammar) error {
	return nil
}

// Execute executes grammar on data
func (app *application) Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error) {
	return nil, nil
}

// Coverages returns the coverages of a grammar
func (app *application) Coverages(grammar grammars.Grammar) (coverages.Coverages, error) {
	return nil, nil
}

// Covered returns the covered tokens
func (app *application) Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error) {
	return nil, nil
}

// Uncovered returns the uncovered tokens
func (app *application) Uncovered(grammar grammars.Grammar) (map[string]map[uint]map[uint]string, error) {
	return nil, nil
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	return nil
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	return nil
}

// Push pushes a context
func (app *application) Push(context uint) error {
	return nil
}

// Close closes a context
func (app *application) Close(context uint) error {
	return nil
}
