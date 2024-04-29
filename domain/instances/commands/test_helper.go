package commands

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
)

// NewCommandsForTests creates a new commands for tests
func NewCommandsForTests(list []Command) Commands {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommandForTests creates a new command for tests
func NewCommandForTests(input []byte, layer layers.Layer, result results.Result, parent Link) Command {
	ins, err := NewCommandBuilder().Create().WithInput(input).WithLayer(layer).WithResult(result).WithParent(parent).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates a new link for tests
func NewLinkForTests(input []byte, link links.Link) Link {
	ins, err := NewLinkBuilder().Create().WithInput(input).WithLink(link).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithCommandForTests creates a new link with command for tests
func NewLinkWithCommandForTests(input []byte, link links.Link, command Command) Link {
	ins, err := NewLinkBuilder().Create().WithInput(input).WithLink(link).WithCommand(command).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
