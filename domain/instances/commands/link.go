package commands

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links"
)

type link struct {
	hash    hash.Hash
	input   []byte
	link    links.Link
	command Command
}

func createLink(
	hash hash.Hash,
	input []byte,
	linkIns links.Link,
) Link {
	return createLinkInternally(hash, input, linkIns, nil)
}

func createLinkWithCommand(
	hash hash.Hash,
	input []byte,
	linkIns links.Link,
	command Command,
) Link {
	return createLinkInternally(hash, input, linkIns, command)
}

func createLinkInternally(
	hash hash.Hash,
	input []byte,
	linkIns links.Link,
	command Command,
) Link {
	out := link{
		hash:    hash,
		input:   input,
		link:    linkIns,
		command: command,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *link) Input() []byte {
	return obj.input
}

// Link returns the link
func (obj *link) Link() links.Link {
	return obj.link
}

// HasCommand returns true if there is a command, false otherwise
func (obj *link) HasCommand() bool {
	return obj.command != nil
}

// Command returns the command
func (obj *link) Command() Command {
	return obj.command
}
