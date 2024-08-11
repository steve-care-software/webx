package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/updates"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSwitcherBuilder creates a new switcher builder
func NewSwitcherBuilder() SwitcherBuilder {
	return createSwitcherBuilder()
}

// Builder represents a switchers builder
type Builder interface {
	Create() Builder
	WithList(list []Switcher) Builder
	Now() (Switchers, error)
}

// Switchers represents switchers
type Switchers interface {
	List() []Switcher
}

// SwitcherBuilder represents a switcher builder
type SwitcherBuilder interface {
	Create() SwitcherBuilder
	WithOriginal(original singles.Single) SwitcherBuilder
	WithUpdated(updated updates.Update) SwitcherBuilder
	Now() (Switcher, error)
}

// Switcher represents an identity switcher
type Switcher interface {
	Current() singles.Single
	HasOriginal() bool
	Original() singles.Single
	HasUpdated() bool
	Updated() updates.Update
}
