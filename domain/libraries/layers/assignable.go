package layers

import "github.com/steve-care-software/identity/domain/hash"

type assignable struct {
	hash      hash.Hash
	bytes     Bytes
	identity  Identity
	execution Execution
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil, nil)
}

func createAssignableWithIdentity(
	hash hash.Hash,
	identity Identity,
) Assignable {
	return createAssignableInternally(hash, nil, identity, nil)
}

func createAssignableWithexecution(
	hash hash.Hash,
	execution Execution,
) Assignable {
	return createAssignableInternally(hash, nil, nil, execution)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes Bytes,
	identity Identity,
	execution Execution,
) Assignable {
	out := assignable{
		hash:      hash,
		bytes:     bytes,
		identity:  identity,
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
func (obj *assignable) Bytes() Bytes {
	return obj.bytes
}

// IsIdentity returns true if there is identity, false otherwise
func (obj *assignable) IsIdentity() bool {
	return obj.identity != nil
}

// Identity returns the identity, if any
func (obj *assignable) Identity() Identity {
	return obj.identity
}

// IsExecution returns true if there is execution, false otherwise
func (obj *assignable) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *assignable) Execution() Execution {
	return obj.execution
}
