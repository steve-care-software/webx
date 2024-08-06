package pointers

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/pointers"

// Builder represents a pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
	NextIndex() (*uint, error)
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithStorage(storage pointers.Pointer) PointerBuilder
	WithBytes(bytes []byte) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Storage() pointers.Pointer
	Bytes() []byte
}