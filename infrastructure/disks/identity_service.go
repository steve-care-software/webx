package disks

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/passwords"
	"github.com/steve-care-software/syntax/domain/syntax/identities"
	"github.com/steve-care-software/syntax/infrastructure/jsons"
)

type identityService struct {
	encryptionBuilder passwords.Builder
	repository        identities.Repository
	basePath          string
	delimiter         string
	extension         string
}

func createIdentityService(
	encryptionBuilder passwords.Builder,
	repository identities.Repository,
	basePath string,
	delimiter string,
	extension string,
) identities.Service {
	out := identityService{
		encryptionBuilder: encryptionBuilder,
		repository:        repository,
		basePath:          basePath,
		delimiter:         delimiter,
		extension:         extension,
	}

	return &out
}

// Insert inserts a new identity
func (app *identityService) Insert(identity identities.Identity, password string) error {
	name := identity.Name()
	_, err := app.repository.Retrieve(name, password)
	if err == nil {
		str := fmt.Sprintf("the identity (name: %s) cannot be inserted because it already exists", name)
		return errors.New(str)
	}

	decrypted, err := jsons.ToJSON(identity)
	if err != nil {
		return err
	}

	encryption, err := app.encryptionBuilder.Create().WithPassword([]byte(password)).Now()
	if err != nil {
		return err
	}

	cipher, err := encryption.Encrypt(decrypted)
	if err != nil {
		return err
	}

	path := filepath.Join(app.basePath, name, app.extension)
	return os.WriteFile(path, cipher, 0666)
}

// Update updates an existing identity
func (app *identityService) Update(identity identities.Identity, currentPassword string, newPassword string) error {
	name := identity.Name()
	_, err := app.repository.Retrieve(name, currentPassword)
	if err != nil {
		str := fmt.Sprintf("the identity (name: %s) cannot be updated because it doesn't already exists", name)
		return errors.New(str)
	}

	decrypted, err := jsons.ToJSON(identity)
	if err != nil {
		return err
	}

	encryption, err := app.encryptionBuilder.Create().WithPassword([]byte(newPassword)).Now()
	if err != nil {
		return err
	}

	cipher, err := encryption.Encrypt(decrypted)
	if err != nil {
		return err
	}

	path := filepath.Join(app.basePath, name, app.extension)
	return os.WriteFile(path, cipher, 0666)
}
