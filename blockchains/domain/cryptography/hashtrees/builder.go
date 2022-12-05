package hashtrees

import (
	"errors"
	"math"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	leafBuilder   LeafBuilder
	leavesBuilder LeavesBuilder
	leavesAdapter LeavesAdapter
	blocks        [][]byte
	pHead         *hash.Hash
	parent        ParentLeaf
}

func createBuilder(
	hashAdapter hash.Adapter,
	leafBuilder LeafBuilder,
	leavesBuilder LeavesBuilder,
	leavesAdapter LeavesAdapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		leafBuilder:   leafBuilder,
		leavesBuilder: leavesBuilder,
		leavesAdapter: leavesAdapter,
		blocks:        nil,
		pHead:         nil,
		parent:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.leafBuilder,
		app.leavesBuilder,
		app.leavesAdapter,
	)
}

// WithBlocks add blocks to the builder
func (app *builder) WithBlocks(blocks [][]byte) Builder {
	app.blocks = blocks
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head hash.Hash) Builder {
	app.pHead = &head
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent ParentLeaf) Builder {
	app.parent = parent
	return app
}

// Now builds a new HashTree instance
func (app *builder) Now() (HashTree, error) {
	if app.pHead != nil && app.parent != nil {
		return createHashTree(*app.pHead, app.parent), nil
	}

	if app.blocks == nil {
		return nil, errors.New("the blocks are mandatory in order to build an HashTree instance")
	}

	if len(app.blocks) <= 1 {
		app.blocks = append(app.blocks, []byte(""))
	}

	hashes := []hash.Hash{}
	for _, oneData := range app.blocks {
		oneHash, err := app.hashAdapter.FromBytes(oneData)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *oneHash)
	}

	//need to make sure the elements are always a power of 2:
	length := len(hashes)
	isPowerOfTwo := (length != 0) && ((length & (length - 1)) == 0)
	if !isPowerOfTwo {
		lengthAsFloat := float64(len(hashes))
		next := uint(math.Pow(2, math.Ceil(math.Log(lengthAsFloat)/math.Log(2))))
		remaining := int(next) - int(lengthAsFloat)
		for i := 0; i < remaining; i++ {
			single, err := app.hashAdapter.FromBytes(nil)
			if err != nil {
				return nil, err
			}

			hashes = append(hashes, *single)
		}
	}

	list := []Leaf{}
	for _, oneBlockHash := range hashes {
		leaf, err := app.leafBuilder.Create().WithHead(oneBlockHash).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, leaf)
	}

	leaves, err := app.leavesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, err
	}

	return app.leavesAdapter.ToHashTree(leaves)
}
