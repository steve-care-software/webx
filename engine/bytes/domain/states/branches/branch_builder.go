package branches

import (
	"errors"

	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers"
)

type branchBuilder struct {
	name      string
	isDeleted bool
	layers    layers.Layers
	metaData  delimiters.Delimiter
	children  Branches
}

func createBranchBuilder() BranchBuilder {
	out := branchBuilder{
		name:      "",
		isDeleted: false,
		layers:    nil,
		metaData:  nil,
		children:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *branchBuilder) Create() BranchBuilder {
	return createBranchBuilder()
}

// WithName adds a name to the builder
func (app *branchBuilder) WithName(name string) BranchBuilder {
	app.name = name
	return app
}

// WithLayers add layers to the builder
func (app *branchBuilder) WithLayers(layers layers.Layers) BranchBuilder {
	app.layers = layers
	return app
}

// WithMetaData add metaData to the builder
func (app *branchBuilder) WithMetaData(metaData delimiters.Delimiter) BranchBuilder {
	app.metaData = metaData
	return app
}

// WithChildren add children to the builder
func (app *branchBuilder) WithChildren(children Branches) BranchBuilder {
	app.children = children
	return app
}

// IsDeleted flags the builder as deleted
func (app *branchBuilder) IsDeleted() BranchBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new Branch instance
func (app *branchBuilder) Now() (Branch, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a branch instance")
	}

	if app.layers != nil && app.metaData != nil && app.children != nil {
		return createBranchWithLayersAndMetaDataAndChildren(
			app.name,
			app.isDeleted,
			app.layers,
			app.metaData,
			app.children,
		), nil
	}

	if app.layers != nil && app.children != nil {
		return createBranchWithLayersAndChildren(
			app.name,
			app.isDeleted,
			app.layers,
			app.children,
		), nil
	}

	if app.layers != nil && app.metaData != nil {
		return createBranchWithLayersAndMetaData(
			app.name,
			app.isDeleted,
			app.layers,
			app.metaData,
		), nil
	}

	if app.metaData != nil && app.children != nil {
		return createBranchWithMetaDataAndChildren(
			app.name,
			app.isDeleted,
			app.metaData,
			app.children,
		), nil
	}

	if app.layers != nil {
		return createBranchWithLayers(
			app.name,
			app.isDeleted,
			app.layers,
		), nil
	}

	if app.metaData != nil {
		return createBranchWithMetaData(
			app.name,
			app.isDeleted,
			app.metaData,
		), nil
	}

	if app.children != nil {
		return createBranchWithChildren(
			app.name,
			app.isDeleted,
			app.children,
		), nil
	}

	return createBranch(app.name, app.isDeleted), nil
}
