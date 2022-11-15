package entries

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
	WithWeightedRelations(weightedRelations []WeightedEntries) EntryBuilder
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
	HasWeightedRelations() bool
	WeightedRelations() []WeightedEntries
}

// WeightedEntriesBuilder represents a weighted entries builder
type WeightedEntriesBuilder interface {
	Create() WeightedEntriesBuilder
	WithList(list []WeightedEntry) WeightedEntriesBuilder
	Now() (WeightedEntries, error)
}

// WeightedEntries represents weighted entries
type WeightedEntries interface {
	List() []WeightedEntry
}

// WeightedEntryBuilder represents a weighted entry builder
type WeightedEntryBuilder interface {
	Create() WeightedEntryBuilder
	WithEntry(entry Entry) WeightedEntryBuilder
	WithWeight(weight uint) WeightedEntryBuilder
	Now() (WeightedEntry, error)
}

// WeightedEntry represents a weighted entry
type WeightedEntry interface {
	Entry() Entry
	Weight() uint
}
