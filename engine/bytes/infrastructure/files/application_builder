package files

import (
	"errors"

	"github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/bytes/domain/pointers"
	"github.com/steve-care-software/webx/engine/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type applicationBuilder struct {
	hashAdapter       hash.Adapter
	statesAdapter     states.Adapter
	statesBuilder     states.Builder
	stateBuilder      states.StateBuilder
	pointersBuilder   pointers.Builder
	pointerBuilder    pointers.PointerBuilder
	entriesBuilder    entries.Builder
	entryBuilder      entries.EntryBuilder
	delimitersBuilder delimiters.Builder
	delimiterBuilder  delimiters.DelimiterBuilder
	basepath          []string
}

func createApplicationBuilder(
	hashAdapter hash.Adapter,
	statesAdapter states.Adapter,
	statesBuilder states.Builder,
	stateBuilder states.StateBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	entriesBuilder entries.Builder,
	entryBuilder entries.EntryBuilder,
	delimitersBuilder delimiters.Builder,
	delimiterBuilder delimiters.DelimiterBuilder,
) applications.Builder {
	out := applicationBuilder{
		hashAdapter:       hashAdapter,
		statesAdapter:     statesAdapter,
		statesBuilder:     statesBuilder,
		stateBuilder:      stateBuilder,
		pointersBuilder:   pointersBuilder,
		pointerBuilder:    pointerBuilder,
		entriesBuilder:    entriesBuilder,
		entryBuilder:      entryBuilder,
		delimitersBuilder: delimitersBuilder,
		delimiterBuilder:  delimiterBuilder,
		basepath:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() applications.Builder {
	return createApplicationBuilder(
		app.hashAdapter,
		app.statesAdapter,
		app.statesBuilder,
		app.stateBuilder,
		app.pointersBuilder,
		app.pointerBuilder,
		app.entriesBuilder,
		app.entryBuilder,
		app.delimitersBuilder,
		app.delimiterBuilder,
	)
}

// WithBasePath adds a basePath to the builder
func (app *applicationBuilder) WithBasePath(basePath []string) applications.Builder {
	app.basepath = basePath
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (applications.Application, error) {
	if app.basepath != nil && len(app.basepath) <= 0 {
		app.basepath = nil
	}

	if app.basepath == nil {
		return nil, errors.New("the basePath is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.hashAdapter,
		app.statesAdapter,
		app.statesBuilder,
		app.stateBuilder,
		app.pointersBuilder,
		app.pointerBuilder,
		app.entriesBuilder,
		app.entryBuilder,
		app.delimitersBuilder,
		app.delimiterBuilder,
		app.basepath,
	), nil
}
