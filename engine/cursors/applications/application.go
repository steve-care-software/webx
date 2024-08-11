package applications

import (
	"errors"

	applications_loaders "github.com/steve-care-software/webx/engine/cursors/applications/loaders"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	loaders_namespace "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces"
	"github.com/steve-care-software/webx/engine/cursors/domain/records"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/branches"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/versions"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/workspaces"
)

type application struct {
	loaderApp              applications_loaders.Application
	loaderFactory          loaders.Factory
	loaderBuilder          loaders.Builder
	loaderNamespaceBuilder loaders_namespace.Builder
	recordsBuilder         records.Builder
	recordBuilder          records.RecordBuilder
	cursorFactory          cursors.CursorFactory
	cursorsBuilder         cursors.Builder
	cursorBuilder          cursors.CursorBuilder
	currentCursor          cursors.Cursor
	currentLoader          loaders.Loader
	records                records.Records
}

func createApplication(
	loaderApp applications_loaders.Application,
	loaderFactory loaders.Factory,
	loaderBuilder loaders.Builder,
	recordsBuilder records.Builder,
	recordBuilder records.RecordBuilder,
	cursorFactory cursors.CursorFactory,
	cursorsBuilder cursors.Builder,
	cursorBuilder cursors.CursorBuilder,
	initialLoader loaders.Loader,
	initialCursor cursors.Cursor,
) Application {
	out := application{
		loaderApp:      loaderApp,
		loaderFactory:  loaderFactory,
		loaderBuilder:  loaderBuilder,
		recordsBuilder: recordsBuilder,
		recordBuilder:  recordBuilder,
		cursorFactory:  cursorFactory,
		cursorsBuilder: cursorsBuilder,
		cursorBuilder:  cursorBuilder,
		currentCursor:  initialCursor,
		currentLoader:  initialLoader,
		records:        nil,
	}

	return &out
}

// Cursor returns the current cursor
func (app *application) Cursor() (cursors.Cursor, error) {
	return app.currentCursor, nil
}

// Records returns the recorded cursors
func (app *application) Records() (records.Records, error) {
	if app.records != nil {
		return app.records, nil
	}

	return nil, errors.New(zeroRecordErr)
}

// Erase erases a cursor using the provided name
func (app *application) Erase(name string) error {
	if app.records == nil {
		return errors.New(zeroRecordErr)
	}

	list, err := app.records.FetchExceptName(name)
	if err != nil {
		return err
	}

	if len(list) <= 0 {
		app.records = nil
		return nil
	}

	updated, err := app.recordsBuilder.Create().WithList(list).Now()
	if err != nil {
		return err
	}

	app.records = updated
	return nil
}

// Record records the current cursor to this name
func (app *application) Record(name string) error {
	record, err := app.recordBuilder.Create().
		WithName(name).
		WithCursor(app.currentCursor).
		Now()

	if err != nil {
		return err
	}

	list := []records.Record{
		record,
	}

	if app.records != nil {
		currentList := app.records.List()
		list = append(list, currentList...)
	}

	records, err := app.recordsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return err
	}

	app.records = records
	return nil
}

// Replace replaces the current cursor by the stored cursor of the provided name
func (app *application) Replace(name string) error {
	if app.records == nil {
		return errors.New(zeroRecordErr)
	}

	record, err := app.records.FetchByName(name)
	if err != nil {
		return err
	}

	app.currentCursor = record.Cursor()
	return nil
}

// MetaData returns the current branch's metaData
func (app *application) MetaData() (delimiters.Delimiter, error) {
	return nil, nil
}

// InstallHeader install the header
func (app *application) InstallHeader(header headers.Header) error {
	/*builder := app.loaderBuilder.Create()
	if app.currentLoader != nil {
		builder.WithInitialLoader(app.currentLoader)
	}*/

	/*if header.HasNamespaces() {
		headerNamespaces := header.Namespaces()
		namespace, err := app.loaderNamespaceBuilder.Create().WithAll(headerNamespaces).Now()
		if err != nil {
			return err
		}

		builder.WithNamespace(namespace)
	}*/

	return nil
}

// InstallVersions installs versions
func (app *application) InstallVersions(versions versions.Versions) error {
	return nil
}

// InstallWorkspaces installs workspaces
func (app *application) InstallWorkspaces(workspaces workspaces.Workspaces) error {
	return nil
}

// InstallRootBranch installs the root branch
func (app *application) InstallRootBranch(rootBRanch branches.Branch) error {
	return nil
}

// InstallBranches installs branches
func (app *application) InstallBranches(branches branches.Branches) error {
	return nil
}

// InstallStates installs states
func (app *application) InstallStates(states states.States) error {
	return nil
}

// Create creates an identity
func (app *application) Create(original originals.Original, password []byte) error {
	return nil
}

// Authenticate authenticated to an identity
func (app *application) Authenticate(name string, password []byte) error {
	return nil
}

// SetPassword updates the password of the authenticated user
func (app *application) SetPassword(newPassword []byte) error {
	return nil
}

// Encrypt encrypts a message and returns the cipher
func (app *application) Encrypt(message []byte) ([]byte, error) {
	return nil, nil
}

// Decrypt decrypts a cipher and returns the message
func (app *application) Decrypt(cipher []byte) ([]byte, error) {
	return nil, nil
}

// Sign signs a message
func (app *application) Sign(message []byte) (signers.Signature, error) {
	return nil, nil
}

// ValidateSignature validates a signature
func (app *application) ValidateSignature(message []byte, sig signers.Signature) (bool, error) {
	return false, nil
}

// Vote votes on a message and the provided ring
func (app *application) Vote(message []byte, ring []signers.PublicKey) (signers.Vote, error) {
	return nil, nil
}

// ValidateVote validates a vote
func (app *application) ValidateVote(message []byte, vote signers.Vote, ring []hash.Hash) (bool, error) {
	return false, nil
}

// Transfer transfers currencies to another wallet
func (app *application) Transfer(toWallet hash.Hash, amount uint64, fees uint64) error {
	return nil
}

// Lock locks a transfer until a given block height
func (app *application) Lock(walletPassword []byte, toWallet hash.Hash, untilBlock uint64) error {
	return nil
}

// Claim claims a lock
func (app *application) Claim(lockPassword []byte, amount uint64, amountSeed []byte) error {
	return nil
}

// Reset resets the cursor and set it at namespace, identity and blockchain
func (app *application) Reset(flag uint8) error {
	cursor, err := app.cursorBuilder.Create().WithFlag(flag).Now()
	if err != nil {
		return err
	}

	app.currentCursor = cursor
	return nil
}

// Set sets the cursor to this element (horizontally)
func (app *application) Set(name string) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Set(name)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Down sets the cursor to this element (sub-element or 'down')
func (app *application) Down(name string) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Down(name)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Climb sets the cursor to this element (parent-element or 'climb')
func (app *application) Climb(name string) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Climb(name)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Insert inserts an original
func (app *application) Insert(original originals.Original) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Insert(original)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Update updates an original
func (app *application) Update(original string, updated originals.Original) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Update(original, updated)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Delete deletes an original by name
func (app *application) Delete(name string) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Delete(name)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Recover recovers an original by name
func (app *application) Recover(name string) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Recover(name)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Purge purges by name
func (app *application) Purge(name string) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Purge(name)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// PurgeAll purges all
func (app *application) PurgeAll() error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.PurgeAll()
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Move moves a development iteration to a production iteration inside the current iteration
func (app *application) Move(name string, devName string, deleteOriginal bool) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Move(name, devName, deleteOriginal)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Merge merges the current branch to its parent
func (app *application) Merge(deleteOriginal bool) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.Merge(deleteOriginal)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// NextIndex returns the next index of data
func (app *application) NextIndex() (*uint, error) {
	return nil, nil
}

// InsertData inserts data to the current state
func (app *application) InsertData(delimiter delimiters.Delimiter) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.InsertData(delimiter)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// UpdateData updates data on the current state
func (app *application) UpdateData(original delimiters.Delimiter, updated []byte) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.UpdateData(original, updated)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// DeleteData deletes data from the current state
func (app *application) DeleteData(delete delimiters.Delimiter) error {
	if app.currentLoader == nil {
		return errors.New(noLoaderCreatedErr)
	}

	loader, err := app.loaderApp.DeleteData(delete)
	if err != nil {
		return err
	}

	app.currentLoader = loader
	return nil
}

// Transact executes a transaction
func (app *application) Transact(trx transactions.Transaction) error {
	return nil
}

// Execute executes the cursor and returns a transaction
func (app *application) Execute() (transactions.Transaction, error) {
	return nil, nil
}
