package values

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type content struct {
	pInput     *uint
	constant   []byte
	pExecution *hash.Hash
	pProgram   *hash.Hash
}

func createContentWithInput(
	pInput *uint,
) Content {
	return createContentInternally(pInput, nil, nil, nil)
}

func createContentWithConstant(
	constant []byte,
) Content {
	return createContentInternally(nil, constant, nil, nil)
}

func createContentWithExecution(
	pExecution *hash.Hash,
) Content {
	return createContentInternally(nil, nil, pExecution, nil)
}

func createContentWithProgram(
	pProgram *hash.Hash,
) Content {
	return createContentInternally(nil, nil, nil, pProgram)
}

func createContentInternally(
	pInput *uint,
	constant []byte,
	pExecution *hash.Hash,
	pProgram *hash.Hash,
) Content {
	out := content{
		pInput:     pInput,
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
