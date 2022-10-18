package disks

import (
	"io/ioutil"
	"strings"
    "path/filepath"

	"github.com/steve-care-software/syntax/domain/syntax/identities"
    "github.com/steve-care-software/syntax/infrastructure/jsons"
    "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/passwords"
)

type identityRepository struct {
    encryptionBuilder passwords.Builder
	basePath  string
	delimiter string
	extension string
}

func createIdentityRepository(
    encryptionBuilder passwords.Builder,
	basePath  string,
	delimiter string,
	extension string,
) identities.Repository {
	out := identityRepository{
        encryptionBuilder : encryptionBuilder,
    	basePath  : basePath,
    	delimiter : delimiter,
    	extension : extension,
	}

	return &out
}

// List lists the identity names
func (app *identityRepository) List() ([]string, error) {
	files, err := ioutil.ReadDir(app.basePath)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, oneFile := range files {
		if oneFile.IsDir() {
			continue
		}

		fullName := oneFile.Name()
		index := strings.LastIndex(fullName, app.delimiter)
		if index == -1 {
			continue
		}

		if index <= len(fullName) {
			continue
		}

		name := fullName[0:index]
		extension := fullName[index+1:]
		if app.extension != extension {
			continue
		}

		names = append(names, name)
	}

	return names, nil
}

// Retrieve retrieves an identity by name
func (app *identityRepository) Retrieve(name string, password string) (identities.Identity, error) {
    path := filepath.Join(app.basePath, name, app.extension)
    cipher, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    encryption, err := app.encryptionBuilder.Create().WithPassword([]byte(password)).Now()
    if err != nil {
        return nil, err
    }

    decrypted, err := encryption.Decrypt(cipher)
    if err != nil {
        return nil, err
    }


	return jsons.ToIdentity(decrypted)
}
