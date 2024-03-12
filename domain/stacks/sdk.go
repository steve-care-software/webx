package stacks

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries"
	"github.com/steve-care-software/datastencil/domain/skeletons"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	framesBuilder := NewFramesBuilder()
	return createBuilder(
		framesBuilder,
	)
}

// NewFramesBuilder creates a new frames builder
func NewFramesBuilder() FramesBuilder {
	return createFramesBuilder()
}

// NewFrameBuilder creates a new frame builder
func NewFrameBuilder() FrameBuilder {
	return createFrameBuilder()
}

// NewAssignmentsBuilder creates a new assignments builder
func NewAssignmentsBuilder() AssignmentsBuilder {
	return createAssignmentsBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewAssignableBuilder creates a new assignable builder
func NewAssignableBuilder() AssignableBuilder {
	return createAssignableBuilder()
}

// Factory represents the stack factory
type Factory interface {
	Create() Stack
}

// Builder represents a stack builder
type Builder interface {
	Create() Builder
	WithFrames(frames Frames) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Frames() Frames
	Head() Frame
	HasBody() bool
	Body() Frames
}

// FramesBuilder represents the frames builder
type FramesBuilder interface {
	Create() FramesBuilder
	WithList(list []Frame) FramesBuilder
	Now() (Frames, error)
}

// Frames represents frames
type Frames interface {
	List() []Frame
}

// FrameBuilder represents the frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithAssignments(assignments Assignments) FrameBuilder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	Fetch(name string) (Assignable, error)
	FetchBool(name string) (bool, error)
	FetchHash(name string) (hash.Hash, error)
	FetchBytes(name string) ([]byte, error)
	FetchUnsignedInt(name string) (*uint, error)
	FetchString(name string) (string, error)
	FetchStringList(name string) ([]string, error)
	FetchAccount(name string) (accounts.Account, error)
	FetchRing(name string) ([]signers.PublicKey, error)
	FetchCredentials(name string) (credentials.Credentials, error)
	FetchInstance(name string) (instances.Instance, error)
	FetchQuery(name string) (queries.Query, error)
	HasAssignments() bool
	Assignments() Assignments
}

// AssignmentsBuilder represents an assignments builder
type AssignmentsBuilder interface {
	Create() AssignmentsBuilder
	WithList(list []Assignment) AssignmentsBuilder
	Now() (Assignments, error)
}

// Assignments represents assignments
type Assignments interface {
	List() []Assignment
	Fetch(name string) (Assignable, error)
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithAssignable(assignable Assignable) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolValue bool) AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	WithHash(hash hash.Hash) AssignableBuilder
	WithHashList(hashList []hash.Hash) AssignableBuilder
	WithStringList(strList []string) AssignableBuilder
	WithUnsignedInt(unsignedInt uint) AssignableBuilder
	WithAccount(account stack_accounts.Account) AssignableBuilder
	WithInstance(ins instances.Instance) AssignableBuilder
	WithSkeleton(skeleton skeletons.Skeleton) AssignableBuilder
	WithQuery(query queries.Query) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBool() bool
	Bool() *bool
	IsBytes() bool
	Bytes() []byte
	IsHash() bool
	Hash() hash.Hash
}
