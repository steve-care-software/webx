package coverages

import "github.com/steve-care-software/webx/domain/trees"

type result struct {
	tree trees.Tree
	err  string
}

func createResultWithTree(
	tree trees.Tree,
) Result {
	return createResultInternally(tree, "")
}

func createResultWithError(
	err string,
) Result {
	return createResultInternally(nil, err)
}

func createResultInternally(
	tree trees.Tree,
	err string,
) Result {
	out := result{
		tree: tree,
		err:  err,
	}

	return &out
}

// IsTree returns true if there is a tree, false otherwise
func (obj *result) IsTree() bool {
	return obj.tree != nil
}

// Tree returns the tree, if any
func (obj *result) Tree() trees.Tree {
	return obj.tree
}

// IsError returns true if there is an error, false otherwise
func (obj *result) IsError() bool {
	return obj.err != ""
}

// Error returns the error, if any
func (obj *result) Error() string {
	return obj.err
}
