package executions

import "github.com/steve-care-software/historydb/domain/hash"

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	IsList() Builder
	Now() (Execution, error)
}

// Execution represents an execution instruction
type Execution interface {
	Hash() hash.Hash
	IsList() bool
	List() string
}
