package stacks

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries"
	"github.com/steve-care-software/datastencil/domain/skeletons"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

// NewFrameForTests creates a new frame for tests
func NewFrameForTests(value Assignments) Frame {
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

// NewAssignableWithHashForTests creates a new assignable with hash for tests
func NewAssignableWithHashForTests(value hash.Hash) Assignable {
	ins, err := NewAssignableBuilder().Create().WithHash(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithHashListForTests creates a new assignable with hash list for tests
func NewAssignableWithHashListForTests(value []hash.Hash) Assignable {
	ins, err := NewAssignableBuilder().Create().WithHashList(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithStringListForTests creates a new assignable with string list for tests
func NewAssignableWithStringListForTests(value []string) Assignable {
	ins, err := NewAssignableBuilder().Create().WithStringList(value).Now()
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

// NewAssignableWithAccountForTests creates a new assignable with account for tests
func NewAssignableWithAccountForTests(value stack_accounts.Account) Assignable {
	ins, err := NewAssignableBuilder().Create().WithAccount(value).Now()
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

// NewAssignableWithSkeletonForTests creates a new assignable with skeleton for tests
func NewAssignableWithSkeletonForTests(value skeletons.Skeleton) Assignable {
	ins, err := NewAssignableBuilder().Create().WithSkeleton(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithQueryForTests creates a new assignable with query for tests
func NewAssignableWithQueryForTests(value queries.Query) Assignable {
	ins, err := NewAssignableBuilder().Create().WithQuery(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
