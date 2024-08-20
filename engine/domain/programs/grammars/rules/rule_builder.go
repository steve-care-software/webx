package rules

import "errors"

type ruleBuilder struct {
	name  string
	bytes []byte
}

func createRuleBuilder() RuleBuilder {
	out := ruleBuilder{
		name:  "",
		bytes: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleBuilder) Create() RuleBuilder {
	return createRuleBuilder()
}

// WithName adds a name to the builder
func (app *ruleBuilder) WithName(name string) RuleBuilder {
	app.name = name
	return app
}

// WithBytes add bytes to the builder
func (app *ruleBuilder) WithBytes(bytes []byte) RuleBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new Rule instance
func (app *ruleBuilder) Now() (Rule, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes are mandatory in order to build a Rule instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Rule instance")
	}

	return createRule(
		app.name,
		app.bytes,
	), nil
}
