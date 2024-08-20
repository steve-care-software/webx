package programs

import (
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

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

// Application represents the program application
type Application interface {
	Parse(grammar grammars.Grammar, input []byte) (programs.Program, error)
	Compile(program programs.Program) (nfts.NFT, error)
	Decompile(nft nfts.NFT) (programs.Program, error)
	Compose(program programs.Program) ([]byte, error)
	Interpret(program programs.Program) (stacks.Stack, error)
}
