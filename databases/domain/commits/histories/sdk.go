package histories

import (
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	historyBuilder := NewHistoryBuilder()
	return createAdapter(builder, historyBuilder)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewHistoryBuilder creates an history builder
func NewHistoryBuilder() HistoryBuilder {
	return createHistoryBuilder()
}

// Adapter represents an histories adapter
type Adapter interface {
	ToHistories(commits []commits.Commit) (Histories, error)
}

// Builder represents a histories builder
type Builder interface {
	Create() Builder
	WithList(list []History) Builder
	Now() (Histories, error)
}

// Histories represents an histories
type Histories interface {
	List() []History
	Compare(ins Histories) ([]History, error)
}

// HistoryBuilder represents an history builder
type HistoryBuilder interface {
	Create() HistoryBuilder
	WithCommit(commit hash.Hash) HistoryBuilder
	WithScore(score uint) HistoryBuilder
	Now() (History, error)
}

// History represents a commit history
type History interface {
	Commit() hash.Hash
	Score() uint
}
