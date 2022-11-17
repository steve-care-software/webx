package selects

import (
	entry_applications "github.com/steve-care-software/webx/applications/databases/entries"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/passwords"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entries"
	database_identity_modifications "github.com/steve-care-software/webx/domain/databases/identities/modifications"
	"github.com/steve-care-software/webx/domain/databases/references"
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

type application struct {
	hashAdapter           hash.Adapter
	encPasswordBuilder    passwords.Builder
	identityBuilder       identities.Builder
	modificationDbAdapter database_identity_modifications.Adapter
	modificationsAdapter  modifications.Adapter
	modificationAdapter   modifications.ModificationAdapter
	entryApplication      entry_applications.Application
	entryBuilder          entries.EntryBuilder
	additionEntryBuilder  entries.AdditionEntryBuilder
	name                  string
}

func createApplication(
	hashAdapter hash.Adapter,
	encPasswordBuilder passwords.Builder,
	identityBuilder identities.Builder,
	modificationDbAdapter database_identity_modifications.Adapter,
	modificationsAdapter modifications.Adapter,
	modificationAdapter modifications.ModificationAdapter,
	entryApplication entry_applications.Application,
	entryBuilder entries.EntryBuilder,
	additionEntryBuilder entries.AdditionEntryBuilder,
	name string,
) Application {
	out := application{
		hashAdapter:           hashAdapter,
		encPasswordBuilder:    encPasswordBuilder,
		identityBuilder:       identityBuilder,
		modificationDbAdapter: modificationDbAdapter,
		modificationsAdapter:  modificationsAdapter,
		modificationAdapter:   modificationAdapter,
		entryApplication:      entryApplication,
		entryBuilder:          entryBuilder,
		additionEntryBuilder:  additionEntryBuilder,
		name:                  name,
	}

	return &out
}

// Retrieve retrieves an application
func (app *application) Retrieve(password []byte) (identities.Identity, error) {
	// retrieve the identity:
	pHash, err := app.hashAdapter.FromBytes([]byte(app.name))
	if err != nil {
		return nil, err
	}

	// retrieve the modifications:
	modificationCiphersList, err := app.entryApplication.ListByLink(references.KindIdentity, *pHash, references.KindIdentityModification)
	if err != nil {
		return nil, err
	}

	// build the encryption instance with the password:
	encryption, err := app.encPasswordBuilder.Create().WithPassword(password).Now()
	if err != nil {
		return nil, err
	}

	decryptedModificationContents := [][]byte{}
	for _, oneModificationCipher := range modificationCiphersList {
		modificationContent, err := encryption.Decrypt(oneModificationCipher)
		if err != nil {
			return nil, err
		}

		decryptedModificationContents = append(decryptedModificationContents, modificationContent)
	}

	dbModifications, err := app.modificationDbAdapter.ToModifications(decryptedModificationContents)
	if err != nil {
		return nil, err
	}

	modifications, err := app.modificationsAdapter.ToModification(dbModifications)
	if err != nil {
		return nil, err
	}

	return app.identityBuilder.Create().
		WithModifications(modifications).
		Now()
}

// Modify modify an identity
func (app *application) Modify(modification modifications.Modification, currentPassword []byte, newPassword []byte) error {
	if newPassword == nil {
		newPassword = currentPassword
	}

	identity, err := app.Retrieve(currentPassword)
	if err != nil {
		return err
	}

	// convert the modification to a db modification:
	dbModification, err := app.modificationAdapter.ToDatabase(modification)
	if err != nil {
		return err
	}

	// convert the database identity modifications to data:
	dbModificationContent, err := app.modificationDbAdapter.ToContent(dbModification)
	if err != nil {
		return err
	}

	// build the encryption instance with the new password:
	encryption, err := app.encPasswordBuilder.Create().WithPassword(newPassword).Now()
	if err != nil {
		return err
	}

	// encrypt the modification:
	modifCipher, err := encryption.Encrypt(dbModificationContent)
	if err != nil {
		return err
	}

	modificationEntry, err := app.entryBuilder.Create().
		WithKind(references.KindIdentityModification).
		WithContent(modifCipher).
		Now()

	if err != nil {
		return err
	}

	// build the addition entry:
	pHash, err := app.hashAdapter.FromBytes([]byte(identity.Name()))
	if err != nil {
		return err
	}

	addition, err := app.additionEntryBuilder.Create().
		WithEntity(*pHash).
		WithEntry(modificationEntry).
		Now()

	if err != nil {
		return err
	}

	return app.entryApplication.Add(addition)
}
