package histories

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// Adapter represents an histories adapter
type Adapter interface {
	ToContent(ins History) ([]byte, error)
	ToHistory(content []byte) (History, error)
}

// Builder represents an history builder
type Builder interface {
	Create() Builder
	WithCommit(commit hash.Hash) Builder
	WithDifficulty(difficulty uint) Builder
	Now() (History, error)
}

// History represents a commit history
type History interface {
	Commit() hash.Hash
	Difficulty() uint
}
