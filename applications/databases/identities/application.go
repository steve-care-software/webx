package identities

import (
	"github.com/steve-care-software/webx/applications/databases/entries"
	"github.com/steve-care-software/webx/applications/databases/identities/selects"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/passwords"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases"
	database_identities "github.com/steve-care-software/webx/domain/databases/identities"
	"github.com/steve-care-software/webx/domain/identities"
)

type application struct {
	hashAdapter        hash.Adapter
	encPasswordBuilder passwords.Builder
	adapter            identities.Adapter
	dbAdapter          database_identities.Adapter
	entryBuilder       databases.EntryBuilder
	entryApplication   entries.Application
	database           databases.Database
}

// List lists the identities
func (app *application) List() ([]string, error) {
	// fetch the identity name sections:

	// retrieve all entries:

	// for each entry, convert the data to a string:

	// returns the list:
	return nil, nil
}

// New creates then save the new identity
func (app *application) New(identity identities.Identity, password []byte) error {
	// convert the identity instance to a database identity:
	dbIdentity, err := app.adapter.ToDatabase(identity)
	if err != nil {
		return err
	}

	// convert the database identity instance to data:
	content, err := app.dbAdapter.ToContent(dbIdentity)
	if err != nil {
		return err
	}

	// encrypt the content using the password:
	encryption, err := app.encPasswordBuilder.Create().WithPassword(password).Now()
	if err != nil {
		return err
	}

	cipher, err := encryption.Encrypt(content)
	if err != nil {
		return err
	}

	// build an entry instance:
	entry, err := app.entryBuilder.Create().
		WithContent(cipher).
		WithKind(databases.KindIdentity).
		Now()

	if err != nil {
		return err
	}

	// save the entry instances:
	_, err = app.entryApplication.Insert(entry)
	if err != nil {
		return err
	}

	return nil
}

// Select returns the selected identity application
func (app *application) Select(name string) (selects.Application, error) {
	return nil, nil
}
