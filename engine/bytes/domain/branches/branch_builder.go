package branches

import (
	"errors"

	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
)

type branchBuilder struct {
	name        string
	description string
	isDeleted   bool
	states      delimiters.Delimiter
	metaData    delimiters.Delimiter
	children    Branches
}

func createBranchBuilder() BranchBuilder {
	out := branchBuilder{
		name:        "",
		description: "",
		isDeleted:   false,
		states:      nil,
		metaData:    nil,
		children:    nil,
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

// WithDescription adds a description to the builder
func (app *branchBuilder) WithDescription(description string) BranchBuilder {
	app.description = description
	return app
}

// WithStates add states to the builder
func (app *branchBuilder) WithStates(states delimiters.Delimiter) BranchBuilder {
	app.states = states
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

	if app.states != nil && app.metaData != nil && app.children != nil {
		return createBranchWithStatesAndMetaDataAndChildren(
			app.name,
			app.description,
			app.isDeleted,
			app.states,
			app.metaData,
			app.children,
		), nil
	}

	if app.states != nil && app.children != nil {
		return createBranchWithStatesAndChildren(
			app.name,
			app.description,
			app.isDeleted,
			app.states,
			app.children,
		), nil
	}

	if app.states != nil && app.metaData != nil {
		return createBranchWithStatesAndMetaData(
			app.name,
			app.description,
			app.isDeleted,
			app.states,
			app.metaData,
		), nil
	}

	if app.metaData != nil && app.children != nil {
		return createBranchWithMetaDataAndChildren(
			app.name,
			app.description,
			app.isDeleted,
			app.metaData,
			app.children,
		), nil
	}

	if app.states != nil {
		return createBranchWithStates(
			app.name,
			app.description,
			app.isDeleted,
			app.states,
		), nil
	}

	if app.metaData != nil {
		return createBranchWithMetaData(
			app.name,
			app.description,
			app.isDeleted,
			app.metaData,
		), nil
	}

	if app.children != nil {
		return createBranchWithChildren(
			app.name,
			app.description,
			app.isDeleted,
			app.children,
		), nil
	}

	return createBranch(
		app.name,
		app.description,
		app.isDeleted,
	), nil
}
