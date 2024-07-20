package executions

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type executions struct {
	hash hash.Hash
	list []Execution
}

func createExecutions(
	hash hash.Hash,
	list []Execution,
) Executions {
	return &executions{
		hash: hash,
		list: list,
	}
}

// Hash returns the hash
func (obj *executions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *executions) List() []Execution {
	return obj.list
}
