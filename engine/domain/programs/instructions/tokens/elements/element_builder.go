package elements

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls"
)

type elementBuilder struct {
	rule        string
	syscall     syscalls.Syscall
	instruction string
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		rule:        "",
		syscall:     nil,
		instruction: "",
	}

	return &out
}

// Create initializes the elementBuilder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithRule adds a rule to the elementBuilder
func (app *elementBuilder) WithRule(rule string) ElementBuilder {
	app.rule = rule
	return app
}

// WithSyscall adds a syscall to the elementBuilder
func (app *elementBuilder) WithSyscall(syscall syscalls.Syscall) ElementBuilder {
	app.syscall = syscall
	return app
}

// WithInstruction adds an instruction to the elementBuilder
func (app *elementBuilder) WithInstruction(instruction string) ElementBuilder {
	app.instruction = instruction
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.rule != "" {
		return createElementWithRule(app.rule), nil
	}

	if app.syscall != nil {
		return createElementWithSyscall(app.syscall), nil
	}

	if app.instruction != "" {
		return createElementWithInstruction(app.instruction), nil
	}

	return nil, errors.New("the Element is invalid")
}
