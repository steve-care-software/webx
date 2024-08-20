package elements

import "errors"

type builder struct {
	rule        string
	syscall     string
	instruction string
}

func createBuilder() Builder {
	out := builder{
		rule:        "",
		syscall:     "",
		instruction: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// WithSyscall adds a syscall to the builder
func (app *builder) WithSyscall(syscall string) Builder {
	app.syscall = syscall
	return app
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(instruction string) Builder {
	app.instruction = instruction
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
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
