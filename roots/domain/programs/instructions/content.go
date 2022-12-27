package instructions

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type content struct {
	pValue     *hash.Hash
	pExecution *hash.Hash
}

func createContentWithValue(
	pValue *hash.Hash,
) Content {
	return createContentInternally(pValue, nil)
}

func createContentWithExecution(
	pExecution *hash.Hash,
) Content {
	return createContentInternally(nil, pExecution)
}

func createContentInternally(
	pValue *hash.Hash,
	pExecution *hash.Hash,
) Content {
	out := content{
		pValue:     pValue,
		pExecution: pExecution,
	}

	return &out
}

// IsValue returns true if there is an value, false otherwise
func (obj *content) IsValue() bool {
	return obj.pValue != nil
}

// Value returns the value, if any
func (obj *content) Value() *hash.Hash {
	return obj.pValue
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *content) IsExecution() bool {
	return obj.pExecution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() *hash.Hash {
	return obj.pExecution
}
