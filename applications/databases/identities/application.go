package identities

import (
	entry_applications "github.com/steve-care-software/webx/applications/databases/entries"
	"github.com/steve-care-software/webx/applications/databases/identities/selects"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/passwords"
	"github.com/steve-care-software/webx/domain/databases/entries"
	database_identity_modifications "github.com/steve-care-software/webx/domain/databases/identities/modifications"
	"github.com/steve-care-software/webx/domain/databases/references"
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

type application struct {
	encPasswordBuilder    passwords.Builder
	modificationAdapter   modifications.Adapter
	modificationDbAdapter database_identity_modifications.Adapter
	entryApplication      entry_applications.Application
	entriesBuilder        entries.Builder
	entryBuilder          entries.EntryBuilder
	selectAppBuilder      selects.Builder
}

func createApplication(
	encPasswordBuilder passwords.Builder,
	modificationAdapter modifications.Adapter,
	modificationDbAdapter database_identity_modifications.Adapter,
	entryApplication entry_applications.Application,
	entriesBuilder entries.Builder,
	entryBuilder entries.EntryBuilder,
	selectAppBuilder selects.Builder,
) Application {
	out := application{
		encPasswordBuilder:    encPasswordBuilder,
		modificationAdapter:   modificationAdapter,
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
	// retrieve the identities:
	names, err := app.entryApplication.List(references.KindIdentity)
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
	// convert the identity modifications to database identity modifications:
	modifications := identity.Modifications()
	dbModifications, err := app.modificationAdapter.ToDatabase(modifications)
	if err != nil {
		return err
	}

	// convert the database identity modifications to data:
	dbModificationContents, err := app.modificationDbAdapter.ToContents(dbModifications)
	if err != nil {
		return err
	}

	// build the encryption instance:
	encryption, err := app.encPasswordBuilder.Create().WithPassword(password).Now()
	if err != nil {
		return err
	}

	modificationEntriesList := []entries.Entry{}
	for _, oneDbModification := range dbModificationContents {
		// encrypt the modification:
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

	identityName := identity.Name()
	ins, err := app.entryBuilder.Create().
		WithKind(references.KindIdentity).
		WithContent([]byte(identityName)).
		WithRelations([]entries.Entries{
			modificationEntries,
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
