package commands

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/receipts/commands/results"
	"github.com/steve-care-software/identity/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewCommandBuilder creates a new command builder
func NewCommandBuilder() CommandBuilder {
	hashAdapter := hash.NewAdapter()
	return createCommandBuilder(
		hashAdapter,
	)
}

// NewLinkBuilder creates a new link builder
func NewLinkBuilder() LinkBuilder {
	hashAdapter := hash.NewAdapter()
	return createLinkBuilder(
		hashAdapter,
	)
}

// Builder represents a commands builder
type Builder interface {
	Create() Builder
	WithList(list []Command) Builder
	Now() (Commands, error)
}

// Commands represents commands
type Commands interface {
	Hash() hash.Hash
	List() []Command
	Last() Command
}

// CommandBuilder represents a command builder
type CommandBuilder interface {
	Create() CommandBuilder
	WithInput(input []byte) CommandBuilder
	WithLayer(layer layers.Layer) CommandBuilder
	WithResult(result results.Result) CommandBuilder
	WithParent(parent Link) CommandBuilder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	Hash() hash.Hash
	Input() []byte
	Layer() layers.Layer
	Result() results.Result
	HasParent() bool
	Parent() Link
}

// LinkBuilder represents a link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithInput(input []byte) LinkBuilder
	WithLink(link links.Link) LinkBuilder
	WithCommand(command Command) LinkBuilder
	Now() (Link, error)
}

// Link represents a link execution
type Link interface {
	Hash() hash.Hash
	Input() []byte
	Link() links.Link
	Command() Command
}
