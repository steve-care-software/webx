package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/updates"
)

type switcher struct {
	original singles.Single
	updated  updates.Update
}

func createSwitcherWithOriginalAndUpdated(
	original singles.Single,
	updated updates.Update,
) Switcher {
	return createSwitcherInternally(original, updated)
}

func createSwitcherWithOriginal(
	original singles.Single,
) Switcher {
	return createSwitcherInternally(original, nil)
}

func createSwitcherWithUpdated(
	updated updates.Update,
) Switcher {
	return createSwitcherInternally(nil, updated)
}

func createSwitcherInternally(
	original singles.Single,
	updated updates.Update,
) Switcher {
	out := switcher{
		original: original,
		updated:  updated,
	}

	return &out
}

// Current returns the current instance
func (obj *switcher) Current() singles.Single {
	if obj.HasUpdated() {
		return obj.updated.Single()
	}

	return obj.original
}

// HasOriginal returns true if there is an original, false otherwise
func (obj *switcher) HasOriginal() bool {
	return obj.original != nil
}

// Original returns the original, if any
func (obj *switcher) Original() singles.Single {
	return obj.original
}

// HasUpdated returns true if there is an updated, false otherwise
func (obj *switcher) HasUpdated() bool {
	return obj.updated != nil
}

// Updated returns the updated, if any
func (obj *switcher) Updated() updates.Update {
	return obj.updated
}