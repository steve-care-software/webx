package hashtrees

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type leafAdapter struct {
	hashAdapter       hash.Adapter
	parentLeafBuilder ParentLeafBuilder
	leafBuilder       LeafBuilder
	leavesBuilder     LeavesBuilder
	hashTreeBuilder   Builder
}

func createLeafAdapter(
	hashAdapter hash.Adapter,
	parentLeafBuilder ParentLeafBuilder,
	leafBuilder LeafBuilder,
	leavesBuilder LeavesBuilder,
	hashTreeBuilder Builder,
) LeafAdapter {
	out := leafAdapter{
		hashAdapter:       hashAdapter,
		parentLeafBuilder: parentLeafBuilder,
		leafBuilder:       leafBuilder,
		leavesBuilder:     leavesBuilder,
		hashTreeBuilder:   hashTreeBuilder,
	}

	return &out
}

// LeafToContent converts leaf to bytes
func (app *leafAdapter) LeafToContent(ins Leaf) ([]byte, error) {
	headBytes := ins.Head().Bytes()

	output := []byte{}
	output = append(output, headBytes...)

	if ins.HasParent() {
		parentBytes, err := app.ParentLeafToContent(ins.Parent())
		if err != nil {
			return nil, err
		}

		output = append(output, parentBytes...)
	}

	return output, nil
}

// LeafToLeaves converts a lead to leaves
func (app *leafAdapter) LeafToLeaves(ins Leaf) (Leaves, error) {
	if ins.HasParent() {
		return app.ParentLeafToLeaves(ins.Parent())
	}

	return app.leavesBuilder.Create().WithList([]Leaf{
		ins,
	}).Now()
}

// ContentToLeaf converts bytes to leaf
func (app *leafAdapter) ContentToLeaf(content []byte) (Leaf, error) {
	contentLength := len(content)
	if contentLength < minLeafSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Leaf instance, %d provided", minLeafSize, contentLength)
		return nil, errors.New(str)
	}

	pHead, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	remaining := content[hash.Size:]
	builder := app.leafBuilder.Create().WithHead(*pHead)
	if len(remaining) > 0 {
		parentLeaf, err := app.ContentToParentLeaf(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithParent(parentLeaf)
	}

	return builder.Now()
}

// ParentLeafToContent converts parent leaf to bytes
func (app *leafAdapter) ParentLeafToContent(ins ParentLeaf) ([]byte, error) {
	leftBytes, err := app.LeafToContent(ins.Left())
	if err != nil {
		return nil, err
	}

	rightBytes, err := app.LeafToContent(ins.Right())
	if err != nil {
		return nil, err
	}

	leftLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(leftLengthBytes, uint64(len(leftBytes)))

	output := []byte{}
	output = append(output, leftLengthBytes...)
	output = append(output, leftBytes...)
	output = append(output, rightBytes...)

	return output, nil
}

// ParentLeafToBlockLeaves converts parentLeaf to leaves
func (app *leafAdapter) ParentLeafToLeaves(ins ParentLeaf) (Leaves, error) {
	leftLeaves, err := app.LeafToLeaves(ins.Left())
	if err != nil {
		return nil, err
	}

	rightLeaves, err := app.LeafToLeaves(ins.Right())
	if err != nil {
		return nil, err
	}

	return leftLeaves.Merge(rightLeaves), nil
}

// ParentLeafToHashTree converts parentLeaf to hashTree
func (app *leafAdapter) ParentLeafToHashTree(ins ParentLeaf) (HashTree, error) {
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		ins.Left().Head().Bytes(),
		ins.Right().Head().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return app.hashTreeBuilder.Create().
		WithHead(*pHash).
		WithParent(ins).
		Now()
}

// ContentToLeaf converts bytes to parent leaf
func (app *leafAdapter) ContentToParentLeaf(content []byte) (ParentLeaf, error) {
	contentLength := len(content)
	if contentLength < minParentLeaf {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a ParentLeaf instance, %d provided", minParentLeaf, contentLength)
		return nil, errors.New(str)
	}

	leftLength := int(binary.LittleEndian.Uint64(content[:8]))
	left, err := app.ContentToLeaf(content[8 : 8+leftLength])
	if err != nil {
		return nil, err
	}

	right, err := app.ContentToLeaf(content[leftLength+8:])
	if err != nil {
		return nil, err
	}

	return app.parentLeafBuilder.Create().
		WithLeft(left).
		WithRight(right).
		Now()
}
