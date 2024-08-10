package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/profiles"
)

// NewSinglesForTests creates a new singles for tests
func NewSinglesForTests(list []Single) Singles {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSingleForTests creates a new single for tests
func NewSingleForTests(profile profiles.Profile, key keys.Key) Single {
	ins, err := NewSingleBuilder().Create().WithProfile(profile).WithKey(key).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
