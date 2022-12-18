package applications

import (
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

const (
	// PendingContentFlag represents the pending content flag
	PendingContentFlag uint8 = iota

	// ActiveContentFlag represents the active content flag
	ActiveContentFlag

	// DeletedContentFlag represents the deleted content flag
	DeletedContentFlag
)

const (
	// ChainBlockchainFlag represents the chain blockchain flag
	ChainBlockchainFlag uint8 = iota

	// BlockBlockchainFlag represents the block blockchain flag
	BlockBlockchainFlag

	// TransactionBlockchainFlag represents the transaction blockchain flag
	TransactionBlockchainFlag
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents the read application
type Application interface {
	Database
	Reference
	Commit
	Content
}

// Commit represents the commit application
type Commit interface {
	Latest() (commits.Commit, error)
	Retrieve(hash hash.Hash) (commits.Commit, error)
}

// Reference represents the reference application
type Reference interface {
	ContentKeys(context uint, kind uint) (references.ContentKeys, error)
	ContentKeysByCommit(context uint, commit hash.Hash) (references.ContentKeys, error)
	ContentKey(context uint, hash hash.Hash, flag uint8) (references.ContentKey, error)
	Commits(context uint) (references.Commits, error)
}

// Database represents the database application
type Database interface {
	Delete(name string) error
	Open(name string, height int) (*uint, error)
	Cancel(context uint) error
	Commit(context uint) error
	Push(context uint) error
	Close(context uint) error
}

// Content represents the content application
type Content interface {
	Read(context uint, pointer references.Pointer) ([]byte, error)
	ReadByHash(content uint, hash hash.Hash) ([]byte, error)
	ReadAll(context uint, pointers []references.Pointer) ([][]byte, error)
	ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error)
	Write(context uint, hash hash.Hash, data []byte, kind uint) error
}
