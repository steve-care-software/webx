package votes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

type vote struct {
	hash     hash.Hash
	create   creates.Create
	validate validates.Validate
}

func createVoteWithCreate(
	hash hash.Hash,
	create creates.Create,
) Vote {
	return createVoteInternally(hash, create, nil)
}

func createVoteWithValidate(
	hash hash.Hash,
	validate validates.Validate,
) Vote {
	return createVoteInternally(hash, nil, validate)
}

func createVoteInternally(
	hash hash.Hash,
	create creates.Create,
	validate validates.Validate,
) Vote {
	out := vote{
		hash:     hash,
		create:   create,
		validate: validate,
	}

	return &out
}

// Hash returns the hash
func (obj *vote) Hash() hash.Hash {
	return obj.hash
}

// IsCreate returns true if create, false otherwise
func (obj *vote) IsCreate() bool {
	return obj.create != nil
}

// Create returns create, if any
func (obj *vote) Create() creates.Create {
	return obj.create
}

// IsValidate returns true if validate, false otherwise
func (obj *vote) IsValidate() bool {
	return obj.validate != nil
}

// Validate returns validate, if any
func (obj *vote) Validate() validates.Validate {
	return obj.validate
}
