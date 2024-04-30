package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewReferencesForTests creates references for tests
func NewReferencesForTests(list []Reference) References {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReferenceForTests creates a new reference for tests
func NewReferenceForTests(variable string, identifier hash.Hash) Reference {
	ins, err := NewReferenceBuilder().Create().
		WithVariable(variable).
		WithIdentifier(identifier).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
