package sessions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

// NewSessionForTests creates a new session for tests
func NewSessionForTests(hash hash.Hash, executions []executions.Executions) Session {
	ins, err := NewBuilder().Create().WithHash(hash).WithExecutions(executions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
