package executables

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the executable adapter
type Adapter interface {
	ToBytes(ins Executable) ([]byte, error)
	ToInstance(bytes []byte) (Executable, error)
}

// Builder represents an executable builder
type Builder interface {
	Create() Builder
	WithLocal(local string) Builder
	WithRemote(remote string) Builder
	Now() (Executable, error)
}

// Executable represents an executable
type Executable interface {
	Hash() hash.Hash
	IsLocal() bool
	Local() string
	IsRemote() bool
	Remote() string
}
