package contexts

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type context struct {
	hash       hash.Hash
	identifier uint
	head       hash.Hash
	executions []hash.Hash
}

func createContext(
	hash hash.Hash,
	identifier uint,
	head hash.Hash,
	executions []hash.Hash,
) Context {
	out := context{
		hash:       hash,
		identifier: identifier,
		head:       head,
		executions: executions,
	}

	return &out
}

// Hash returns the hash
func (obj *context) Hash() hash.Hash {
	return obj.hash
}

// Identifier returns the identifier
func (obj *context) Identifier() uint {
	return obj.identifier
}

// Head returns the head
func (obj *context) Head() hash.Hash {
	return obj.head
}

// Executions returns the executions
func (obj *context) Executions() []hash.Hash {
	return obj.executions
}
