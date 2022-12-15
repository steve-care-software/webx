package contents

import "github.com/steve-care-software/webx/domain/databases/entities"

type value struct {
	pByte *byte
	tree  entities.Identifier
}

func createValueWithByte(
	pByte *byte,
) Value {
	return createValueInternally(pByte, nil)
}

func createValueWithTree(
	tree entities.Identifier,
) Value {
	return createValueInternally(nil, tree)
}

func createValueInternally(
	pByte *byte,
	tree entities.Identifier,
) Value {
	out := value{
		pByte: nil,
		tree:  nil,
	}

	return &out
}

// IsByte returns true if there is a byte, false otherwise
func (obj *value) IsByte() bool {
	return obj.pByte != nil
}

// Byte returns the byte, if any
func (obj *value) Byte() *byte {
	return obj.pByte
}

// IsTree returns true if there is a tree, false otherwise
func (obj *value) IsTree() bool {
	return obj.tree != nil
}

// Tree returns the tree, if any
func (obj *value) Tree() entities.Identifier {
	return obj.tree
}
