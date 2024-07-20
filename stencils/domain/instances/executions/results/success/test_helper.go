package success

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/outputs/kinds"
)

// NewSuccessForTests creates a new success for tests
func NewSuccessForTests(output outputs.Output, kind kinds.Kind) Success {
	ins, err := NewBuilder().Create().WithOutput(output).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
