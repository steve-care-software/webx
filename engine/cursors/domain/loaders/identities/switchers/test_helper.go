package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/updates"
)

// NewSwitchersForTests creates a new switchers for tests
func NewSwitchersForTests(list []Switcher) Switchers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherForTests creates a new switcher for tests
func NewSwitcherForTests(original singles.Single, updated updates.Update) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithOriginal(original).WithUpdated(updated).Now()
	if err != nil {
		panic(err)
	}

	return ins
}