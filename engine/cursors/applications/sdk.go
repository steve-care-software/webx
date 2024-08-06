package applications

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/records"
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

const zeroRecordErr = "there are currently no record"

const (
	// NamespaceFlag represents the namespace flag
	NamespaceFlag (uint8) = iota

	// IdentityFlag represents the identy flag
	IdentityFlag

	// BlockchainFlag represents the blockchain flag
	BlockchainFlag
)

// Application represents an application
type Application interface {
	// cursors:
	Cursor() (cursors.Cursor, error)   // returns the current cursor
	Records() (records.Records, error) // returns the recorded cursors
	Erase(name string) error           // erase a cursor using its name
	Record(name string) error          // record the cursor to this name
	Replace(name string) error         // put the cursor to what the name was pointing to

	// metadata:
	MetaData() (delimiters.Delimiter, error) // returns the current branch meta data

	// install:
	InstallHeader(header headers.Header) error
	InstallVersions(versions versions.Versions) error
	InstallWorkspaces(workspaces workspaces.Workspaces) error
	InstallRootBranch(rootBRanch branches.Branch) error
	InstallBranches(branches branches.Branches) error
	InstallStates(states states.States) error

	// identitity:
	Create(original originals.Original, password []byte) error
	Authenticate(name string, password []byte) error
	SetPassword(newPassword []byte) error // update the password of the authenticated user
	Encrypt(message []byte) ([]byte, error)
	Decrypt(cipher []byte) ([]byte, error)
	Sign(message []byte) (signers.Signature, error)
	ValidateSignature(message []byte, sig signers.Signature) (bool, error)
	Vote(message []byte, ring []signers.PublicKey) (signers.Vote, error)
	ValidateVote(message []byte, vote signers.Vote, ring []hash.Hash) (bool, error)

	// currency:
	Transfer(toWallet hash.Hash, amount uint64, fees uint64) error
	Lock(walletPassword []byte, toWallet hash.Hash, untilBlock uint64) error
	Claim(lockPassword []byte, amount uint64, amountSeed []byte) error

	// switch between namespace, identity and blockchain
	Switch(flag uint8) error

	// write generics:
	Set(name string) error   // set the cursor to this element (horizontally)
	Down(name string) error  // set the cursor to this element (sub-element or 'down')
	Climb(name string) error // set the cursor to this element (parent-element or 'climb')
	Insert(original originals.Original) error
	Update(original string, updated originals.Original) error
	Delete(name string) error
	Recover(name string) error
	Purge(name string) error
	PurgeAll() error
	Move(name string, devName string, deleteOriginal bool) error // moves a development iteration to a production iteration inside the current iteration
	Merge(deleteOriginal bool) error

	// data:
	NextIndex() (*uint, error) // returns the next index for data
	InsertData(delimiter delimiters.Delimiter) error
	UpdateData(original delimiters.Delimiter, updated []byte) error
	DeleteData(delete delimiters.Delimiter) error

	// injects a transaction:
	Transact(trx transactions.Transaction) error

	// execute:
	Execute() (transactions.Transaction, error)
}
