package processors

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type processor struct {
	execution   executions.Execution
	replacement elements.Element
}

func createProcessorWithExecution(
	execution executions.Execution,
) Processor {
	return createProcessorInternally(execution, nil)
}

func createProcessorWithReplacement(
	replacement elements.Element,
) Processor {
	return createProcessorInternally(nil, replacement)
}

func createProcessorInternally(
	execution executions.Execution,
	replacement elements.Element,
) Processor {
	out := processor{
		execution:   execution,
		replacement: replacement,
	}

	return &out
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *processor) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *processor) Execution() executions.Execution {
	return obj.execution
}

// IsReplacement returns true if there is a replacement, false otherwise
func (obj *processor) IsReplacement() bool {
	return obj.replacement != nil
}

// Replacement returns the replacement, if any
func (obj *processor) Replacement() elements.Element {
	return obj.replacement
}
