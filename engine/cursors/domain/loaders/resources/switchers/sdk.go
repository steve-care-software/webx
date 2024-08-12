package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
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
	FetchByDelimiterIndex(delIndex uint64) (Switcher, error)
}

// SwitcherBuilder represents a switcher builder
type SwitcherBuilder interface {
	Create() SwitcherBuilder
	WithOriginal(original singles.Single) SwitcherBuilder
	WithUpdated(updated storages.Storage) SwitcherBuilder
	WithDeleted(deleted storages.Storage) SwitcherBuilder
	Now() (Switcher, error)
}

// Switcher represents a switcher switcher
type Switcher interface {
	Current() singles.Single
	HasOriginal() bool
	Original() singles.Single
	HasUpdated() bool
	Updated() storages.Storage
	HasDeleted() bool
	Deleted() storages.Storage
}
