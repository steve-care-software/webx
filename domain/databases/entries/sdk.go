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
	WithLinks(links Links) EntryBuilder
	WithRelations(relations Relations) EntryBuilder
	Now() (Entry, error)
}

// Entry represents an entry
type Entry interface {
	Kind() uint8
	Content() []byte
	HasLinks() bool
	Links() Links
	HasRelations() bool
	Relations() Relations
}

// RelationsBuilder represents a relations builder
type RelationsBuilder interface {
	Create() RelationsBuilder
	WithList(list []Relation) RelationsBuilder
	Now() (Relations, error)
}

// Relations represents relations
type Relations interface {
	List() []Relation
}

// RelationBuilder represents a relation builder
type RelationBuilder interface {
	Create() RelationBuilder
	WithNew(new Entries) RelationBuilder
	WithExisting(existing []hash.Hash) RelationBuilder
	Now() (Relation, error)
}

// Relation represents a relation
type Relation interface {
	IsNew() bool
	New() Entries
	IsExisting() bool
	Existing() []hash.Hash
}

// LinksBuilder represents a links builder
type LinksBuilder interface {
	Create() LinksBuilder
	WithList(list []Link) LinksBuilder
	Now() (Links, error)
}

// Links represents links
type Links interface {
	List() []Link
}

// LinkBuilder represents a link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithNew(new Entry) LinkBuilder
	WithExisting(existing hash.Hash) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	IsNew() bool
	New() Entry
	IsExisting() bool
	Existing() *hash.Hash
}

// AdditionBuilder represents an addition entry builder
type AdditionBuilder interface {
	Create() AdditionBuilder
	WithEntry(entry Entry) AdditionBuilder
	WithLinks(links []hash.Hash) AdditionBuilder
	WithRelations(relations [][]hash.Hash) AdditionBuilder
	Now() (Addition, error)
}

// Addition represents an addition entry
type Addition interface {
	Entry() Entry
	HasLinks() bool
	Links() []hash.Hash
	HasRelations() bool
	Relations() [][]hash.Hash
}
