package values

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values/transforms"
)

// NewValueWithInstanceForTests creates a new value with instance for tests
func NewValueWithInstanceForTests(instance instances.Instance) Value {
	ins, err := NewBuilder().Create().WithInstance(instance).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewValueWithTransformForTests creates a new value with transform for tests
func NewValueWithTransformForTests(transform transforms.Transform) Value {
	ins, err := NewBuilder().Create().WithTransform(transform).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
