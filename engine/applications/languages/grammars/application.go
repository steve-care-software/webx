package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
)

type application struct {
	ruleBuilder                rules.RuleBuilder
	ruleNameValueSeparator     byte
	possibleRuleNameCharacters []byte
	ruleNameSeparator          byte
	ruleValuePrefix            byte
	ruleValueSuffix            byte
	ruleValueEscape            byte
}

func createApplication(
	ruleBuilder rules.RuleBuilder,
	ruleNameValueSeparator byte,
	possibleRuleNameCharacters []byte,
	ruleNameSeparator byte,
	ruleValuePrefix byte,
	ruleValueSuffix byte,
	ruleValueEscape byte,
) Application {
	out := application{
		ruleBuilder:                ruleBuilder,
		ruleNameValueSeparator:     ruleNameValueSeparator,
		possibleRuleNameCharacters: possibleRuleNameCharacters,
		ruleNameSeparator:          ruleNameSeparator,
		ruleValuePrefix:            ruleValuePrefix,
		ruleValueSuffix:            ruleValueSuffix,
		ruleValueEscape:            ruleValueEscape,
	}

	return &out
}

// Parse parses the input and returns a grammar instance
func (app *application) Parse(lexedInput []byte) (grammars.Grammar, error) {
	return nil, nil
}

// Compile compiles a grammar to an AST
func (app *application) Compile(grammar grammars.Grammar) (asts.AST, error) {
	return nil, nil
}

// Decompile decompiles an AST to a grammar instance
func (app *application) Decompile(ast asts.AST) (grammars.Grammar, error) {
	return nil, nil
}

// Compose composes a grammar instance to a grammar input
func (app *application) Compose(grammar grammars.Grammar) ([]byte, error) {
	return nil, nil
}

func (app *application) bytesToRule(input []byte) (rules.Rule, []byte, error) {
	name, value, remaining, err := bytesToRuleNameAndValue(
		input,
		app.ruleNameValueSeparator,
		app.possibleRuleNameCharacters,
		app.ruleNameSeparator,
		app.ruleValuePrefix,
		app.ruleValueSuffix,
		app.ruleValueEscape,
	)

	if err != nil {
		return nil, nil, err
	}

	ins, err := app.ruleBuilder.Create().
		WithName(string(name)).
		WithBytes(value).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}
