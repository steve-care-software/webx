package hashtrees

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type leavesAdapter struct {
	hashAdapter       hash.Adapter
	parentLeafBuilder ParentLeafBuilder
	leafAdapter       LeafAdapter
	leavesBuilder     LeavesBuilder
	leafBuilder       LeafBuilder
}

func createLeavesAdapter(
	hashAdapter hash.Adapter,
	parentLeafBuilder ParentLeafBuilder,
	leafAdapter LeafAdapter,
	leavesBuilder LeavesBuilder,
	leafBuilder LeafBuilder,
) LeavesAdapter {
	out := leavesAdapter{
		hashAdapter:       hashAdapter,
		parentLeafBuilder: parentLeafBuilder,
		leafAdapter:       leafAdapter,
		leavesBuilder:     leavesBuilder,
		leafBuilder:       leafBuilder,
	}

	return &out
}

// ToContent converts leaves to bytes
func (app *leavesAdapter) ToContent(ins Leaves) ([]byte, error) {
	list := ins.Leaves()
	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(len(list)))

	output := []byte{}
	output = append(output, lengthBytes...)

	for _, oneLeaf := range list {
		content, err := app.leafAdapter.LeafToContent(oneLeaf)
		if err != nil {
			return nil, err
		}

		leafLengthBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(leafLengthBytes, uint64(len(content)))

		output = append(output, leafLengthBytes...)
		output = append(output, content...)
	}

	return output, nil
}

// ToLeaves converts bytes to leaves
func (app *leavesAdapter) ToLeaves(content []byte) (Leaves, error) {
	contentLength := len(content)
	if contentLength < minLeavesSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Leaves instance, %d provided", minLeavesSize, contentLength)
		return nil, errors.New(str)
	}

	list := []Leaf{}
	lastEndsOn := 8
	length := int(binary.LittleEndian.Uint64(content[:lastEndsOn]))
	for i := 0; i < length; i++ {
		lengthBeginsOn := lastEndsOn
		lengthEndsOn := lengthBeginsOn + 8
		leafLength := int(binary.LittleEndian.Uint64(content[lengthBeginsOn:lengthEndsOn]))

		endsOn := lengthEndsOn + leafLength
		ins, err := app.leafAdapter.ContentToLeaf(content[lengthEndsOn:endsOn])
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
		lastEndsOn = endsOn
	}

	return app.leavesBuilder.Create().WithList(list).Now()
}

// ToHashTree converts a leaves to hashtree instance
func (app *leavesAdapter) ToHashTree(ins Leaves) (HashTree, error) {
	list := ins.Leaves()
	length := len(list)
	if length == 2 {
		left := list[0]
		right := list[1]
		parent, err := app.parentLeafBuilder.Create().WithLeft(left).WithRight(right).Now()
		if err != nil {
			return nil, err
		}

		return app.leafAdapter.ParentLeafToHashTree(parent)
	}

	childrenLeaves, err := app.createChildrenLeaves(list)
	if err != nil {
		return nil, err
	}

	return app.ToHashTree(childrenLeaves)
}

func (app *leavesAdapter) createChildrenLeaves(list []Leaf) (Leaves, error) {
	childrenLeaves := []Leaf{}
	for index, oneLeaf := range list {

		if index%2 != 0 {
			continue
		}

		left := oneLeaf
		right := list[index+1]
		child, err := app.createChildLeaf(left, right)
		if err != nil {
			return nil, err
		}

		parent := createParentLeaf(left, right)
		childWithParent := createLeafWithParent(child.Head(), parent)
		childrenLeaves = append(childrenLeaves, childWithParent)
	}

	return createLeaves(childrenLeaves), nil
}

func (app *leavesAdapter) createChildLeaf(left Leaf, right Leaf) (Leaf, error) {
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		left.Head().Bytes(),
		right.Head().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return app.leafBuilder.Create().
		WithHead(*pHash).
		Now()
}
