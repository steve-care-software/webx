package signs

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

type sign struct {
	hash     hash.Hash
	create   creates.Create
	validate validates.Validate
}

func createSignWithCreate(
	hash hash.Hash,
	create creates.Create,
) Sign {
	return createSignInternally(hash, create, nil)
}

func createSignWithValidate(
	hash hash.Hash,
	validate validates.Validate,
) Sign {
	return createSignInternally(hash, nil, validate)
}

func createSignInternally(
	hash hash.Hash,
	create creates.Create,
	validate validates.Validate,
) Sign {
	out := sign{
		hash:     hash,
		create:   create,
		validate: validate,
	}

	return &out
}

// Hash returns the hash
func (obj *sign) Hash() hash.Hash {
	return obj.hash
}

// IsCreate returns true if create, false otherwise
func (obj *sign) IsCreate() bool {
	return obj.create != nil
}

// Create returns the create, if any
func (obj *sign) Create() creates.Create {
	return obj.create
}

// IsValidate returns true if validate, false otherwise
func (obj *sign) IsValidate() bool {
	return obj.validate != nil
}

// Validate returns the validate, if any
func (obj *sign) Validate() validates.Validate {
	return obj.validate
}
