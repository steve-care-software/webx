package gui

import (
	"github.com/steve-care-software/webx/gui/grids"
	"github.com/steve-care-software/webx/gui/menus"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	gridBuilder := grids.NewBuilder()
	menuBuilder := menus.NewBuilder()
	return createBuilder(
		gridBuilder,
		menuBuilder,
	)
}

// Builder represents the gui builder
type Builder interface {
	Create() Builder
	WithTitle(title string) Builder
	WithWidth(width float32) Builder
	WithHeight(height float32) Builder
	Now() (Gui, error)
}

// Gui represents the gui
type Gui interface {
	Execute() error
}
