package headers

import "github.com/steve-care-software/webx/engine/hashes/domain/pointers"

// NewHeaderForTests creates a new header for tests
func NewHeaderForTests() Header {
	ins, err := NewBuilder().Create().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewHeaderWithIdentitiesForTests creates a new header with identities for tests
func NewHeaderWithIdentitiesForTests(identities pointers.Pointer) Header {
	ins, err := NewBuilder().Create().WithIdentities(identities).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
