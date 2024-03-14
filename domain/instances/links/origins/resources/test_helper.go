package resources

import "github.com/steve-care-software/datastencil/domain/hash"

// NewResourceWithIsMandatoryForTests creates a new origin resource with isMandatory for tests
func NewResourceWithIsMandatoryForTests(layer hash.Hash) Resource {
	ins, err := NewBuilder().Create().WithLayer(layer).IsMandatory().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceForTests creates a new origin resource for tests
func NewResourceForTests(layer hash.Hash) Resource {
	ins, err := NewBuilder().Create().WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
