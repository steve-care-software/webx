package pointers

import "github.com/steve-care-software/datastencil/domain/hash"

// Pointers represents pointers
type Pointers interface {
	Hash() hash.Hash
	List() []Pointer
}

// Pointer represents a pointer
type Pointer interface {
	Hash() hash.Hash
	Path() []string
	Identifier() hash.Hash
}
