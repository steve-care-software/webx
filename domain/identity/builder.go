package identity

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

type builder struct {
	pID     *uuid.UUID
	public  publics.Public
	pk      signatures.PrivateKey
	wallets wallets.Wallets
}

func createBuilder() Builder {
	out := builder{
		pID:     nil,
		public:  nil,
		pk:      nil,
		wallets: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id uuid.UUID) Builder {
	app.pID = &id
	return app
}

// WithPublic adds a public to the builder
func (app *builder) WithPublic(public publics.Public) Builder {
	app.public = public
	return app
}

// WithPrivateKey adds a pk to the builder
func (app *builder) WithPrivateKey(pk signatures.PrivateKey) Builder {
	app.pk = pk
	return app
}

// WithWallets add wallets to the builder
func (app *builder) WithWallets(wallets wallets.Wallets) Builder {
	app.wallets = wallets
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.pID == nil {
		return nil, errors.New("the ID is mandatory in order to build an Identity instance")
	}

	if app.public == nil {
		return nil, errors.New("the public is mandatory in order to build an Identity instance")
	}

	if app.pk == nil {
		return nil, errors.New("the pk is mandatory in order to build an Identity instance")
	}

	if app.wallets != nil {
		return createIdentityWithWallets(*app.pID, app.public, app.pk, app.wallets), nil
	}

	return createIdentity(*app.pID, app.public, app.pk), nil
}
