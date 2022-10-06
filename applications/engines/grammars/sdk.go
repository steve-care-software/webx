package grammars

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/coverages"
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
	treeContentsBuilder := trees.NewContentsBuilder()
	treeContentBuilder := trees.NewContentBuilder()
	treeValueBuilder := trees.NewValueBuilder()
	coveragesBuilder := coverages.NewBuilder()
	coverageBuilder := coverages.NewCoverageBuilder()
	coverageExecutionsBuilder := coverages.NewExecutionsBuilder()
	coverageExecutionBuilder := coverages.NewExecutionBuilder()
	coverageResultBuilder := coverages.NewResultBuilder()
	return createApplication(
		grammarTokenBuilder,
		treesBuilder,
		treeBuilder,
		treeBlockBuilder,
		treeLineBuilder,
		treeElementsBuilder,
		treeElementBuilder,
		treeContentsBuilder,
		treeContentBuilder,
		treeValueBuilder,
		coveragesBuilder,
		coverageBuilder,
		coverageExecutionsBuilder,
		coverageExecutionBuilder,
		coverageResultBuilder,
	)
}

// Application represents a grammar application
type Application interface {
	Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error)
	Coverages(grammar grammars.Grammar) (coverages.Coverages, error)
	Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error)
	Uncovered(grammar grammars.Grammar) (map[string]map[uint]map[uint]string, error)
}
