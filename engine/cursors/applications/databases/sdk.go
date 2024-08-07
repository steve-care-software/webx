package databases

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"

// FileBuilder represents a file application builder
type FileBuilder interface {
	Create() FileBuilder
	WithPath(path []string) FileBuilder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Reset()
	Read(delimiter delimiters.Delimiter) ([]byte, error)
	CopyBeforeThenWrite(index uint64, bytes []byte) error
	Close() error
}
