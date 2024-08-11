package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/updates"
)

// Builder represents a switchers builder
type Builder interface {
	Create() Builder
	WithList(list []Switcher) Builder
	Now() (Switchers, error)
}

// Switchers represents switchers
type Switchers interface {
	List() []Switcher
	Names() []string
	Fetch(name string) (Switcher, error)
}

// SwitcherBuilder represents a switcher builder
type SwitcherBuilder interface {
	Create() SwitcherBuilder
	WithOriginal(original singles.Single) SwitcherBuilder
	WithUpdated(updated updates.Update) SwitcherBuilder
	Now() (Switcher, error)
}

// Switcher represents a version switcher
type Switcher interface {
	Current() singles.Single
	HasOriginal() bool
	Original() singles.Single
	HasUpdated() bool
	Updated() updates.Update
}
