package grammars

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
)

// NewApplication creates a new application
func NewApplication() Application {
	grammarTokenBuilder := grammars.NewTokenBuilder()
	treesBuilder := trees.NewBuilder()
	treeBuilder := trees.NewTreeBuilder()
	treeBlockBuilder := trees.NewBlockBuilder()
	treeLineBuilder := trees.NewLineBuilder()
	treeElementsBuilder := trees.NewElementsBuilder()
	treeElementBuilder := trees.NewElementBuilder()
	treeContentBuilder := trees.NewContentBuilder()
	treeValueBuilder := trees.NewValueBuilder()
	return createApplication(
		grammarTokenBuilder,
		treesBuilder,
		treeBuilder,
		treeBlockBuilder,
		treeLineBuilder,
		treeElementsBuilder,
		treeElementBuilder,
		treeContentBuilder,
		treeValueBuilder,
	)
}

// Application represents a grammar application
type Application interface {
	Execute(grammar grammars.Grammar, values []byte) (trees.Tree, []byte, error)
}
