package lines

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type line struct {
	tokens      tokens.Tokens
	execution   executions.Execution
	replacement elements.Element
}

func createLineWithExecution(
	tokens tokens.Tokens,
	execution executions.Execution,
) Line {
	return createLineInternally(tokens, execution, nil)
}

func createLineWithReplacement(
	tokens tokens.Tokens,
	replacement elements.Element,
) Line {
	return createLineInternally(tokens, nil, replacement)
}

func createLine(
	tokens tokens.Tokens,
) Line {
	return createLineInternally(tokens, nil, nil)
}

func createLineInternally(
	tokens tokens.Tokens,
	execution executions.Execution,
	replacement elements.Element,
) Line {
	out := line{
		tokens:      tokens,
		execution:   execution,
		replacement: replacement,
	}

	return &out
}

// Tokens returns the tokens
func (obj *line) Tokens() tokens.Tokens {
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
func (obj *line) Replacement() elements.Element {
	return obj.replacement
}
