package switchers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
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
	FetchByName(name string) (Switcher, error)
}

// SwitcherBuilder represents a switcher builder
type SwitcherBuilder interface {
	Create() SwitcherBuilder
	WithOriginal(original singles.Single) SwitcherBuilder
	WithInsert(insert inserts.Insert) SwitcherBuilder
	WithUpdate(update updates.Update) SwitcherBuilder
	WithDelete(delete deletes.Delete) SwitcherBuilder
	Now() (Switcher, error)
}

// Switcher represents a switcher switcher
type Switcher interface {
	Name() string
	HasOriginal() bool
	Original() singles.Single
	HasInsert() bool
	Insert() inserts.Insert
	HasUpdate() bool
	Update() updates.Update
	HasDelete() bool
	Delete() deletes.Delete
}
