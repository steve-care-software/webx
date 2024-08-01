package states

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"

type state struct {
	isDeleted bool
	pointers  pointers.Pointers
}

func createState() State {
	return createStateInternally(false, nil)
}

func createStateWithPointers(
	pointers pointers.Pointers,
) State {
	return createStateInternally(false, pointers)
}

func createStateWithDeleted(
	pointers pointers.Pointers,
) State {
	return createStateInternally(true, pointers)
}

func createStateWithPointersAndDeleted(
	pointers pointers.Pointers,
) State {
	return createStateInternally(true, pointers)
}

func createStateInternally(
	isDeleted bool,
	pointers pointers.Pointers,
) State {
	out := state{
		isDeleted: isDeleted,
		pointers:  pointers,
	}

	return &out
}

// IsDeleted returns true if deleted, false otherwise
func (obj *state) IsDeleted() bool {
	return obj.isDeleted
}

// HasPointers returns true if there is pointers, false otherwise
func (obj *state) HasPointers() bool {
	return obj.pointers != nil
}

// Pointers returns the pointers, if any
func (obj *state) Pointers() pointers.Pointers {
	return obj.pointers
}
