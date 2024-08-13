package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

// NewSwitchersForTests creates a new switchers for tests
func NewSwitchersForTests(list []Switcher) Switchers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithOriginalAndInsertForTests creates a new switcher with original and insert for tests
func NewSwitcherWithOriginalAndInsertForTests(original singles.Single, insert inserts.Insert) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithOriginal(original).WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithOriginalAndUpdateForTests creates a new switcher with original and update for tests
func NewSwitcherWithOriginalAndUpdateForTests(original singles.Single, update updates.Update) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithOriginal(original).WithUpdate(update).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithOriginalAndDeleteForTests creates a new switcher with original and delete for tests
func NewSwitcherWithOriginalAndDeleteForTests(original singles.Single, delete deletes.Delete) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithOriginal(original).WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithInsertForTests creates a new switcher with insert for tests
func NewSwitcherWithInsertForTests(insert inserts.Insert) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithUpdateForTests creates a new switcher with update for tests
func NewSwitcherWithUpdateForTests(update updates.Update) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithUpdate(update).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithDeleteForTests creates a new switcher with delete for tests
func NewSwitcherWithDeleteForTests(delete deletes.Delete) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSwitcherWithOriginalForTests creates a new switcher with original for tests
func NewSwitcherWithOriginalForTests(original singles.Single) Switcher {
	ins, err := NewSwitcherBuilder().Create().WithOriginal(original).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
