package cursors

import (
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

const invalidKindErrPattern = "the provided kind (%d) is invalid"

// NewApplication creates a new application
func NewApplication() Application {
	variablesBuilder := variables.NewBuilder()
	variableBuilder := variables.NewVariableBuilder()
	return createApplication(
		variablesBuilder,
		variableBuilder,
	)
}

// Application represents the cursor application
type Application interface {
	Fingerprint() (variables.Variables, error)
	Amount(kind uint8) (*uint, error)
	Fetch(index uint, kind uint8) (any, error)
	Push(value any, kind uint8) error
	PushAsStringBytes(valueStrAsBytes []byte, kind uint8) error
	Save(index uint, kind uint8, variable string, replaceIfExists bool) error
	Remove(index uint, kind uint8) error
	Clear(kind uint8) error
	ClearAll() error
}
