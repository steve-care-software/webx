package entries

import "github.com/steve-care-software/webx/domain/cryptography/hash"

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
	WithKind(kind uint8) EntryBuilder
	WithContent(content []byte) EntryBuilder
	WithLinks(links []Entry) EntryBuilder
	WithRelations(relations []Entries) EntryBuilder
	Now() (Entry, error)
}

// Entry represents an entry
type Entry interface {
	Kind() uint8
	Content() []byte
	HasLinks() bool
	Links() []Entry
	HasRelations() bool
	Relations() []Entries
}

// AdditionEntryBuilder represents an addition entry builder
type AdditionEntryBuilder interface {
	Create() AdditionEntryBuilder
	WithEntry(entry Entry) AdditionEntryBuilder
	WithLinks(links []hash.Hash) AdditionEntryBuilder
	WithRelations(relations [][]hash.Hash) AdditionEntryBuilder
	Now() (AdditionEntry, error)
}

// AdditionEntry represents an addition entry
type AdditionEntry interface {
	Entry() Entry
	HasLinks() bool
	Links() []hash.Hash
	HasRelations() bool
	Relations() [][]hash.Hash
}
