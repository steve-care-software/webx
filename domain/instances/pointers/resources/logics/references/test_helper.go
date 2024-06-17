package references

import "github.com/steve-care-software/datastencil/domain/instances"

// NewReferences creates new references
func NewReferences(list []Reference) References {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReference creates a new reference
func NewReference(variable string, instance instances.Instance) Reference {
	ins, err := NewReferenceBuilder().Create().WithVariable(variable).WithInstance(instance).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
