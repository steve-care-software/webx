package assignments

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"
)

type assignment struct {
	hash       hash.Hash
	name       string
	assignable assignables.Assignable
}

func createAssignment(
	hash hash.Hash,
	name string,
	assignable assignables.Assignable,
) Assignment {
	out := assignment{
		hash:       hash,
		name:       name,
		assignable: assignable,
	}

	return &out
}

// Hash returns the hash
func (obj *assignment) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *assignment) Name() string {
	return obj.name
}

// Assignable returns the assignable
func (obj *assignment) Assignable() assignables.Assignable {
	return obj.assignable
}
