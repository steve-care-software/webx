package stacks

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/signers"
)

// NewStackForTests creates a new stack for tests
func NewStackForTests(frames Frames) Stack {
	ins, err := NewBuilder().Create().WithFrames(frames).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFramesForTests creates a new frames for tests
func NewFramesForTests(frames []Frame) Frames {
	ins, err := NewFramesBuilder().Create().WithList(frames).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFrameForTests creates a new frame for tests
func NewFrameForTests() Frame {
	ins, err := NewFrameBuilder().Create().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFrameWithAssignmentsForTests creates a new frame with assignments for tests
func NewFrameWithAssignmentsForTests(value Assignments) Frame {
	ins, err := NewFrameBuilder().Create().WithAssignments(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignmentsForTests creates a new assignments for tests
func NewAssignmentsForTests(list []Assignment) Assignments {
	ins, err := NewAssignmentsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignmentForTests creates a new assignment for tests
func NewAssignmentForTests(name string, assignable Assignable) Assignment {
	ins, err := NewAssignmentBuilder().Create().WithName(name).WithAssignable(assignable).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignablesForTests creates a new assignables for tests
func NewAssignablesForTests(value []Assignable) Assignables {
	ins, err := NewAssignablesBuilder().Create().WithList(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithBoolForTests creates a new assignable with bool for tests
func NewAssignableWithBoolForTests(value bool) Assignable {
	ins, err := NewAssignableBuilder().Create().WithBool(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithBytesForTests creates a new assignable with bytes for tests
func NewAssignableWithBytesForTests(value []byte) Assignable {
	ins, err := NewAssignableBuilder().Create().WithBytes(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithStringForTests creates a new assignable with string for tests
func NewAssignableWithStringForTests(value string) Assignable {
	ins, err := NewAssignableBuilder().Create().WithString(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithHashForTests creates a new assignable with hash for tests
func NewAssignableWithHashForTests(value hash.Hash) Assignable {
	ins, err := NewAssignableBuilder().Create().WithHash(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithUnsignedIntForTests creates a new assignable with unsigned int for tests
func NewAssignableWithUnsignedIntForTests(value uint) Assignable {
	ins, err := NewAssignableBuilder().Create().WithUnsignedInt(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithInstanceForTests creates a new assignable with instance for tests
func NewAssignableWithInstanceForTests(value instances.Instance) Assignable {
	ins, err := NewAssignableBuilder().Create().WithInstance(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithEncryptorForTests creates a new assignable with encryptor for tests
func NewAssignableWithEncryptorForTests(value encryptors.Encryptor) Assignable {
	ins, err := NewAssignableBuilder().Create().WithEncryptor(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithPublicKeyForTests creates a new assignable with encryptor public key for tests
func NewAssignableWithEncryptorPublicKeyForTests(value encryptors.PublicKey) Assignable {
	ins, err := NewAssignableBuilder().Create().WithEncryptorPubKey(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithSignerForTests creates a new assignable with signer for tests
func NewAssignableWithSignerForTests(value signers.Signer) Assignable {
	ins, err := NewAssignableBuilder().Create().WithSigner(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithSignerPublicKeyForTests creates a new assignable with signer public key for tests
func NewAssignableWithSignerPublicKeyForTests(value signers.PublicKey) Assignable {
	ins, err := NewAssignableBuilder().Create().WithSignerPubKey(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithSignatureForTests creates a new assignable with signature for tests
func NewAssignableWithSignatureForTests(value signers.Signature) Assignable {
	ins, err := NewAssignableBuilder().Create().WithSignature(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithListForTests creates a new assignable with list for tests
func NewAssignableWithListForTests(value Assignables) Assignable {
	ins, err := NewAssignableBuilder().Create().WithList(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithVoteForTests creates a new assignable with vote for tests
func NewAssignableWithVoteForTests(value signers.Vote) Assignable {
	ins, err := NewAssignableBuilder().Create().WithVote(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
