package entries

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type addition struct {
	entry     Entry
	links     []hash.Hash
	relations [][]hash.Hash
}

func createAdditionWithLinks(
	entry Entry,
	links []hash.Hash,
) Addition {
	return createAdditionInternally(entry, links, nil)
}

func createAdditionWithRelations(
	entry Entry,
	relations [][]hash.Hash,
) Addition {
	return createAdditionInternally(entry, nil, relations)
}

func createAdditionWithLinksAndRelations(
	entry Entry,
	links []hash.Hash,
	relations [][]hash.Hash,
) Addition {
	return createAdditionInternally(entry, links, relations)
}

func createAdditionInternally(
	entry Entry,
	links []hash.Hash,
	relations [][]hash.Hash,
) Addition {
	out := addition{
		entry:     entry,
		links:     links,
		relations: relations,
	}

	return &out
}

// Entry returns the entry
func (obj *addition) Entry() Entry {
	return obj.entry
}

// HasLinks returns true if there is links, false otherwise
func (obj *addition) HasLinks() bool {
	return obj.links != nil
}

// Links returns the links, if any
func (obj *addition) Links() []hash.Hash {
	return obj.links
}

// HasRelations returns true if there is relations, false otherwise
func (obj *addition) HasRelations() bool {
	return obj.relations != nil
}

// Relations returns the relations, if any
func (obj *addition) Relations() [][]hash.Hash {
	return obj.relations
}
