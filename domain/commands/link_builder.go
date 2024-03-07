package commands

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	input       []byte
	link        links.Link
	command     Command
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		input:       nil,
		link:        nil,
		command:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder(
		app.hashAdapter,
	)
}

// WithInput adds input to the builder
func (app *linkBuilder) WithInput(input []byte) LinkBuilder {
	app.input = input
	return app
}

// WithLink adds link to the builder
func (app *linkBuilder) WithLink(link links.Link) LinkBuilder {
	app.link = link
	return app
}

// WithCommand adds command to the builder
func (app *linkBuilder) WithCommand(command Command) LinkBuilder {
	app.command = command
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Link instance")
	}

	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Link instance")
	}

	if app.command == nil {
		return nil, errors.New("the command is mandatory in order to build a Link instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.input,
		app.link.Hash().Bytes(),
		app.command.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*pHash, app.input, app.link, app.command), nil
}
