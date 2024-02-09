package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type voteBuilder struct {
	hashAdapter hash.Adapter
	ring        string
	message     string
}

func createVoteBuilder(
	hashAdapter hash.Adapter,
) VoteBuilder {
	out := voteBuilder{
		hashAdapter: hashAdapter,
		ring:        "",
		message:     "",
	}

	return &out
}

// Create initializes the builder
func (app *voteBuilder) Create() VoteBuilder {
	return createVoteBuilder(
		app.hashAdapter,
	)
}

// WithRing adds a ring to the builder
func (app *voteBuilder) WithRing(ring string) VoteBuilder {
	app.ring = ring
	return app
}

// WithMessage adds a message to the builder
func (app *voteBuilder) WithMessage(message string) VoteBuilder {
	app.message = message
	return app
}

// Now builds a new Vote instance
func (app *voteBuilder) Now() (Vote, error) {
	if app.ring == "" {
		return nil, errors.New("the ring variable is mandatory in order to build a Vote instance")
	}

	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Vote instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.ring),
		[]byte(app.message),
	})

	if err != nil {
		return nil, err
	}

	return createVote(*pHash, app.ring, app.message), nil
}
