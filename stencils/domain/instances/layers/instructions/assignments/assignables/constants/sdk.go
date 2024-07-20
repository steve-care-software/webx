package constants

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewConstantBuilder creates a new constant builder
func NewConstantBuilder() ConstantBuilder {
	hashAdapter := hash.NewAdapter()
	return createConstantBuilder(
		hashAdapter,
	)
}

// Adapter represents the constant adapter
type Adapter interface {
	InstanceToBytes(ins Constant) ([]byte, error)
	BytesToInstance(bytes []byte) (Constant, error)
	InstancesToBytes(ins Constants) ([]byte, error)
	BytesToInstances(bytes []byte) (Constants, error)
}

// Builder represents a constants builder
type Builder interface {
	Create() Builder
	WithList(list []Constant) Builder
	Now() (Constants, error)
}

// Constants represents constants
type Constants interface {
	Hash() hash.Hash
	List() []Constant
}

// ConstantBuilder represents a constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithBool(boolValue bool) ConstantBuilder
	WithString(strValue string) ConstantBuilder
	WithInt(intValue int) ConstantBuilder
	WithUint(uintValue uint) ConstantBuilder
	WithFloat(floatVal float64) ConstantBuilder
	WithList(list Constants) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant assignable
type Constant interface {
	Hash() hash.Hash
	IsBool() bool
	Bool() *bool
	IsString() bool
	String() *string
	IsInt() bool
	Int() *int
	IsUint() bool
	Uint() *uint
	IsFloat() bool
	Float() *float64
	IsList() bool
	List() Constants
}
