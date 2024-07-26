package entries

import (
	"github.com/steve-care-software/webx/engine/databases/domain/headers/containers/pointers"
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

type Repository interface {
	Retrieve(pointer pointers.Pointer) ([]byte, error)
	RetrieveAll(pointers pointers.Pointers) ([][]byte, error)
}

// Service represents an instance service
type Service interface {
	Insert(entry Entry) error
	InsertAll(entries Entries) error
	Delete(pointer pointers.Pointer) error
	DeleteAll(pointers pointers.Pointers) error
}
