package applications

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/branches"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/versions"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/workspaces"
	"github.com/steve-care-software/webx/engine/cursors/domain/transactions"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Cursor returns the current cursor
func (app *application) Cursor() (cursors.Cursor, error) {
	return nil, nil
}

// Records returns the recorded cursors
func (app *application) Records() (cursors.Cursors, error) {
	return nil, nil
}

// Erase erases a cursor using the provided name
func (app *application) Erase(name string) error {
	return nil
}

// Record records the current cursor to this name
func (app *application) Record(name string) error {
	return nil
}

// Replace replaces the current cursor by the stored cursor of the provided name
func (app *application) Replace(name string) error {
	return nil
}

// MetaData returns the current branch's metaData
func (app *application) MetaData() (delimiters.Delimiter, error) {
	return nil, nil
}

// InstallHeader install the header
func (app *application) InstallHeader(header headers.Header) error {
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

// Switch switches the head cursor between namespace, identity and blockchain
func (app *application) Switch(flag uint8) error {
	return nil
}

// Set sets the cursor to this element (horizontally)
func (app *application) Set(name string) error {
	return nil
}

// Down sets the cursor to this element (sub-element or 'down')
func (app *application) Down(name string) error {
	return nil
}

// Climb sets the cursor to this element (parent-element or 'climb')
func (app *application) Climb(name string) error {
	return nil
}

// Insert inserts an original
func (app *application) Insert(original originals.Original) error {
	return nil
}

// Update updates an original
func (app *application) Update(original string, updated originals.Original) error {
	return nil
}

// Delete deletes an original by name
func (app *application) Delete(name string) error {
	return nil
}

// Recover recovers an original by name
func (app *application) Recover(name string) error {
	return nil
}

// Purge purges by name
func (app *application) Purge(name string) error {
	return nil
}

// PurgeAll purges all
func (app *application) PurgeAll() error {
	return nil
}

// Move moves a development iteration to a production iteration inside the current iteration
func (app *application) Move(name string, devName string, deleteOriginal bool) error {
	return nil
}

// Merge merges the current branch to its parent
func (app *application) Merge(deleteOriginal bool) error {
	return nil
}

// NextIndex returns the next index of data
func (app *application) NextIndex() (*uint, error) {
	return nil, nil
}

// InsertData inserts data to the current state
func (app *application) InsertData(delimiter delimiters.Delimiter) error {
	return nil
}

// UpdateData updates data on the current state
func (app *application) UpdateData(original delimiters.Delimiter, updated []byte) error {
	return nil
}

// DeleteData deletes data from the current state
func (app *application) DeleteData(delete delimiters.Delimiter) error {
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
