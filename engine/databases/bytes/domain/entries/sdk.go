package entries

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers/pointers"
)

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
	WithPointer(pointer pointers.Pointer) EntryBuilder
	WithBytes(bytes []byte) EntryBuilder
	Now() (Entry, error)
}

// Entry represents an entry
type Entry interface {
	Pointer() pointers.Pointer
	Bytes() []byte
}
