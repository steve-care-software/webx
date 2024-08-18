package lines

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/replacements"
)

type line struct {
	tokens      []string
	execution   executions.Execution
	replacement replacements.Replacement
}

func createLineWithExecutionAndReplacement(
	tokens []string,
	execution executions.Execution,
	replacement replacements.Replacement,
) Line {
	return createLineInternally(tokens, execution, replacement)
}

func createLineWithExecution(
	tokens []string,
	execution executions.Execution,
) Line {
	return createLineInternally(tokens, execution, nil)
}

func createLineWithReplacement(
	tokens []string,
	replacement replacements.Replacement,
) Line {
	return createLineInternally(tokens, nil, replacement)
}

func createLineInternally(
	tokens []string,
	execution executions.Execution,
	replacement replacements.Replacement,
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
	return obj.replacement != nil
}

// Replacement returns the replacement, if any
func (obj *line) Replacement() replacements.Replacement {
	return obj.replacement
}
