package modifications

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a modification adapter
type Adapter interface {
	ToContents(list []Modification) ([][]byte, error)
	ToContent(ins Modification) ([]byte, error)
	ToModifications(contents [][]byte) ([]Modification, error)
	ToModification(content []byte) (Modification, error)
}

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSignature(sig []byte) Builder
	WithEncryption(enc []byte) Builder
	Now() (Modification, error)
}

// Modification represents a modifucation
type Modification interface {
	Hash() hash.Hash
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
