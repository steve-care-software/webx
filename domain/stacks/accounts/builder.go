package accounts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

type builder struct {
	accountIns  accounts.Account
	credentials credentials.Credentials
	ring        []signers.PublicKey
	signature   signers.Signature
	vote        signers.Vote
}

func createBuilder() Builder {
	out := builder{
		accountIns:  nil,
		credentials: nil,
		ring:        nil,
		signature:   nil,
		vote:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithAccount adds an account to the builder
func (app *builder) WithAccount(account accounts.Account) Builder {
	app.accountIns = account
	return app
}

// WithCredentials add credentials to the builder
func (app *builder) WithCredentials(credentials credentials.Credentials) Builder {
	app.credentials = credentials
	return app
}

// WithRing add ring to the builder
func (app *builder) WithRing(ring []signers.PublicKey) Builder {
	app.ring = ring
	return app
}

// WithSignature add signature to the builder
func (app *builder) WithSignature(signature signers.Signature) Builder {
	app.signature = signature
	return app
}

// WithVote add vote to the builder
func (app *builder) WithVote(vote signers.Vote) Builder {
	app.vote = vote
	return app
}

// Now builds a new Account instance
func (app *builder) Now() (Account, error) {
	if app.accountIns != nil {
		return createAccountWithAccount(app.accountIns), nil
	}

	if app.credentials != nil {
		return createAccountWithCredentials(app.credentials), nil
	}

	if app.ring != nil && len(app.ring) <= 0 {
		app.ring = nil
	}

	if app.ring != nil {
		return createAccountWithRing(app.ring), nil
	}

	if app.signature != nil {
		return createAccountWithSignature(app.signature), nil
	}

	if app.vote != nil {
		return createAccountWithVote(app.vote), nil
	}

	return nil, errors.New("The Account is invalid")
}
