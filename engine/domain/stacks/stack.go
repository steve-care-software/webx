package stacks

import (
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

type stack struct {
	frame  frames.Frame
	parent Stack
}

func createStack(
	frame frames.Frame,
) Stack {
	return createStackInternally(frame, nil)
}

func createStackWithParent(
	frame frames.Frame,
	parent Stack,
) Stack {
	return createStackInternally(frame, parent)
}

func createStackInternally(
	frame frames.Frame,
	parent Stack,
) Stack {
	out := stack{
		frame:  frame,
		parent: parent,
	}

	return &out
}

// Frame returns the frame
func (obj *stack) Frame() frames.Frame {
	return obj.frame
}

// Fetch fetches a variable by name
func (obj *stack) Fetch(name string) (variables.Variable, error) {
	retVariable, err := obj.frame.Fetch(name)
	if err != nil {
		if obj.HasParent() {
			return obj.parent.Fetch(name)
		}

		return nil, err
	}

	return retVariable, nil
}

// Height returns the stack height
func (obj *stack) Height() uint {
	if !obj.HasParent() {
		return 1
	}

	return obj.parent.Height() + 1
}

// HasParent returns true if there is a parent, false otherwise
func (obj *stack) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *stack) Parent() Stack {
	return obj.parent
}
