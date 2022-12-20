package histories

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type historyBuilder struct {
	pCommit *hash.Hash
	pScore  *uint
}

func createHistoryBuilder() HistoryBuilder {
	out := historyBuilder{
		pCommit: nil,
		pScore:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *historyBuilder) Create() HistoryBuilder {
	return createHistoryBuilder()
}

// WithCommit adds a commit to the builder
func (app *historyBuilder) WithCommit(commit hash.Hash) HistoryBuilder {
	app.pCommit = &commit
	return app
}

// WithScore adds a score to the builder
func (app *historyBuilder) WithScore(score uint) HistoryBuilder {
	app.pScore = &score
	return app
}

// Now builds a new History instance
func (app *historyBuilder) Now() (History, error) {
	if app.pCommit == nil {
		return nil, errors.New("the commit is mandatory in order to build an History instance")
	}

	if app.pScore == nil {
		return nil, errors.New("the score is mandatory in order to build an History instance")
	}

	return createHistory(*app.pCommit, *app.pScore), nil
}
