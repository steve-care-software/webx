package combiners

import (
	"github.com/steve-care-software/syntax/domain/syntax/trees"
)

// Application represents the combine application
type Application interface {
	Execute(trees []trees.Tree, includeChannels bool) ([]byte, error)
}
