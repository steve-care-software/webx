package selectors

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// Application represents a selector application
type Application interface {
	Matches(grammar grammars.Grammar, selector selectors.Selector) (bool, error)
	Execute(selector selectors.Selector, script []byte) (interface{}, bool, []byte, error)
}
