package trees

import "github.com/steve-care-software/syntax/domain/bytes/grammars/values"

type content struct {
	value values.Value
	tree  Tree
}

func createContentWithValue(
	value values.Value,
) Content {
	return createContentInternally(value, nil)
}

func createContentWithTree(
	tree Tree,
) Content {
	return createContentInternally(nil, tree)
}

func createContentInternally(
	value values.Value,
	tree Tree,
) Content {
	out := content{
		value: value,
		tree:  tree,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *content) IsValue() bool {
	return obj.value != nil
}

// Value returns the value if any
func (obj *content) Value() values.Value {
	return obj.value
}

// IsTree returns true if there is a tree, false otherwise
func (obj *content) IsTree() bool {
	return obj.tree != nil
}

// Tree returns the tree if any
func (obj *content) Tree() Tree {
	return obj.tree
}
