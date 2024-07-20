package signs

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// NewSignWithCreateForTests creates a new sign with create for tests
func NewSignWithCreateForTests(create creates.Create) Sign {
	ins, err := NewBuilder().Create().WithCreate(create).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignWithValidateForTests creates a new sign with validate for tests
func NewSignWithValidateForTests(validate validates.Validate) Sign {
	ins, err := NewBuilder().Create().WithValidate(validate).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
