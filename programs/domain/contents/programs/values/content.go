package values

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type content struct {
	pInput     *uint
	pValue     *hash.Hash
	constant   []byte
	pExecution *hash.Hash
	pProgram   *hash.Hash
}

func createContentWithInput(
	pInput *uint,
) Content {
	return createContentInternally(pInput, nil, nil, nil, nil)
}

func createContentWithValue(
	pValue *hash.Hash,
) Content {
	return createContentInternally(nil, pValue, nil, nil, nil)
}

func createContentWithConstant(
	constant []byte,
) Content {
	return createContentInternally(nil, nil, constant, nil, nil)
}

func createContentWithExecution(
	pExecution *hash.Hash,
) Content {
	return createContentInternally(nil, nil, nil, pExecution, nil)
}

func createContentWithProgram(
	pProgram *hash.Hash,
) Content {
	return createContentInternally(nil, nil, nil, nil, pProgram)
}

func createContentInternally(
	pInput *uint,
	pValue *hash.Hash,
	constant []byte,
	pExecution *hash.Hash,
	pProgram *hash.Hash,
) Content {
	out := content{
		pInput:     pInput,
		pValue:     pValue,
		constant:   constant,
		pExecution: pExecution,
		pProgram:   pProgram,
	}

	return &out
}

// IsInput returns true if there is an input, false otherwise
func (obj *content) IsInput() bool {
	return obj.pInput != nil
}

// Input returns the input, if any
func (obj *content) Input() *uint {
	return obj.pInput
}

// IsValue returns true if there is an value, false otherwise
func (obj *content) IsValue() bool {
	return obj.pValue != nil
}

// Value returns the value, if any
func (obj *content) Value() *hash.Hash {
	return obj.pValue
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *content) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *content) Constant() []byte {
	return obj.constant
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *content) IsExecution() bool {
	return obj.pExecution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() *hash.Hash {
	return obj.pExecution
}

// IsProgram returns true if there is a program, false otherwise
func (obj *content) IsProgram() bool {
	return obj.pProgram != nil
}

// Program returns the program, if any
func (obj *content) Program() *hash.Hash {
	return obj.pProgram
}
