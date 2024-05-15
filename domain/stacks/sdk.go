package stacks

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/keys/encryptors"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
)

// NewFactory creates a new factory
func NewFactory() Factory {
	return createFactory()
}

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
	FetchStringList(name string) ([]string, error)
	FetchRing(name string) ([]signers.PublicKey, error)
	FetchInstance(name string) (instances.Instance, error)
	FetchEncryptor(name string) (encryptors.Encryptor, error)
	FetchEncryptorPubKey(name string) (encryptors.PublicKey, error)
	FetchSigner(name string) (signers.Signer, error)
	FetchSignerPubKey(name string) (signers.PublicKey, error)
	FetchSignature(name string) (signers.Signature, error)
	FetchVote(name string) (signers.Vote, error)
	FetchList(name string) (Assignables, error)
	FetchModifications(name string) (modifications.Modifications, error)
	FetchString(name string) (string, error)
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

// AssignablesBuilder represents an assignables builder
type AssignablesBuilder interface {
	Create() AssignablesBuilder
	WithList(list []Assignable) AssignablesBuilder
	Now() (Assignables, error)
}

// Assignables represents assignables
type Assignables interface {
	List() []Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolValue bool) AssignableBuilder
	WithString(stringValue string) AssignableBuilder
	WithFloat(floatVal float64) AssignableBuilder
	WithInt(intVal int) AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	WithHash(hash hash.Hash) AssignableBuilder
	WithHashList(hashList []hash.Hash) AssignableBuilder
	WithStringList(strList []string) AssignableBuilder
	WithUnsignedInt(unsignedInt uint) AssignableBuilder
	WithInstance(ins instances.Instance) AssignableBuilder
	WithEncryptor(encryptor encryptors.Encryptor) AssignableBuilder
	WithEncryptorPubKey(encryptorPubKey encryptors.PublicKey) AssignableBuilder
	WithSigner(signer signers.Signer) AssignableBuilder
	WithSignerPubKey(signerPubKey signers.PublicKey) AssignableBuilder
	WithSignature(signature signers.Signature) AssignableBuilder
	WithVote(vote signers.Vote) AssignableBuilder
	WithList(list Assignables) AssignableBuilder
	WithAction(action actions.Action) AssignableBuilder
	WithCommit(commit commits.Commit) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBool() bool
	Bool() *bool
	IsString() bool
	String() *string
	IsBytes() bool
	Bytes() []byte
	IsHash() bool
	Hash() hash.Hash
	IsHashList() bool
	HashList() []hash.Hash
	IsStringList() bool
	StringList() []string
	IsUnsignedInt() bool
	UnsignedInt() *uint
	IsInstance() bool
	Instance() instances.Instance
	IsModification() bool
	Modification() modifications.Modification
	IsAction() bool
	Action() actions.Action
}
