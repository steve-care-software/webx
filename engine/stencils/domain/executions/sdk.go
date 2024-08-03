package executions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Identifier() uint64
	Value() executions.Execution
}
