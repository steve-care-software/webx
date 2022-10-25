package selections

import (
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/trees"
	"github.com/steve-care-software/webx/domain/trees/selections"
)

// NewApplication creates a new application
func NewApplication() Application {
	builder := selections.NewBuilder()
	selectionBuilder := selections.NewSelectionBuilder()
	elementBuilder := selections.NewElementBuilder()
	childrenBuilder := selections.NewChildrenBuilder()
	childBuilder := selections.NewChildBuilder()
	return createApplication(
		builder,
		selectionBuilder,
		elementBuilder,
		childrenBuilder,
		childBuilder,
	)
}

// Application represents a selection application
type Application interface {
	Convert(tree trees.Tree, includeChannelBytes bool) (selections.Selections, error)
	Search(selections selections.Selections, criteria criterias.Criteria) (selections.Selections, error)
}
