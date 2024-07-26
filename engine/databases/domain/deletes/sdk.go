package deletes

import "github.com/steve-care-software/webx/engine/states/domain/databases/pointers"

// Builder represents a deletes builder
type Builder interface {
	Create() Builder
	WithList(list []Delete) Builder
	Now() (Deletes, error)
}

// Deletes represents a deletes
type Deletes interface {
	List() []Delete
}

// DeleteBuilder represents a delete builder
type DeleteBuilder interface {
	Create() DeleteBuilder
	WithKeyname(keyname string) DeleteBuilder
	Now() (Delete, error)
}

// Delete represents a delete pointer
type Delete interface {
	Keyname() string
	Pointer() pointers.Pointer
}
