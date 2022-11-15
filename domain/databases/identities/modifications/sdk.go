package modifications

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a modification adapter
type Adapter interface {
	ToContents(list []Modification) ([][]byte, error)
	ToContent(ins Modification) ([]byte, error)
}

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier entities.Identifier) Builder
	WithName(name string) Builder
	WithSignature(sig []byte) Builder
	WithEncryption(enc []byte) Builder
	Now() (Modification, error)
}

// Modification represents a modifucation
type Modification interface {
	Identifier() entities.Identifier
	Content() Content
}

// Content represents a modification content
type Content interface {
	HasName() bool
	Name() string
	HasSignature() bool
	Signature() []byte
	HasEncryption() bool
	Encryption() []byte
}
