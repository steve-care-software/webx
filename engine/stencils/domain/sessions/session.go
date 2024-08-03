package sessions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

type session struct {
	hash       hash.Hash
	executions []executions.Executions
	token      hash.Hash
}

func createSession(
	hash hash.Hash,
	executions []executions.Executions,
	token hash.Hash,
) Session {
	out := session{
		hash:       hash,
		executions: executions,
		token:      token,
	}

	return &out
}

// Hash returns the hash
func (obj *session) Hash() hash.Hash {
	return obj.hash
}

// Executions returns the executions
func (obj *session) Executions() []executions.Executions {
	return obj.executions
}

// Token returns the token
func (obj *session) Token() hash.Hash {
	return obj.token
}
