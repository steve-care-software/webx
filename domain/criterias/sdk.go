package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewNodeBuilder creates a new node builder
func NewNodeBuilder() NodeBuilder {
	return createNodeBuilder()
}

// NewTailBuilder creates a new tail builder
func NewTailBuilder() TailBuilder {
	hashAdapter := hash.NewAdapter()
	return createTailBuilder(hashAdapter)
}

// NewDelimiterBuilder creates a new delimiter builder
func NewDelimiterBuilder() DelimiterBuilder {
	hashAdapter := hash.NewAdapter()
	return createDelimiterBuilder(hashAdapter)
}

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithCurrent(current Tail) Builder
	WithNext(next Node) Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Hash() hash.Hash
	Current() Tail
	HasNext() bool
	Next() Node
}

// NodeBuilder represents the node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithNext(next Criteria) NodeBuilder
	WithTail(tail Tail) NodeBuilder
	Now() (Node, error)
}

// Node represents the criteria node
type Node interface {
	Hash() hash.Hash
	IsNext() bool
	Next() Criteria
	IsTail() bool
	Tail() Tail
}

// TailBuilder represents a tail builder
type TailBuilder interface {
	Create() TailBuilder
	WithName(name string) TailBuilder
	WithDelimiter(delimiter Delimiter) TailBuilder
	Now() (Tail, error)
}

// Tail represents a criteria tail
type Tail interface {
	Hash() hash.Hash
	Name() string
	HasDelimiter() bool
	Delimiter() Delimiter
}

// DelimiterBuilder represents a delimiter builder
type DelimiterBuilder interface {
	Create() DelimiterBuilder
	WithIndex(index uint) DelimiterBuilder
	WithAmount(amount uint) DelimiterBuilder
	Now() (Delimiter, error)
}

// Delimiter represents a delimiter
type Delimiter interface {
	Hash() hash.Hash
	Index() uint
	HasAmount() bool
	Amount() *uint
}
