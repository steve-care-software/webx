package resources

import "github.com/steve-care-software/datastencil/domain/instances"

// NewResourceForTests creates a new resource for tests
func NewResourceForTests(path []string, ins instances.Instance) Resource {
	retIns, err := NewBuilder().Create().
		WithPath(path).
		WithInstance(ins).
		Now()

	if err != nil {
		panic(err)
	}

	return retIns
}
