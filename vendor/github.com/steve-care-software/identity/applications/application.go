package applications

import (
	"github.com/steve-care-software/identity/domain/accounts"
	"github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/credentials"
)

type application struct {
	builder          accounts.Builder
	signerFactory    signers.Factory
	encryptorBuilder encryptors.Builder
	repository       accounts.Repository
	service          accounts.Service
	bitrate          int
}

func createApplication(
	builder accounts.Builder,
	signerFactory signers.Factory,
	encryptorBuilder encryptors.Builder,
	repository accounts.Repository,
	service accounts.Service,
	bitrate int,
) Application {
	out := application{
		builder:          builder,
		signerFactory:    signerFactory,
		encryptorBuilder: encryptorBuilder,
		repository:       repository,
		service:          service,
		bitrate:          bitrate,
	}

	return &out
}

// List returns the list of account usernames
func (app *application) List() ([]string, error) {
	return app.repository.List()
}

// Exists returns true if the username exists false otherwise
func (app *application) Exists(username string) (bool, error) {
	return app.repository.Exists(username)
}

// Insert inserts an account
func (app *application) Insert(credentials credentials.Credentials) error {
	username := credentials.Username()
	encryptor, err := app.encryptorBuilder.Create().
		WithBitRate(app.bitrate).
		Now()

	if err != nil {
		return err
	}

	signer := signers.NewFactory().Create()
	ins, err := app.builder.Create().
		WithEncryptor(encryptor).
		WithSigner(signer).
		WithUsername(username).
		Now()

	if err != nil {
		return err
	}

	password := credentials.Password()
	return app.service.Insert(ins, password)
}

// Retrieve retrieves an account
func (app *application) Retrieve(credentials credentials.Credentials) (accounts.Account, error) {
	return app.repository.Retrieve(credentials)
}

// Update updates an account
func (app *application) Update(credentials credentials.Credentials, criteria accounts.UpdateCriteria) error {
	return app.service.Update(credentials, criteria)
}

// Delete deletes an account
func (app *application) Delete(credentials credentials.Credentials) error {
	return app.service.Delete(credentials)
}
