package resources

import (
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
)

// NewResourceWithValueForTests creates a new resource with value for tests
func NewResourceWithValueForTests(value interface{}) Resource {
	ins, err := NewBuilder().Create().WithValue(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceWithFieldForTests creates a new resource with field for tests
func NewResourceWithFieldForTests(field pointers.Pointer) Resource {
	ins, err := NewBuilder().Create().WithField(field).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
