package constants

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Adapter represents the constant adapter
type Adapter interface {
	ToBytes(ins Constant) ([]byte, error)
	ToInstance(bytes []byte) (Constant, error)
}

// Builder represents a constant builder
type Builder interface {
	Create() Builder
	WithBool(boolValue bool) Builder
	WithString(strValue string) Builder
	WithInt(intValue int) Builder
	WithUint(uintValue uint) Builder
	WithFoat(floatVal float64) Builder
	WithList(list []string) Builder
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
	Foat() *float64
	IsList() bool
	List() []Constant
}
