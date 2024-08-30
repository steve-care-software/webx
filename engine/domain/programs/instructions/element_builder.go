package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
)

type elementBuilder struct {
	rule        rules.Rule
	instruction Instruction
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		rule:        nil,
		instruction: nil,
	}

	return &out
}

// Create initializes the elementBuilder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithRule adds a rule to the elementBuilder
func (app *elementBuilder) WithRule(rule rules.Rule) ElementBuilder {
	app.rule = rule
	return app
}

// WithInstruction adds an instruction to the elementBuilder
func (app *elementBuilder) WithInstruction(instruction Instruction) ElementBuilder {
	app.instruction = instruction
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.rule != nil {
		return createElementWithRule(app.rule), nil
	}

	if app.instruction != nil {
		return createElementWithInstruction(app.instruction), nil
	}

	return nil, errors.New("the Element is invalid")
}
