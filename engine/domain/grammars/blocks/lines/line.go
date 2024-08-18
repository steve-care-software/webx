package lines

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
)

type line struct {
	tokens      []string
	execution   executions.Execution
	replacement string
}

func createLineWithExecutionAndReplacement(
	tokens []string,
	execution executions.Execution,
	replacement string,
) Line {
	return createLineInternally(tokens, execution, "")
}

func createLineWithExecution(
	tokens []string,
	execution executions.Execution,
) Line {
	return createLineInternally(tokens, execution, "")
}

func createLineWithReplacement(
	tokens []string,
	replacement string,
) Line {
	return createLineInternally(tokens, nil, replacement)
}

func createLineInternally(
	tokens []string,
	execution executions.Execution,
	replacement string,
) Line {
	out := line{
		tokens:      tokens,
		execution:   execution,
		replacement: replacement,
	}

	return &out
}

// Tokens returns the tokens
func (obj *line) Tokens() []string {
	return obj.tokens
}

// HasExecution returns true if there is an execution, false otherwise
func (obj *line) HasExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *line) Execution() executions.Execution {
	return obj.execution
}

// HasReplacement returns true if there is a replacement, false otherwise
func (obj *line) HasReplacement() bool {
	return obj.replacement != ""
}

// Replacement returns the replacement, if any
func (obj *line) Replacement() string {
	return obj.replacement
}
