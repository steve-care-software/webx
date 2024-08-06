package pointers

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/pointers"

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
}

// Pointer represents a pointer
type Pointer interface {
	Storage() pointers.Pointer
	Bytes() []byte
}
