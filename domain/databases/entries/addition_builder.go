package entries

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type additionBuilder struct {
	entry     Entry
	links     []hash.Hash
	relations [][]hash.Hash
}

func createAdditionBuilder() AdditionBuilder {
	out := additionBuilder{
		entry:     nil,
		links:     nil,
		relations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *additionBuilder) Create() AdditionBuilder {
	return createAdditionBuilder()
}

// WithEntry adds an entry to the builder
func (app *additionBuilder) WithEntry(entry Entry) AdditionBuilder {
	app.entry = entry
	return app
}

// WithLinks add links to the builder
func (app *additionBuilder) WithLinks(links []hash.Hash) AdditionBuilder {
	app.links = links
	return app
}

// WithRelations add relations to the builder
func (app *additionBuilder) WithRelations(relations [][]hash.Hash) AdditionBuilder {
	app.relations = relations
	return app
}

// Now builds a new Addition instance
func (app *additionBuilder) Now() (Addition, error) {
	if app.entry == nil {

	}

	if app.links != nil && len(app.links) <= 0 {
		app.links = nil
	}

	if app.relations != nil && len(app.relations) <= 0 {
		app.relations = nil
	}

	if app.relations != nil {
		for idx, oneList := range app.relations {
			if oneList != nil && len(oneList) <= 0 {
				oneList = nil
			}

			if oneList == nil {
				str := fmt.Sprintf("the relation (index: %d) was expected to contain an Hash list, but was empty", idx)
				return nil, errors.New(str)
			}
		}
	}

	if app.links != nil && app.relations != nil {
		return createAdditionWithLinksAndRelations(app.entry, app.links, app.relations), nil
	}

	if app.links != nil {
		return createAdditionWithLinks(app.entry, app.links), nil
	}

	if app.relations != nil {
		return createAdditionWithRelations(app.entry, app.relations), nil
	}

	return nil, errors.New("the Addition is invalid")
}
