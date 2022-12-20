package histories

import (
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// Adapter represents an histories adapter
type Adapter interface {
	ToHistories(commits []commits.Commit) (Histories, error)
}

// Builder represents a histories builder
type Builder interface {
	Create() Builder
	WithList(list []History) Builder
	WithMatrix(matrix []Histories) Builder
	Now() (Histories, error)
}

// Histories represents an histories
type Histories interface {
	List() ([]History, error)
	Compare(ins Histories) ([]History, error)
}

// History represents a commit history
type History interface {
	Commit() hash.Hash
	Difficulty() uint
}
