package votes

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// NewVoteWithCreateForTests creates a new vote with create for tests
func NewVoteWithCreateForTests(create creates.Create) Vote {
	ins, err := NewBuilder().Create().WithCreate(create).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewVoteWithValidateForTests creates a new vote with validate for tests
func NewVoteWithValidateForTests(validate validates.Validate) Vote {
	ins, err := NewBuilder().Create().WithValidate(validate).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
