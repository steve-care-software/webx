package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"
)

type application struct {
	tokensBuilder            tokens.Builder
	tokenBuilder             tokens.TokenBuilder
	elementBuilder           elements.Builder
	ruleBuilder              rules.RuleBuilder
	cardinalityBuilder       cardinalities.Builder
	ruleNameValueSeparator   byte
	possibleLetters          []byte
	possibleLowerCaseLetters []byte
	possibleUpperCaseLetters []byte
	possibleNumbers          []byte
	tokenReferenceSeparator  byte
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
	tokensBuilder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
	elementBuilder elements.Builder,
	ruleBuilder rules.RuleBuilder,
	cardinalityBuilder cardinalities.Builder,
	ruleNameValueSeparator byte,
	possibleLetters []byte,
	possibleLowerCaseLetters []byte,
	possibleUpperCaseLetters []byte,
	possibleNumbers []byte,
	tokenReferenceSeparator byte,
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
		tokensBuilder:            tokensBuilder,
		tokenBuilder:             tokenBuilder,
		elementBuilder:           elementBuilder,
		ruleBuilder:              ruleBuilder,
		cardinalityBuilder:       cardinalityBuilder,
		ruleNameValueSeparator:   ruleNameValueSeparator,
		possibleLetters:          possibleLetters,
		possibleLowerCaseLetters: possibleLowerCaseLetters,
		possibleUpperCaseLetters: possibleUpperCaseLetters,
		possibleNumbers:          possibleNumbers,
		tokenReferenceSeparator:  tokenReferenceSeparator,
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

func (app *application) bytesToToken(input []byte) (tokens.Token, []byte, error) {
	if len(input) <= 0 {
		return nil, nil, errors.New("the token was expected to contain at least 1 byte")
	}

	if input[0] != app.tokenReferenceSeparator {
		return nil, nil, errors.New("the token was expected to contain the tokenReference byte at its prefix")
	}

	// try to match a rule
	input = input[1:]
	elementBuilder := app.elementBuilder.Create()
	ruleName, retRemaining, err := app.bytesToRuleName(input)
	if err != nil {
		// there is no rule, so try to match a block
		blockName, retBlockRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleLetters)
		if err != nil {
			return nil, nil, err
		}

		elementBuilder.WithBlock(string(blockName))
		retRemaining = retBlockRemaining
	}

	if err == nil {
		elementBuilder.WithRule(ruleName)
	}

	element, err := elementBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	cardinalityIns, retRemainingAfterCardinality, err := app.bytesToCardinality(retRemaining)
	if err != nil {
		ins, err := app.cardinalityBuilder.Create().WithMin(1).WithMax(1).Now()
		if err != nil {
			return nil, nil, err
		}

		cardinalityIns = ins
	}

	if err == nil {
		retRemaining = retRemainingAfterCardinality
	}

	ins, err := app.tokenBuilder.Create().
		WithCardinality(cardinalityIns).
		WithElement(element).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
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

func (app *application) bytesToRuleName(input []byte) (string, []byte, error) {
	retRuleName, retRemaining, err := bytesToRuleName(
		input,
		app.possibleUpperCaseLetters,
		app.ruleNameSeparator,
	)

	if err != nil {
		return "", nil, err
	}

	return string(retRuleName), retRemaining, nil
}
