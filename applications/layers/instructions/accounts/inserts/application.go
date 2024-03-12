package inserts

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	repository       accounts.Repository
	service          accounts.Service
	accountBuilder   accounts.Builder
	signerFactory    signers.Factory
	encryptorBuilder encryptors.Builder
	bitRate          int
}

func createApplication(
	repository accounts.Repository,
	service accounts.Service,
	accountBuilder accounts.Builder,
	signerFactory signers.Factory,
	encryptorBuilder encryptors.Builder,
	bitRate int,
) Application {
	out := application{
		repository:       repository,
		service:          service,
		accountBuilder:   accountBuilder,
		signerFactory:    signerFactory,
		encryptorBuilder: encryptorBuilder,
		bitRate:          bitRate,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, instruction inserts.Insert) (*uint, error) {
	userVar := instruction.Username()
	userBytes, err := frame.FetchBytes(userVar)
	if err != nil {
		code := failures.CouldNotFetchUsernameFromFrame
		return &code, err
	}

	username := string(userBytes)
	exists, err := app.repository.Exists(username)
	if err != nil {
		return nil, err
	}

	if exists {
		code := failures.AccountWithSameUsernameAlreadyExists
		str := fmt.Sprintf("the account (name: %s) already exists", username)
		return &code, errors.New(str)
	}

	passVar := instruction.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		code := failures.CouldNotFetchPasswordFromFrame
		return &code, err
	}

	pPrivateKey, err := rsa.GenerateKey(rand.Reader, app.bitRate)
	if err != nil {
		return nil, err
	}

	encryptor, err := app.encryptorBuilder.Create().
		WithPK(*pPrivateKey).
		WithBitRate(app.bitRate).
		Now()

	if err != nil {
		return nil, err
	}

	signer := app.signerFactory.Create()
	account, err := app.accountBuilder.Create().
		WithUsername(username).
		WithEncryptor(encryptor).
		WithSigner(signer).
		Now()

	if err != nil {
		return nil, err
	}

	err = app.service.Insert(account, password)
	if err != nil {
		code := failures.CouldNotInsertAccountInDatabase
		return &code, err
	}

	return nil, nil
}
