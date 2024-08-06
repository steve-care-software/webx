package validates

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	vote        string
	message     string
	hashedRing  string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		vote:        "",
		message:     "",
		hashedRing:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithVote adds a vote to the builder
func (app *builder) WithVote(vote string) Builder {
	app.vote = vote
	return app
}

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// WithHashedRing adds an hashed ring to the builder
func (app *builder) WithHashedRing(hashedRing string) Builder {
	app.hashedRing = hashedRing
	return app
}

// Now builds a new Validate instance
func (app *builder) Now() (Validate, error) {
	if app.vote == "" {
		return nil, errors.New("the vote is mandatory in order to build a Validate instance")
	}

	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Validate instance")
	}

	if app.hashedRing == "" {
		return nil, errors.New("the hashedRing is mandatory in order to build a Validate instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.vote),
		[]byte(app.message),
		[]byte(app.hashedRing),
	})

	if err != nil {
		return nil, err
	}

	return createValidate(*pHash, app.vote, app.message, app.hashedRing), nil
}
