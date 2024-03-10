package commands

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type commands struct {
	hash hash.Hash
	list []Command
}

func createCommands(
	hash hash.Hash,
	list []Command,
) Commands {
	out := commands{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *commands) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *commands) List() []Command {
	return obj.list
}

// Last returns the last command
func (obj *commands) Last() Command {
	return obj.list[len(obj.list)-1]
}
