package variables

const (
	// KindUint8 represents the uint8
	KindUint8 (uint8) = iota

	// KindUint16 represents the uint16
	KindUint16

	// KindUint32 represents the uint32
	KindUint32

	// KindUint64 represents the uint64
	KindUint64

	// KindInt8 represents the int8
	KindInt8

	// KindInt16 represents the int16
	KindInt16

	// KindInt32 represents the int32
	KindInt32

	// KindInt64 represents the int64
	KindInt64

	// KindFloat32 represents the float32
	KindFloat32

	// KindFloat64 represents the float64
	KindFloat64

	// KindBool represents the bool
	KindBool

	// KindString represents the string
	KindString

	// KindStack represents the stack
	KindStack
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewVariableBuilder creates a new variable builder
func NewVariableBuilder() VariableBuilder {
	return createVariableBuilder()
}

// Builder represents a variables builder
type Builder interface {
	Create() Builder
	WithList(list []Variable) Builder
	Now() (Variables, error)
}

// Variables represents the variables
type Variables interface {
	List() []Variable
	Fetch(name string) (Variable, error)
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithName(name string) VariableBuilder
	WithValue(value any) VariableBuilder
	WithKind(kind uint8) VariableBuilder
	ReplaceIfExists() VariableBuilder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	Name() string
	Value() any
	Kind() uint8
	ReplaceIfExists() bool
}
