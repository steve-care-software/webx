package assignables

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
)

type assignable struct {
	hash     hash.Hash
	bytes    bytes.Bytes
	constant constants.Constant
	account  accounts.Account
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes bytes.Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil, nil)
}

func createAssignableWithConstant(
	hash hash.Hash,
	constant constants.Constant,
) Assignable {
	return createAssignableInternally(hash, nil, constant, nil)
}

func createAssignableWithAccount(
	hash hash.Hash,
	account accounts.Account,
) Assignable {
	return createAssignableInternally(hash, nil, nil, account)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes bytes.Bytes,
	constant constants.Constant,
	account accounts.Account,
) Assignable {
	out := assignable{
		hash:     hash,
		bytes:    bytes,
		constant: constant,
		account:  account,
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

// IsConstant returns true if there is constant, false otherwise
func (obj *assignable) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *assignable) Constant() constants.Constant {
	return obj.constant
}

// IsAccount returns true if there is an account, false otherwise
func (obj *assignable) IsAccount() bool {
	return obj.account != nil
}

// Account returns the account, if any
func (obj *assignable) Account() accounts.Account {
	return obj.account
}
