package elements

import "errors"

type elementBuilder struct {
	rule        string
	syscall     string
	instruction string
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		rule:        "",
		syscall:     "",
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
func (app *elementBuilder) WithSyscall(syscall string) ElementBuilder {
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

	if app.syscall != "" {
		return createElementWithSyscall(app.syscall), nil
	}

	if app.instruction != "" {
		return createElementWithInstruction(app.instruction), nil
	}

	return nil, errors.New("the Element is invalid")
}
