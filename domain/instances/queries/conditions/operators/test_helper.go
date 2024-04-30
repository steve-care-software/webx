package operators

import (
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/integers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/relationals"
)

// NewOperatorWithIntegerForTests creates a new operator with integer for tests
func NewOperatorWithIntegerForTests(integer integers.Integer) Operator {
	ins, err := NewBuilder().Create().WithInteger(integer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithRelationalForTests creates a new operator with relational for tests
func NewOperatorWithRelationalForTests(relational relationals.Relational) Operator {
	ins, err := NewBuilder().Create().WithRelational(relational).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithEqualForTests creates a new operator with equal for tests
func NewOperatorWithEqualForTests() Operator {
	ins, err := NewBuilder().Create().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
