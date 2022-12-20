package histories

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	pCommit *hash.Hash
	pScore  *uint
}

func createBuilder() Builder {
	out := builder{
		pCommit: nil,
		pScore:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCommit adds a commit to the builder
func (app *builder) WithCommit(commit hash.Hash) Builder {
	app.pCommit = &commit
	return app
}

// WithScore adds a score to the builder
func (app *builder) WithScore(score uint) Builder {
	app.pScore = &score
	return app
}

// Now builds a new History instance
func (app *builder) Now() (History, error) {
	if app.pCommit == nil {
		return nil, errors.New("the commit is mandatory in order to build an History instance")
	}

	if app.pScore == nil {
		return nil, errors.New("the score is mandatory in order to build an History instance")
	}

	return createHistory(*app.pCommit, *app.pScore), nil
}
