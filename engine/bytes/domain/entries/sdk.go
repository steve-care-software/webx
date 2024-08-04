package entries

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
)

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewEntryBuilder creates a new entry builder
func NewEntryBuilder() EntryBuilder {
	return createEntryBuilder()
}

// Builder represents an entries builder
type Builder interface {
	Create() Builder
	WithList(list []Entry) Builder
	Now() (Entries, error)
}

// Entries represents entries
type Entries interface {
	List() []Entry
}

// EntryBuilder represents an entry builder
type EntryBuilder interface {
	Create() EntryBuilder
	WithDelimiter(delimiter delimiters.Delimiter) EntryBuilder
	WithBytes(bytes []byte) EntryBuilder
	Now() (Entry, error)
}

// Entry represents an entry
type Entry interface {
	Delimiter() delimiters.Delimiter
	Bytes() []byte
}
