package entries

type entry struct {
	kind      uint8
	content   []byte
	links     Links
	relations Relations
}

func createEntry(
	kind uint8,
	content []byte,
) Entry {
	return createEntryInternally(kind, content, nil, nil)
}

func createEntryWithLinks(
	kind uint8,
	content []byte,
	links Links,
) Entry {
	return createEntryInternally(kind, content, links, nil)
}

func createEntryWithRelations(
	kind uint8,
	content []byte,
	relations Relations,
) Entry {
	return createEntryInternally(kind, content, nil, relations)
}

func createEntryWithLinksAndRelations(
	kind uint8,
	content []byte,
	links Links,
	relations Relations,
) Entry {
	return createEntryInternally(kind, content, links, relations)
}

func createEntryInternally(
	kind uint8,
	content []byte,
	links Links,
	relations Relations,
) Entry {
	out := entry{
		kind:      kind,
		content:   content,
		links:     links,
		relations: relations,
	}

	return &out
}

// Kind returns the kind
func (obj *entry) Kind() uint8 {
	return obj.kind
}

// Content returns the content
func (obj *entry) Content() []byte {
	return obj.content
}

// HasLinks returns true if there is links, false otherwise
func (obj *entry) HasLinks() bool {
	return obj.links != nil
}

// Links returns the links, if any
func (obj *entry) Links() Links {
	return obj.links
}

// HasRelations returns true if there is relations, false otherwise
func (obj *entry) HasRelations() bool {
	return obj.relations != nil
}

// Relations returns the relations, if any
func (obj *entry) Relations() Relations {
	return obj.relations
}
