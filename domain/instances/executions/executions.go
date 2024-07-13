package executions

import "github.com/steve-care-software/historydb/domain/hash"

type executions struct {
	hash hash.Hash
	list []Execution
}

func createExecutions(
	hash hash.Hash,
	list []Execution,
) Executions {
	out := executions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *executions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *executions) List() []Execution {
	return obj.list
}

// Databases returns the executed database paths
func (obj *executions) Databases() ([][]string, error) {
	return nil, nil
}

// basePath returns the executed links paths
func (obj *executions) Links(basePath []string) ([][]string, error) {
	return nil, nil
}
