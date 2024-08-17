package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
)

type application struct {
	ruleBuilder              rules.RuleBuilder
	cardinalityBuilder       cardinalities.Builder
	ruleNameValueSeparator   byte
	possibleUpperCaseLetters []byte
	possibleNumbers          []byte
	ruleNameSeparator        byte
	ruleValuePrefix          byte
	ruleValueSuffix          byte
	ruleValueEscape          byte
	cardinalityOpen          byte
	cardinalityClose         byte
	cardinalitySeparator     byte
	cardinalityZeroPlus      byte
	cardinalityOnePlus       byte
}

func createApplication(
	ruleBuilder rules.RuleBuilder,
	cardinalityBuilder cardinalities.Builder,
	ruleNameValueSeparator byte,
	possibleUpperCaseLetters []byte,
	possibleNumbers []byte,
	ruleNameSeparator byte,
	ruleValuePrefix byte,
	ruleValueSuffix byte,
	ruleValueEscape byte,
	cardinalityOpen byte,
	cardinalityClose byte,
	cardinalitySeparator byte,
	cardinalityZeroPlus byte,
	cardinalityOnePlus byte,
) Application {
	out := application{
		ruleBuilder:              ruleBuilder,
		cardinalityBuilder:       cardinalityBuilder,
		ruleNameValueSeparator:   ruleNameValueSeparator,
		possibleUpperCaseLetters: possibleUpperCaseLetters,
		possibleNumbers:          possibleNumbers,
		ruleNameSeparator:        ruleNameSeparator,
		ruleValuePrefix:          ruleValuePrefix,
		ruleValueSuffix:          ruleValueSuffix,
		ruleValueEscape:          ruleValueEscape,
		cardinalityOpen:          cardinalityOpen,
		cardinalityClose:         cardinalityClose,
		cardinalitySeparator:     cardinalitySeparator,
		cardinalityZeroPlus:      cardinalityZeroPlus,
		cardinalityOnePlus:       cardinalityOnePlus,
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

func (app *application) bytesToCardinality(input []byte) (cardinalities.Cardinality, []byte, error) {
	retMin, pRetMax, retRemaining, err := bytesToMinMax(
		input,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.cardinalitySeparator,
		app.cardinalityZeroPlus,
		app.cardinalityOnePlus,
	)

	if err != nil {
		return nil, nil, err
	}

	builder := app.cardinalityBuilder.Create().WithMin(retMin)
	if pRetMax != nil {
		builder.WithMax(*pRetMax)
	}

	retIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, retRemaining, nil
}

func (app *application) bytesToRule(input []byte) (rules.Rule, []byte, error) {
	name, value, remaining, err := bytesToRuleNameAndValue(
		input,
		app.ruleNameValueSeparator,
		app.possibleUpperCaseLetters,
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
