package updates

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/criterias"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	repository       accounts.Repository
	service          accounts.Service
	accountBuilder   accounts.Builder
	criteriaBuilder  criterias.Builder
	signerFactory    signers.Factory
	encryptorBuilder encryptors.Builder
	bitRate          int
}

func createApplication(
	repository accounts.Repository,
	service accounts.Service,
	accountBuilder accounts.Builder,
	criteriaBuilder criterias.Builder,
	signerFactory signers.Factory,
	encryptorBuilder encryptors.Builder,
	bitRate int,
) Application {
	out := application{
		repository:       repository,
		service:          service,
		accountBuilder:   accountBuilder,
		criteriaBuilder:  criteriaBuilder,
		signerFactory:    signerFactory,
		encryptorBuilder: encryptorBuilder,
		bitRate:          bitRate,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, instruction updates.Update) (*uint, error) {
	credentialsVar := instruction.Credentials()
	credentials, err := frame.FetchCredentials(credentialsVar)
	if err != nil {
		code := failures.CouldNotFetchCredentialsFromFrame
		str := fmt.Sprintf("the variable (name: %s) was expected to contain credentials, but was NOT declared", credentialsVar)
		return &code, errors.New(str)
	}

	insCriteria := instruction.Criteria()
	builder := app.criteriaBuilder.Create()
	if insCriteria.ChangeSigner() {
		signer := app.signerFactory.Create()
		builder.WithSigner(signer)
	}

	if insCriteria.ChangeEncryptor() {
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

		builder.WithEncryptor(encryptor)
	}

	if insCriteria.HasUsername() {
		userVar := insCriteria.Username()
		userBytes, err := frame.FetchBytes(userVar)
		if err != nil {
			code := failures.CouldNotFetchUsernameFromFrame
			return &code, err
		}

		username := string(userBytes)
		exists := app.repository.Exists(username)
		if exists {
			code := failures.AccountWithSameUsernameAlreadyExists
			str := fmt.Sprintf("the account (name: %s) already exists", username)
			return &code, errors.New(str)
		}

		builder.WithUsername(username)
	}

	if insCriteria.HasPassword() {
		passVar := insCriteria.Password()
		password, err := frame.FetchBytes(passVar)
		if err != nil {
			code := failures.CouldNotFetchPasswordFromFrame
			return &code, err
		}

		builder.WithPassword(password)
	}

	criteria, err := builder.Now()
	if err != nil {
		return nil, err
	}

	err = app.service.Update(credentials, criteria)
	if err != nil {
		code := failures.CouldNotUpdateAccountInDatabase
		return &code, err
	}

	return nil, nil
}
