package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

type switcher struct {
	original singles.Single
	insert   inserts.Insert
	update   updates.Update
	delete   deletes.Delete
}

func createSwitcherWithOriginal(
	original singles.Single,
) Switcher {
	return createSwitcherInternally(original, nil, nil, nil)
}

func createSwitcherWithInsert(
	insert inserts.Insert,
) Switcher {
	return createSwitcherInternally(nil, insert, nil, nil)
}

func createSwitcherWithUpdate(
	update updates.Update,
) Switcher {
	return createSwitcherInternally(nil, nil, update, nil)
}

func createSwitcherWithDelete(
	delete deletes.Delete,
) Switcher {
	return createSwitcherInternally(nil, nil, nil, delete)
}

func createSwitcherWithOriginalAndInsert(
	original singles.Single,
	insert inserts.Insert,
) Switcher {
	return createSwitcherInternally(original, insert, nil, nil)
}

func createSwitcherWithOriginalAndUpdate(
	original singles.Single,
	update updates.Update,
) Switcher {
	return createSwitcherInternally(original, nil, update, nil)
}

func createSwitcherWithOriginalAndDelete(
	original singles.Single,
	delete deletes.Delete,
) Switcher {
	return createSwitcherInternally(original, nil, nil, delete)
}

func createSwitcherInternally(
	original singles.Single,
	insert inserts.Insert,
	update updates.Update,
	delete deletes.Delete,
) Switcher {
	out := switcher{
		original: original,
		insert:   insert,
		update:   update,
		delete:   delete,
	}

	return &out
}

// Name returns the name
func (obj *switcher) Name() string {
	if obj.HasOriginal() {
		return obj.original.Storage().Name()
	}

	if obj.HasInsert() {
		return obj.insert.Name()
	}

	if obj.HasUpdate() {
		return obj.insert.Name()
	}

	return obj.delete.Name()
}

// HasOriginal returns true if there is an original, false otherwise
func (obj *switcher) HasOriginal() bool {
	return obj.original != nil
}

// Original returns the original, if any
func (obj *switcher) Original() singles.Single {
	return obj.original
}

// HasInsert returns true if there is an insert, false otherwise
func (obj *switcher) HasInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *switcher) Insert() inserts.Insert {
	return obj.insert
}

// HasUpdate returns true if there is an update, false otherwise
func (obj *switcher) HasUpdate() bool {
	return obj.update != nil
}

// Update returns the update, if any
func (obj *switcher) Update() updates.Update {
	return obj.update
}

// HasDelete returns true if there is a delete, false otherwise
func (obj *switcher) HasDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *switcher) Delete() deletes.Delete {
	return obj.delete
}
