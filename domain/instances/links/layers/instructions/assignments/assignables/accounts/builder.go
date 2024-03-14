package accounts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
)

type builder struct {
	hashAdapter   hash.Adapter
	list          string
	credentials   credentials.Credentials
	retrieve      retrieves.Retrieve
	communication communications.Communication
	encryption    encryptions.Encryption
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		list:          "",
		credentials:   nil,
		retrieve:      nil,
		communication: nil,
		encryption:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list string) Builder {
	app.list = list
	return app
}

// WithCredentials add credentials to the builder
func (app *builder) WithCredentials(credentials credentials.Credentials) Builder {
	app.credentials = credentials
	return app
}

// WithRetrieve add retrieve to the builder
func (app *builder) WithRetrieve(retrieve retrieves.Retrieve) Builder {
	app.retrieve = retrieve
	return app
}

// WithCommunication add communication to the builder
func (app *builder) WithCommunication(communication communications.Communication) Builder {
	app.communication = communication
	return app
}

// WithEncryption add encryption to the builder
func (app *builder) WithEncryption(encryption encryptions.Encryption) Builder {
	app.encryption = encryption
	return app
}

// Now builds a new Account instance
func (app *builder) Now() (Account, error) {
	data := [][]byte{}
	if app.list != "" {
		data = append(data, []byte("list"))
		data = append(data, []byte(app.list))
	}

	if app.credentials != nil {
		data = append(data, []byte("credentials"))
		data = append(data, app.credentials.Hash().Bytes())
	}

	if app.retrieve != nil {
		data = append(data, []byte("retrieve"))
		data = append(data, app.retrieve.Hash().Bytes())
	}

	if app.communication != nil {
		data = append(data, []byte("communication"))
		data = append(data, app.communication.Hash().Bytes())
	}

	if app.encryption != nil {
		data = append(data, []byte("encryption"))
		data = append(data, app.encryption.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Account is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.list != "" {
		return createAccountWithList(*pHash, app.list), nil
	}

	if app.credentials != nil {
		return createAccountWithCredentials(*pHash, app.credentials), nil
	}

	if app.retrieve != nil {
		return createAccountWithRetrieve(*pHash, app.retrieve), nil
	}

	if app.communication != nil {
		return createAccountWithCommunication(*pHash, app.communication), nil
	}

	return createAccountWithEncryption(*pHash, app.encryption), nil
}
