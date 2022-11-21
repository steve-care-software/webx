package entries

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/steve-care-software/webx/domain/databases/entries"
)

type builder struct {
	builder          entries.Builder
	entryBuilder     entries.EntryBuilder
	relationsBuilder entries.RelationsBuilder
	relationBuilder  entries.RelationBuilder
	linksBuilder     entries.LinksBuilder
	linkBuilder      entries.LinkBuilder
	additionBuilder  entries.AdditionBuilder
	basePath         string
	relativeFilePath string
}

func createBuilder(
	entriesBuilder entries.Builder,
	entryBuilder entries.EntryBuilder,
	relationsBuilder entries.RelationsBuilder,
	relationBuilder entries.RelationBuilder,
	linksBuilder entries.LinksBuilder,
	linkBuilder entries.LinkBuilder,
	additionBuilder entries.AdditionBuilder,
	basePath string,
) Builder {
	out := builder{
		builder:          entriesBuilder,
		entryBuilder:     entryBuilder,
		relationsBuilder: relationsBuilder,
		relationBuilder:  relationBuilder,
		linksBuilder:     linksBuilder,
		linkBuilder:      linkBuilder,
		additionBuilder:  additionBuilder,
		basePath:         basePath,
		relativeFilePath: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.builder,
		app.entryBuilder,
		app.relationsBuilder,
		app.relationBuilder,
		app.linksBuilder,
		app.linkBuilder,
		app.additionBuilder,
		app.basePath,
	)
}

// WithRelativeFilePath adds a relative filePath to the builder
func (app *builder) WithRelativeFilePath(relFilePath string) Builder {
	return createBuilder(
		app.builder,
		app.entryBuilder,
		app.relationsBuilder,
		app.relationBuilder,
		app.linksBuilder,
		app.linkBuilder,
		app.additionBuilder,
		app.basePath,
	)
}

// Now builds a new builder instance
func (app *builder) Now() (Application, error) {
	if app.relativeFilePath == "" {
		return nil, errors.New("the relative filePath is mandatory in order to build an entry file disk Application")
	}

	filePath, err := filepath.Abs(filepath.Join(app.basePath, app.relativeFilePath))
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(filePath, app.basePath) {
		str := fmt.Sprintf("the relative path (%s) must be contained within the base path (%s)", app.relativeFilePath, app.basePath)
		return nil, errors.New(str)
	}

	// read the database flags in the database file:

	// read the reference content:

	// convert the content to a reference instance:

	return createApplication(
		app.builder,
		app.entryBuilder,
		app.relationsBuilder,
		app.relationBuilder,
		app.linksBuilder,
		app.linkBuilder,
		app.additionBuilder,
		nil,
		filePath,
	), nil
}
