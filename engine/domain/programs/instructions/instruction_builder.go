package instructions

import (
	"errors"
)

type instructionBuilder struct {
	block   string
	pLine   *uint
	tokens  Tokens
	syscall Syscall
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		block:   "",
		pLine:   nil,
		tokens:  nil,
		syscall: nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithBlock adds a block to the builder
func (app *instructionBuilder) WithBlock(block string) InstructionBuilder {
	app.block = block
	return app
}

// WithLine adds a line to the builder
func (app *instructionBuilder) WithLine(line uint) InstructionBuilder {
	app.pLine = &line
	return app
}

// WithTokens add tokens to the builder
func (app *instructionBuilder) WithTokens(tokens Tokens) InstructionBuilder {
	app.tokens = tokens
	return app
}

// WithSyscall adds a syscall to the builder
func (app *instructionBuilder) WithSyscall(syscall Syscall) InstructionBuilder {
	app.syscall = syscall
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.block == "" {
		return nil, errors.New("the block is mandatory in order to build an Instruction")
	}

	if app.pLine == nil {
		return nil, errors.New("the line is mandatory in order to build an Instruction")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens is mandatory in order to build an Instruction")
	}

	if app.syscall != nil {
		return createInstructionWithSyscall(
			app.block,
			*app.pLine,
			app.tokens,
			app.syscall,
		), nil
	}

	return createInstruction(app.block, *app.pLine, app.tokens), nil
}
