package databases

import "errors"

type entriesBuilder struct {
	list []Entry
}

func createEntriesBuilder() EntriesBuilder {
	out := entriesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *entriesBuilder) Create() EntriesBuilder {
	return createEntriesBuilder()
}

// WithList adds a list to the builder
func (app *entriesBuilder) WithList(list []Entry) EntriesBuilder {
	app.list = list
	return app
}

// Now builds a new Entries instance
func (app *entriesBuilder) Now() (Entries, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Entry in order to build an Entries instance")
	}

	return createEntries(app.list), nil
}
