package applications

import (
	"github.com/steve-care-software/webx/databases/domain/configs"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
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
	Content
}

// Reference represents the reference application
type Reference interface {
	Content(context uint) (references.Content, error)
	Commits(context uint) (references.Commits, error)
}

// Database represents the database application
type Database interface {
	Delete(name string) error
	Open(name string, height int) (*uint, error)
	Cancel(context uint) error
	Commit(context uint) error
	Push(context uint, config configs.Config) error
	Close(context uint) error
}

// Content represents the content application
type Content interface {
	Read(context uint, pointer references.Pointer) ([]byte, error)
	ReadByHash(context uint, hash hash.Hash) ([]byte, error)
	ReadAll(context uint, pointers []references.Pointer) ([][]byte, error)
	ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error)
	Write(context uint, hash hash.Hash, data []byte, kind uint) error
}
