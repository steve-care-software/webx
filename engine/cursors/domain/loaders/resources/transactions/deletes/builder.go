package deletes

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	vote        signers.Vote
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		vote:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithVote adds a vote to the builder
func (app *builder) WithVote(vote signers.Vote) Builder {
	app.vote = vote
	return app
}

// Now builds a new Delete instance
func (app *builder) Now() (Delete, error) {
	if app.name == "" {

	}

	if app.vote == nil {

	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		[]byte(app.vote.String()),
	})

	if err != nil {
		return nil, err
	}

	return createDelete(
		*pHash,
		app.name,
		app.vote,
	), nil
}
