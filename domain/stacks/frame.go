package stacks

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type frame struct {
	assignments Assignments
}

func createFrame() Frame {
	return createFrameInternally(
		nil,
	)
}

func createFrameWithAssignments(
	assignments Assignments,
) Frame {
	return createFrameInternally(
		assignments,
	)
}

func createFrameInternally(
	assignments Assignments,
) Frame {
	out := frame{
		assignments: assignments,
	}

	return &out
}

// Fetch fetches an assignable by name
func (obj *frame) Fetch(name string) (Assignable, error) {
	if !obj.HasAssignments() {
		str := fmt.Sprintf("the frame contains no assignment, therefore the assignable (name: %s) could not be found", name)
		return nil, errors.New(str)
	}

	return obj.Assignments().Fetch(name)
}

// FetchBool fetches a bool by name
func (obj *frame) FetchBool(name string) (bool, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return false, err
	}

	if !assignable.IsBool() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a bool", name)
		return false, errors.New(str)
	}

	pBool := assignable.Bool()
	return *pBool, nil
}

// FetchHash fetches an hash by name
func (obj *frame) FetchHash(name string) (hash.Hash, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsHash() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Hash", name)
		return nil, errors.New(str)
	}

	return assignable.Hash(), nil
}

// FetchBytes fetches a bytes by name
func (obj *frame) FetchBytes(name string) ([]byte, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsBytes() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a []byte", name)
		return nil, errors.New(str)
	}

	return assignable.Bytes(), nil
}

// HasAssignments returns true if there is assignments, false otherwise
func (obj *frame) HasAssignments() bool {
	return obj.assignments != nil
}

// Assignments fetches the assignments, if any
func (obj *frame) Assignments() Assignments {
	return obj.assignments
}
