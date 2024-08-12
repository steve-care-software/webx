package databases

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"

// FileBuilder represents a file application builder
type FileBuilder interface {
	Create() FileBuilder
	WithOriginal(original []string) FileBuilder
	WithDestination(destination []string) FileBuilder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Reset()
	Read(delimiter delimiters.Delimiter) ([]byte, error)
	CopyBeforeThenWrite(startAtIndex uint64, index uint64, bytes []byte) error
	CloseThenDeleteOriginal() error
	Close() error
}
