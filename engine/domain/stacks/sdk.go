package stacks

import (
	"github.com/steve-care-software/webx/engine/domain/hash"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

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
	Hash() hash.Hash
	Frame() frames.Frame
	Fetch(name string) (variables.Variable, error)
	Height() uint
	HasParent() bool
	Parent() Stack
}
