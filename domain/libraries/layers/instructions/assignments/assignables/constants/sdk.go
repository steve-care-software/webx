package constants

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewConstantBuilder creates a new constant builder
func NewConstantBuilder() ConstantBuilder {
	hashAdapter := hash.NewAdapter()
	return createConstantBuilder(
		hashAdapter,
	)
}

// ConstantBuilder represents a constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithBool(boolValue bool) ConstantBuilder
	WithBytes(bytes []byte) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant assignable
type Constant interface {
	Hash() hash.Hash
	IsBool() bool
	Bool() *bool
	IsBytes() bool
	Bytes() []byte
}
