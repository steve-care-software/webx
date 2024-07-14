package executes

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

// Execute represents an execute
type Execute interface {
	Context() string
	Input() inputs.Input
	Return() string
	HasLayer() bool
	Layer() inputs.Input
}
