package stacks

import (
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

// NewFactory creates a new factory
func NewFactory() Factory {
	builder := NewBuilder()
	frameFactory := frames.NewFactory()
	return createFactory(
		builder,
		frameFactory,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Factory represents a stack factory
type Factory interface {
	Create() (Stack, error)
}

// Builder represents a stack builder
type Builder interface {
	Create() Builder
	WithFrame(frame frames.Frame) Builder
	WithParent(parent Stack) Builder
	Now() (Stack, error)
}

// Stack represents a stack instance
type Stack interface {
	Frame() frames.Frame
	Fetch(name string) (variables.Variable, error)
	Height() uint
	HasParent() bool
	Parent() Stack
}
