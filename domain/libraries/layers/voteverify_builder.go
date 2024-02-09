package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type voteVerifyBuilder struct {
	hashAdapter hash.Adapter
	vote        string
	message     string
	hashedRing  string
}

func createVoteVerifyBuilder(
	hashAdapter hash.Adapter,
) VoteVerifyBuilder {
	out := voteVerifyBuilder{
		hashAdapter: hashAdapter,
		vote:        "",
		message:     "",
		hashedRing:  "",
	}

	return &out
}

// Create initializes the builder
func (app *voteVerifyBuilder) Create() VoteVerifyBuilder {
	return createVoteVerifyBuilder(
		app.hashAdapter,
	)
}

// WithVote adds a vote to the builder
func (app *voteVerifyBuilder) WithVote(vote string) VoteVerifyBuilder {
	app.vote = vote
	return app
}

// WithMessage adds a message to the builder
func (app *voteVerifyBuilder) WithMessage(msg string) VoteVerifyBuilder {
	app.message = msg
	return app
}

// WithHashedRing adds an hashed ring to the builder
func (app *voteVerifyBuilder) WithHashedRing(hashedRing string) VoteVerifyBuilder {
	app.hashedRing = hashedRing
	return app
}

// Now builds a new VoteVerify instance
func (app *voteVerifyBuilder) Now() (VoteVerify, error) {
	if app.vote == "" {
		return nil, errors.New("the vote variable is mandatory in order to build a VoteVerify instance")
	}

	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a VoteVerify instance")
	}

	if app.hashedRing == "" {
		return nil, errors.New("the hashedRing variable is mandatory in order to build a VoteVerify instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.vote),
		[]byte(app.message),
		[]byte(app.hashedRing),
	})

	if err != nil {
		return nil, err
	}

	return createVoteVerify(*pHash, app.vote, app.message, app.hashedRing), nil
}
