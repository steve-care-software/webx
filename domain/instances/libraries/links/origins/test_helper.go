package origins

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/resources"
)

// NewOriginForTests creates a new origin for tests
func NewOriginForTests(resource resources.Resource, operator operators.Operator, next Value) Origin {
	ins, err := NewBuilder().Create().WithResource(resource).WithOperator(operator).WithNext(next).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewValueWithOriginForTests creates a new origin value with origin for tests
func NewValueWithOriginForTests(origin Origin) Value {
	ins, err := NewValueBuilder().Create().WithOrigin(origin).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewValueWithResourceForTests creates a new origin value with resource for tests
func NewValueWithResourceForTests(resource resources.Resource) Value {
	ins, err := NewValueBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
