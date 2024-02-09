package stacks

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type assignable struct {
	pBool *bool
	bytes []byte
	hash  hash.Hash
}

func createAssignableWithBool(
	pBool *bool,
) Assignable {
	return createAssignableInternally(
		pBool,
		nil,
		nil,
	)
}

func createAssignableWithBytes(
	bytes []byte,
) Assignable {
	return createAssignableInternally(
		nil,
		bytes,
		nil,
	)
}

func createAssignableWithHash(
	hash hash.Hash,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		hash,
	)
}

func createAssignableInternally(
	pBool *bool,
	bytes []byte,
	hash hash.Hash,
) Assignable {
	out := assignable{
		pBool: pBool,
		bytes: bytes,
		hash:  hash,
	}

	return &out
}

// IsBool returns true if bool, false otherwise
func (obj *assignable) IsBool() bool {
	return obj.pBool != nil
}

// Bool returns bool, if any
func (obj *assignable) Bool() *bool {
	return obj.pBool
}

// IsBytes returns true if bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns bytes, if any
func (obj *assignable) Bytes() []byte {
	return obj.bytes
}

// IsHash returns true if hash, false otherwise
func (obj *assignable) IsHash() bool {
	return obj.hash != nil
}

// Hash returns hash, if any
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}
