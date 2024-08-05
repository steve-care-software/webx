package states

import "github.com/steve-care-software/webx/engine/bytes/domain/pointers"

type state struct {
	message   string
	isDeleted bool
	pointers  pointers.Pointers
}

func createState(
	message string,
	isDeleted bool,
) State {
	return createStateInternally(message, isDeleted, nil)
}

func createStateWithPointers(
	message string,
	isDeleted bool,
	pointers pointers.Pointers,
) State {
	return createStateInternally(message, isDeleted, pointers)
}

func createStateInternally(
	message string,
	isDeleted bool,
	pointers pointers.Pointers,
) State {
	out := state{
		message:   message,
		isDeleted: isDeleted,
		pointers:  pointers,
	}

	return &out
}

// Message returns the message
func (obj *state) Message() string {
	return obj.message
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
