package contents

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

const minContentSize = hash.Size + 8 + 1
const minContentsSize = 8 + minContentSize

// NewAdapter creates a new adapter for tests
func NewAdapter() Adapter {
	builder := NewBuilder()
	contentAdapter := NewContentAdapter()
	return createAdapter(builder, contentAdapter)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewContentAdapter creates a new content adapter instance
func NewContentAdapter() ContentAdapter {
	hashAdapter := hash.NewAdapter()
	builder := NewContentBuilder()
	return createContentAdapter(hashAdapter, builder)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// Adapter represents a connections adapter
type Adapter interface {
	ToContent(ins Contents) ([]byte, error)
	ToInstance(content []byte) (Contents, error)
}

// Builder represents a contents builder
type Builder interface {
	Create() Builder
	WithList(list []Content) Builder
	Now() (Contents, error)
}

// Contents represents contents
type Contents interface {
	List() []Content
}

// ContentAdapter represents a connection adapter
type ContentAdapter interface {
	ToContent(ins Content) ([]byte, error)
	ToInstance(content []byte) (Content, error)
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithHash(hash hash.Hash) ContentBuilder
	WithData(data []byte) ContentBuilder
	WithKind(kind uint) ContentBuilder
	Now() (Content, error)
}

// Content represents a connection's content
type Content interface {
	Hash() hash.Hash
	Data() []byte
	Kind() uint
}
