package selectors

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/trees"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	return createApplication()
}

// Application represents a selector application
type Application interface {
	Matches(grammar grammars.Grammar, selector selectors.Selector) (bool, error)
	Execute(selector selectors.Selector, treeIns trees.Tree) (interface{}, bool, []byte, error)
}
