package votes

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	message     string
	ring        string
	account     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		message:     "",
		ring:        "",
		account:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// WithRing adds a ring to the builder
func (app *builder) WithRing(ring string) Builder {
	app.ring = ring
	return app
}

// WithAccount adds an account to the builder
func (app *builder) WithAccount(account string) Builder {
	app.account = account
	return app
}

// Now builds a new Vote instance
func (app *builder) Now() (Vote, error) {
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Vote instance")
	}

	if app.ring == "" {
		return nil, errors.New("the ring is mandatory in order to build a Vote instance")
	}

	if app.account == "" {
		return nil, errors.New("the account is mandatory in order to build a Vote instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.message),
		[]byte(app.ring),
		[]byte(app.account),
	})

	if err != nil {
		return nil, err
	}

	return createVote(*pHash, app.message, app.ring, app.account), nil
}
