package identities

import (
	entry_applications "github.com/steve-care-software/webx/applications/databases/entries"
	"github.com/steve-care-software/webx/applications/databases/identities/selects"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/passwords"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entries"
	database_identities "github.com/steve-care-software/webx/domain/databases/identities"
	database_identity_modifications "github.com/steve-care-software/webx/domain/databases/identities/modifications"
	"github.com/steve-care-software/webx/domain/databases/references"
	"github.com/steve-care-software/webx/domain/identities"
)

type application struct {
	hashAdapter           hash.Adapter
	encPasswordBuilder    passwords.Builder
	adapter               identities.Adapter
	dbAdapter             database_identities.Adapter
	modificationDbAdapter database_identity_modifications.Adapter
	entryApplication      entry_applications.Application
	entriesBuilder        entries.Builder
	entryBuilder          entries.EntryBuilder
	selectAppBuilder      selects.Builder
}

func createApplication(
	hashAdapter hash.Adapter,
	encPasswordBuilder passwords.Builder,
	adapter identities.Adapter,
	dbAdapter database_identities.Adapter,
	modificationDbAdapter database_identity_modifications.Adapter,
	entryApplication entry_applications.Application,
	entriesBuilder entries.Builder,
	entryBuilder entries.EntryBuilder,
	selectAppBuilder selects.Builder,
) Application {
	out := application{
		hashAdapter:           hashAdapter,
		encPasswordBuilder:    encPasswordBuilder,
		adapter:               adapter,
		dbAdapter:             dbAdapter,
		modificationDbAdapter: modificationDbAdapter,
		entryApplication:      entryApplication,
		entriesBuilder:        entriesBuilder,
		entryBuilder:          entryBuilder,
		selectAppBuilder:      selectAppBuilder,
	}

	return &out
}

// List lists the identities
func (app *application) List() ([]string, error) {
	// retrieve the identity names:
	names, err := app.entryApplication.List(references.KindIdentityName)
	if err != nil {
		return nil, err
	}

	list := []string{}
	for _, oneName := range names {
		list = append(list, string(oneName))
	}

	return list, nil
}

// New creates then save the new identity
func (app *application) New(identity identities.Identity, password []byte) error {
	// convert the identity instance to a database identity and modifications:
	dbIdentity, dbModifications, err := app.adapter.ToDatabase(identity)
	if err != nil {
		return err
	}

	// convert the database identity instance to data:
	dbIdentityContent, err := app.dbAdapter.ToContent(dbIdentity)
	if err != nil {
		return err
	}

	// convert the database identity modifications to data:
	dbModificationContents, err := app.modificationDbAdapter.ToContents(dbModifications)
	if err != nil {
		return err
	}

	// encrypt the content using the password:
	encryption, err := app.encPasswordBuilder.Create().WithPassword(password).Now()
	if err != nil {
		return err
	}

	modificationEntriesList := []entries.Entry{}
	for _, oneDbModification := range dbModificationContents {
		modifCipher, err := encryption.Encrypt(oneDbModification)
		if err != nil {
			return err
		}

		ins, err := app.entryBuilder.Create().
			WithKind(references.KindIdentityModification).
			WithContent(modifCipher).
			Now()

		if err != nil {
			return err
		}

		modificationEntriesList = append(modificationEntriesList, ins)
	}

	modificationEntries, err := app.entriesBuilder.Create().
		WithList(modificationEntriesList).
		Now()

	if err != nil {
		return err
	}

	identityCipher, err := encryption.Encrypt(dbIdentityContent)
	if err != nil {
		return err
	}

	identityName := identity.Name()
	nameEntry, err := app.entryBuilder.Create().
		WithKind(references.KindIdentityName).
		WithContent([]byte(identityName)).
		Now()

	if err != nil {
		return err
	}

	ins, err := app.entryBuilder.Create().
		WithKind(references.KindIdentity).
		WithContent(identityCipher).
		WithRelations([]entries.Entries{
			modificationEntries,
		}).
		WithLinks([]entries.Entry{
			nameEntry,
		}).
		Now()

	if err != nil {
		return err
	}

	return app.entryApplication.Insert(ins)
}

// Select returns the selected identity application
func (app *application) Select(name string) (selects.Application, error) {
	return app.selectAppBuilder.Create().
		WithName(name).
		Now()
}
