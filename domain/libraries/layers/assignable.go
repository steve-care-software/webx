package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/executions"
)

type assignable struct {
	hash      hash.Hash
	bytes     bytes.Bytes
	execution executions.Execution
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes bytes.Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil)
}

func createAssignableWithexecution(
	hash hash.Hash,
	execution executions.Execution,
) Assignable {
	return createAssignableInternally(hash, nil, execution)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes bytes.Bytes,
	execution executions.Execution,
) Assignable {
	out := assignable{
		hash:      hash,
		bytes:     bytes,
		execution: execution,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *assignable) Bytes() bytes.Bytes {
	return obj.bytes
}

// IsExecution returns true if there is execution, false otherwise
func (obj *assignable) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *assignable) Execution() executions.Execution {
	return obj.execution
}
