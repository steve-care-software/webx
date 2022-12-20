package contents

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

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
