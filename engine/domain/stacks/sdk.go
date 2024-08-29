package stacks

import (
	"github.com/steve-care-software/webx/engine/domain/hash"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
)

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
	Height() uint
	HasParent() bool
	Parent() Stack
}
