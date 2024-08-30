package lines

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/processors"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
)

type line struct {
	tokens    tokens.Tokens
	processor processors.Processor
	syscall   executions.Execution
}

func createLine(
	tokens tokens.Tokens,
) Line {
	return createLineInternally(tokens, nil, nil)
}

func createLineWithProcessor(
	tokens tokens.Tokens,
	processor processors.Processor,
) Line {
	return createLineInternally(tokens, processor, nil)
}

func createLineWithSyscall(
	tokens tokens.Tokens,
	syscall executions.Execution,
) Line {
	return createLineInternally(tokens, nil, syscall)
}

func createLineWithProcessorAndSyscall(
	tokens tokens.Tokens,
	processor processors.Processor,
	syscall executions.Execution,
) Line {
	return createLineInternally(tokens, processor, syscall)
}

func createLineInternally(
	tokens tokens.Tokens,
	processor processors.Processor,
	syscall executions.Execution,
) Line {
	out := line{
		tokens:    tokens,
		processor: processor,
		syscall:   syscall,
	}

	return &out
}

// Tokens returns the tokens
func (obj *line) Tokens() tokens.Tokens {
	return obj.tokens
}

// HasProcessor returns true if there is a processor, false otherwise
func (obj *line) HasProcessor() bool {
	return obj.processor != nil
}

// Processor returns the processor, if any
func (obj *line) Processor() processors.Processor {
	return obj.processor
}

// HasSyscall returns true if there is a syscall, false otherwise
func (obj *line) HasSyscall() bool {
	return obj.syscall != nil
}

// Syscall returns the syscall, if any
func (obj *line) Syscall() executions.Execution {
	return obj.syscall
}
