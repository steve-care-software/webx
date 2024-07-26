package contexts

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type context struct {
	hash       hash.Hash
	identifier uint
	executions []hash.Hash
	head       hash.Hash
}

func createContext(
	hash hash.Hash,
	identifier uint,
	executions []hash.Hash,
) Context {
	return createContextInternally(hash, identifier, executions, nil)
}

func createContextWithHead(
	hash hash.Hash,
	identifier uint,
	executions []hash.Hash,
	head hash.Hash,
) Context {
	return createContextInternally(hash, identifier, executions, head)
}

func createContextInternally(
	hash hash.Hash,
	identifier uint,
	executions []hash.Hash,
	head hash.Hash,
) Context {
	out := context{
		hash:       hash,
		identifier: identifier,
		executions: executions,
		head:       head,
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

// Executions returns the executions
func (obj *context) Executions() []hash.Hash {
	return obj.executions
}

// HasHead returns true if there is a head, false otherwise
func (obj *context) HasHead() bool {
	return obj.head != nil
}

// Head returns the head
func (obj *context) Head() hash.Hash {
	return obj.head
}
