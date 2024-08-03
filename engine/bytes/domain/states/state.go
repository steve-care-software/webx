package states

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
)

type state struct {
	isDeleted bool
	root      delimiters.Delimiter
	pointers  pointers.Pointers
}

func createState(
	isDeleted bool,
) State {
	return createStateInternally(isDeleted, nil, nil)
}

func createStateWithRoot(
	isDeleted bool,
	root delimiters.Delimiter,
) State {
	return createStateInternally(isDeleted, root, nil)
}

func createStateWithPointers(
	isDeleted bool,
	pointers pointers.Pointers,
) State {
	return createStateInternally(isDeleted, nil, pointers)
}

func createStateWithRootAndPointers(
	isDeleted bool,
	root delimiters.Delimiter,
	pointers pointers.Pointers,
) State {
	return createStateInternally(isDeleted, root, pointers)
}

func createStateInternally(
	isDeleted bool,
	root delimiters.Delimiter,
	pointers pointers.Pointers,
) State {
	out := state{
		isDeleted: isDeleted,
		root:      root,
		pointers:  pointers,
	}

	return &out
}

// IsDeleted returns true if deleted, false otherwise
func (obj *state) IsDeleted() bool {
	return obj.isDeleted
}

// HasRoot returns true if there is a root, false otherwise
func (obj *state) HasRoot() bool {
	return obj.root != nil
}

// Root returns the root, if any
func (obj *state) Root() delimiters.Delimiter {
	return obj.root
}

// HasPointers returns true if there is pointers, false otherwise
func (obj *state) HasPointers() bool {
	return obj.pointers != nil
}

// Pointers returns the pointers, if any
func (obj *state) Pointers() pointers.Pointers {
	return obj.pointers
}
