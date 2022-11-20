package entries

import "errors"

type entryBuilder struct {
	pKind     *uint8
	content   []byte
	links     Links
	relations Relations
}

func createEntryBuilder() EntryBuilder {
	out := entryBuilder{
		pKind:     nil,
		content:   nil,
		links:     nil,
		relations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *entryBuilder) Create() EntryBuilder {
	return createEntryBuilder()
}

// WithKind adds a kind to the builder
func (app *entryBuilder) WithKind(kind uint8) EntryBuilder {
	app.pKind = &kind
	return app
}

// WithContent adds content to the builder
func (app *entryBuilder) WithContent(content []byte) EntryBuilder {
	app.content = content
	return app
}

// WithLinks add links to the builder
func (app *entryBuilder) WithLinks(links Links) EntryBuilder {
	app.links = links
	return app
}

// WithRelations add relations to the builder
func (app *entryBuilder) WithRelations(relations Relations) EntryBuilder {
	app.relations = relations
	return app
}

// Now builds a new Entry instance
func (app *entryBuilder) Now() (Entry, error) {
	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build an Entry instance")
	}

	if app.content != nil && len(app.content) <= 0 {
		app.content = nil
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Entry instance")
	}

	if app.links != nil && app.relations != nil {
		return createEntryWithLinksAndRelations(*app.pKind, app.content, app.links, app.relations), nil
	}

	if app.links != nil {
		return createEntryWithLinks(*app.pKind, app.content, app.links), nil
	}

	if app.relations != nil {
		return createEntryWithRelations(*app.pKind, app.content, app.relations), nil
	}

	return createEntry(*app.pKind, app.content), nil
}
