package hashtrees

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// MinHashtreeSize represents the minimum hashtree size
const MinHashtreeSize = hash.Size + minParentLeaf

const minLeafSize = hash.Size
const minParentLeaf = minLeafSize * 2
const minLeavesSize = minLeafSize + 8 + 8
const minCompactSize = hash.Size + minLeafSize

var leafAdapterIns LeafAdapter

func init() {
	hashAdapter := hash.NewAdapter()
	parentLeafBuilder := NewParentLeafBuilder()
	leafBuilder := NewLeafBuilder()
	leavesBuilder := NewLeavesBuilder()
	hashTreeBuilder := NewBuilder()
	leafAdapterIns = createLeafAdapter(
		hashAdapter,
		parentLeafBuilder,
		leafBuilder,
		leavesBuilder,
		hashTreeBuilder,
	)
}

// NewAdapter creates a newadapter instance
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	compactBuilder := NewCompactBuilder()
	builder := NewBuilder()
	return createAdapter(
		hashAdapter,
		leafAdapterIns,
		compactBuilder,
		builder,
	)
}

// NewBuilder creates a new hashtree builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	leafBuilder := NewLeafBuilder()
	leavesBuilder := NewLeavesBuilder()
	leavesAdapter := NewLeavesAdapter()
	return createBuilder(hashAdapter, leafBuilder, leavesBuilder, leavesAdapter)
}

// NewParentLeafBuilder creates a new parent leaf builder
func NewParentLeafBuilder() ParentLeafBuilder {
	return createParentLeafBuilder()
}

// NewLeafAdapter creates a new leaf adapter
func NewLeafAdapter() LeafAdapter {
	hashAdapter := hash.NewAdapter()
	parentLeafBuilder := NewParentLeafBuilder()
	leafBuilder := NewLeafBuilder()
	leavesBuilder := NewLeavesBuilder()
	hashTreeBuilder := NewBuilder()
	return createLeafAdapter(
		hashAdapter,
		parentLeafBuilder,
		leafBuilder,
		leavesBuilder,
		hashTreeBuilder,
	)
}

// NewLeafBuilder creates a new leaf builder
func NewLeafBuilder() LeafBuilder {
	return createLeafBuilder()
}

// NewLeavesAdapter creates a new leaves adapter
func NewLeavesAdapter() LeavesAdapter {
	hashAdapter := hash.NewAdapter()
	parentLeafBuilder := NewParentLeafBuilder()
	leavesBuilder := NewLeavesBuilder()
	leafBuilder := NewLeafBuilder()
	return createLeavesAdapter(
		hashAdapter,
		parentLeafBuilder,
		leafAdapterIns,
		leavesBuilder,
		leafBuilder,
	)
}

// NewLeavesBuilder creates a new leaves builder
func NewLeavesBuilder() LeavesBuilder {
	return createLeavesBuilder()
}

// NewCompactAdapter creates a new compact adapter
func NewCompactAdapter() CompactAdapter {
	hashAdapter := hash.NewAdapter()
	leavesAdapter := NewLeavesAdapter()
	builder := NewCompactBuilder()
	return createCompactAdapter(hashAdapter, leavesAdapter, builder)
}

// NewCompactBuilder creates a new compact builder
func NewCompactBuilder() CompactBuilder {
	return createCompactBuilder()
}

// Adapter represents the adapter
type Adapter interface {
	ToContent(ins HashTree) ([]byte, error)
	ToHashTree(content []byte) (HashTree, error)
	ToLength(ins HashTree) (*uint, error)
	ToCompact(ins HashTree) (Compact, error)
	ToOrder(ins HashTree, mixed [][]byte) ([][]byte, error)
}

// Builder represents an hashtree builder
type Builder interface {
	Create() Builder
	WithBlocks(blocks [][]byte) Builder
	WithHead(head hash.Hash) Builder
	WithParent(parent ParentLeaf) Builder
	Now() (HashTree, error)
}

// HashTree represents an hashtree
type HashTree interface {
	Height() uint
	Head() hash.Hash
	Parent() ParentLeaf
}

// ParentLeafBuilder represents a parent leaf builder
type ParentLeafBuilder interface {
	Create() ParentLeafBuilder
	WithLeft(left Leaf) ParentLeafBuilder
	WithRight(right Leaf) ParentLeafBuilder
	Now() (ParentLeaf, error)
}

// ParentLeaf represents an hashtree parent leaf
type ParentLeaf interface {
	Left() Leaf
	Right() Leaf
}

// LeafAdapter represents the leaf adapter
type LeafAdapter interface {
	LeafToContent(ins Leaf) ([]byte, error)
	LeafToLeaves(ins Leaf) (Leaves, error)
	ContentToLeaf(content []byte) (Leaf, error)
	ParentLeafToContent(ins ParentLeaf) ([]byte, error)
	ParentLeafToLeaves(ins ParentLeaf) (Leaves, error)
	ParentLeafToHashTree(ins ParentLeaf) (HashTree, error)
	ContentToParentLeaf(content []byte) (ParentLeaf, error)
}

// LeafBuilder creates a new leaf builder
type LeafBuilder interface {
	Create() LeafBuilder
	WithHead(head hash.Hash) LeafBuilder
	WithParent(parent ParentLeaf) LeafBuilder
	Now() (Leaf, error)
}

// Leaf represents an hashtree leaf
type Leaf interface {
	Head() hash.Hash
	Height() uint
	HasParent() bool
	Parent() ParentLeaf
}

// LeavesAdapter represents the leaves adapter
type LeavesAdapter interface {
	ToContent(ins Leaves) ([]byte, error)
	ToLeaves(content []byte) (Leaves, error)
	ToHashTree(ins Leaves) (HashTree, error)
}

// LeavesBuilder represents a leaves builder
type LeavesBuilder interface {
	Create() LeavesBuilder
	WithList(list []Leaf) LeavesBuilder
	Now() (Leaves, error)
}

// Leaves represents a list of Leaf instances
type Leaves interface {
	Leaves() []Leaf
	Merge(lves Leaves) Leaves
}

// CompactAdapter represents the compact adapter
type CompactAdapter interface {
	ToContent(ins Compact) ([]byte, error)
	ToCompact(content []byte) (Compact, error)
}

// CompactBuilder represents a compact builder
type CompactBuilder interface {
	Create() CompactBuilder
	WithHead(head hash.Hash) CompactBuilder
	WithLeaves(leaves Leaves) CompactBuilder
	Now() (Compact, error)
}

// Compact represents a compact hashtree
type Compact interface {
	Head() hash.Hash
	Leaves() Leaves
	Length() uint
}
