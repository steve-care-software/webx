package executions

const (
	// TypeCastIntHeight represents @type.cast.int8(myToken)
	TypeCastIntHeight (uint16) = iota

	// TypeCastIntSixteen represents @type.cast.int16(myToken)
	TypeCastIntSixteen

	// TypeCastIntThirtyTwo represents @type.cast.int32(myToken)
	TypeCastIntThirtyTwo

	// TypeCastIntSixtyFour represents @type.cast.int64(myToken)
	TypeCastIntSixtyFour

	// TypeCastUintHeight represents @type.cast.int8(myToken)
	TypeCastUintHeight

	// TypeCastUintSixteen represents @type.cast.int16(myToken)
	TypeCastUintSixteen

	// TypeCastUintThirtyTwo represents @type.cast.int32(myToken)
	TypeCastUintThirtyTwo

	// TypeCastUintSixtyFour represents @type.cast.int64(myToken)
	TypeCastUintSixtyFour

	// TypeCastFloatThirtyTwo represents @type.cast.float32(myToken)
	TypeCastFloatThirtyTwo

	// TypeCastFloatSixtyFour represents @type.cast.float64(myToken)
	TypeCastFloatSixtyFour

	// TypeCastFloatString represents @type.cast.string(myToken)
	TypeCastFloatString

	// TypeCastFloatBytes represents @type.cast.bytes(myToken)
	TypeCastFloatBytes

	// MathOperationArithmeticAddition represents @math.operation.atithmetic.addition(intSixtyFour, intSixtyFour)
	MathOperationArithmeticAddition

	// MathOperationArithmeticSubstraction represents @math.operation.atithmetic.substraction(intSixtyFour, intSixtyFour)
	MathOperationArithmeticSubstraction

	// StackAssignment represents @stack.assignment(variableName, assignable)
	StackAssignment
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder([]uint16{
		TypeCastIntHeight,
		TypeCastIntSixteen,
		TypeCastIntThirtyTwo,
		TypeCastIntSixtyFour,
		TypeCastUintHeight,
		TypeCastUintSixteen,
		TypeCastUintThirtyTwo,
		TypeCastUintSixtyFour,
		TypeCastFloatThirtyTwo,
		TypeCastFloatSixtyFour,
		TypeCastFloatString,
		TypeCastFloatBytes,
		MathOperationArithmeticAddition,
		MathOperationArithmeticSubstraction,
		StackAssignment,
	})
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithTokens(tokens []string) Builder
	WithFuncFlag(fnFlag uint16) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Tokens() []string
	FuncFlag() uint16
}
