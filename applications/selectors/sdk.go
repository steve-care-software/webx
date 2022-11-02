package selectors

import (
	"github.com/steve-care-software/webx/domain/selectors"
	"github.com/steve-care-software/webx/domain/trees"
)

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents the selector application
type Application interface {
	Execute(selector selectors.Selector, tree trees.Tree) (interface{}, bool, error)
}
