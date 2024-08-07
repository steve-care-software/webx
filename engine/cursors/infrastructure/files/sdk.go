package files

import "github.com/steve-care-software/webx/engine/cursors/applications/databases"

const (
	// BeginFlag represents the begin flag
	BeginFlag (uint8) = iota

	// EndFlag represents the end flag
	EndFlag

	// CurrentFlag represents the current flag
	CurrentFlag
)

const readChunkSize = uint64(10)

// NewDatabaseApplicationBuilder creates a new file database application builder
func NewDatabaseApplicationBuilder() databases.FileBuilder {
	return createDatabaseApplicationBuilder()
}

// Application represents a file application
type Application interface {
	Open(path string) (*uint, error)
	Length(identifier uint) (*uint64, error)
	Seek(identifier uint, flag uint8) error
	Read(identifier uint, index uint64, length uint64) ([]byte, error)
	ReadAll(identifier uint) ([]byte, error)
	Close(identifier uint) error
}
