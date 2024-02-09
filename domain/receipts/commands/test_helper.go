package commands

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/receipts/commands/results"
)

// NewCommandsForTests creates a new commands for tests
func NewCommandsForTests(list []Command) Commands {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommandWithParentForTests creates a new command with parent for tests
func NewCommandWithParentForTests(input []byte, layer layers.Layer, result results.Result, parent Link) Command {
	ins, err := NewCommandBuilder().Create().WithInput(input).WithLayer(layer).WithResult(result).WithParent(parent).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommandForTests creates a new command for tests
func NewCommandForTests(input []byte, layer layers.Layer, result results.Result) Command {
	ins, err := NewCommandBuilder().Create().WithInput(input).WithLayer(layer).WithResult(result).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates a new link for tests
func NewLinkForTests(input []byte, link links.Link, command Command) Link {
	ins, err := NewLinkBuilder().Create().WithInput(input).WithLink(link).WithCommand(command).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
