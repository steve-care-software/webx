package hashtrees

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type adapter struct {
	hashAdapter    hash.Adapter
	leafAdapter    LeafAdapter
	compactBuilder CompactBuilder
	builder        Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	leafAdapter LeafAdapter,
	compactBuilder CompactBuilder,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:    hashAdapter,
		leafAdapter:    leafAdapter,
		compactBuilder: compactBuilder,
		builder:        builder,
	}

	return &out
}

// ToContent converts hashTree to bytes
func (app *adapter) ToContent(ins HashTree) ([]byte, error) {
	headBytes := ins.Head().Bytes()
	parentBytes, err := app.leafAdapter.ParentLeafToContent(ins.Parent())
	if err != nil {
		return nil, err
	}

	output := []byte{}
	output = append(output, headBytes...)
	output = append(output, parentBytes...)

	return output, nil
}

// ToHashTree converts bytes to hashtree
func (app *adapter) ToHashTree(content []byte) (HashTree, error) {
	contentLength := len(content)
	if contentLength < MinHashtreeSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to an HashTree instance, %d provided", MinHashtreeSize, contentLength)
		return nil, errors.New(str)
	}

	pHead, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	parentLeaf, err := app.leafAdapter.ContentToParentLeaf(content[hash.Size:])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithHead(*pHead).
		WithParent(parentLeaf).
		Now()
}

// ToLength converts an hashtree to its length
func (app *adapter) ToLength(ins HashTree) (*uint, error) {
	leaves, err := app.leafAdapter.ParentLeafToLeaves(ins.Parent())
	if err != nil {
		return nil, err
	}

	length := uint(len(leaves.Leaves()))
	return &length, nil
}

// ToCompact converts an hashtree to a compact instance
func (app *adapter) ToCompact(ins HashTree) (Compact, error) {
	leaves, err := app.leafAdapter.ParentLeafToLeaves(ins.Parent())
	if err != nil {
		return nil, err
	}

	head := ins.Head()
	return app.compactBuilder.Create().
		WithHead(head).
		WithLeaves(leaves).
		Now()
}

// ToOrder takes an hashtree and mixed data and re-order the data according to the hashTree
func (app *adapter) ToOrder(ins HashTree, mixed [][]byte) ([][]byte, error) {
	hashed := map[string][]byte{}
	for _, oneData := range mixed {
		hsh, err := app.hashAdapter.FromBytes(oneData)
		if err != nil {
			return nil, err
		}

		hashAsString := hsh.String()
		hashed[hashAsString] = oneData
	}

	out := [][]byte{}
	leaves, err := app.leafAdapter.ParentLeafToLeaves(ins.Parent())
	if err != nil {
		return nil, err
	}

	list := leaves.Leaves()
	for _, oneLeaf := range list {
		leafHashAsString := oneLeaf.Head().String()
		if oneData, ok := hashed[leafHashAsString]; ok {
			out = append(out, oneData)
			continue
		}

		//must be a filling Leaf, so continue:
		continue
	}

	if len(out) != len(mixed) {
		str := fmt.Sprintf("the length of the input data (%d) does not match the length of the output (%d), therefore, some data blocks could not be found in the hash leaves", len(mixed), len(out))
		return nil, errors.New(str)
	}

	return out, nil
}
