package grammars

import (
	"bytes"
	"errors"

	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"
)

type application struct {
	blocksBuilder              blocks.Builder
	blockBuilder               blocks.BlockBuilder
	suitesBuilder              suites.Builder
	suiteBuilder               suites.SuiteBuilder
	linesBuilder               lines.Builder
	lineBuilder                lines.LineBuilder
	executionBuilder           executions.Builder
	tokensBuilder              tokens.Builder
	tokenBuilder               tokens.TokenBuilder
	elementsBuilder            elements.Builder
	elementBuilder             elements.ElementBuilder
	ruleBuilder                rules.RuleBuilder
	cardinalityBuilder         cardinalities.Builder
	suiteSeparatorPrefix       []byte
	possibleLetters            []byte
	possibleLowerCaseLetters   []byte
	possibleUpperCaseLetters   []byte
	possibleNumbers            []byte
	possibleFuncNameCharacters []byte
	blockSuffix                byte
	suiteLineSuffix            byte
	failSeparator              byte
	blockDefinitionSeparator   byte
	linesSeparator             byte
	lineSeparator              byte
	tokenReferenceSeparator    byte
	ruleNameSeparator          byte
	ruleNameValueSeparator     byte
	ruleValuePrefix            byte
	ruleValueSuffix            byte
	ruleValueEscape            byte
	cardinalityOpen            byte
	cardinalityClose           byte
	cardinalitySeparator       byte
	cardinalityZeroPlus        byte
	cardinalityOnePlus         byte
}

func createApplication(
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	linesBuilder lines.Builder,
	lineBuilder lines.LineBuilder,
	executionBuilder executions.Builder,
	tokensBuilder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
	elementsBuilder elements.Builder,
	elementBuilder elements.ElementBuilder,
	ruleBuilder rules.RuleBuilder,
	cardinalityBuilder cardinalities.Builder,
	suiteSeparatorPrefix []byte,
	possibleLetters []byte,
	possibleLowerCaseLetters []byte,
	possibleUpperCaseLetters []byte,
	possibleNumbers []byte,
	possibleFuncNameCharacters []byte,
	blockSuffix byte,
	suiteLineSuffix byte,
	failSeparator byte,
	blockDefinitionSeparator byte,
	linesSeparator byte,
	lineSeparator byte,
	tokenReferenceSeparator byte,
	ruleNameSeparator byte,
	ruleNameValueSeparator byte,
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
		blocksBuilder:              blocksBuilder,
		blockBuilder:               blockBuilder,
		suitesBuilder:              suitesBuilder,
		suiteBuilder:               suiteBuilder,
		linesBuilder:               linesBuilder,
		lineBuilder:                lineBuilder,
		executionBuilder:           executionBuilder,
		tokensBuilder:              tokensBuilder,
		tokenBuilder:               tokenBuilder,
		elementsBuilder:            elementsBuilder,
		elementBuilder:             elementBuilder,
		ruleBuilder:                ruleBuilder,
		cardinalityBuilder:         cardinalityBuilder,
		suiteSeparatorPrefix:       suiteSeparatorPrefix,
		possibleLetters:            possibleLetters,
		possibleLowerCaseLetters:   possibleLowerCaseLetters,
		possibleUpperCaseLetters:   possibleUpperCaseLetters,
		possibleNumbers:            possibleNumbers,
		possibleFuncNameCharacters: possibleFuncNameCharacters,
		suiteLineSuffix:            suiteLineSuffix,
		failSeparator:              failSeparator,
		blockDefinitionSeparator:   blockDefinitionSeparator,
		blockSuffix:                blockSuffix,
		linesSeparator:             linesSeparator,
		lineSeparator:              lineSeparator,
		tokenReferenceSeparator:    tokenReferenceSeparator,
		ruleNameSeparator:          ruleNameSeparator,
		ruleNameValueSeparator:     ruleNameValueSeparator,
		ruleValuePrefix:            ruleValuePrefix,
		ruleValueSuffix:            ruleValueSuffix,
		ruleValueEscape:            ruleValueEscape,
		cardinalityOpen:            cardinalityOpen,
		cardinalityClose:           cardinalityClose,
		cardinalitySeparator:       cardinalitySeparator,
		cardinalityZeroPlus:        cardinalityZeroPlus,
		cardinalityOnePlus:         cardinalityOnePlus,
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

func (app *application) bytesToBlocks(input []byte) (blocks.Blocks, []byte, error) {
	remaining := input
	list := []blocks.Block{}
	for {
		retBlock, retRemaining, err := app.bytesToBlock(remaining)
		if err != nil {
			break
		}

		list = append(list, retBlock)
		remaining = retRemaining
	}

	ins, err := app.blocksBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) bytesToBlock(input []byte) (blocks.Block, []byte, error) {
	blockName, retBlockNameRemaining, err := app.bytesToBlockDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	retLines, retLinesRemaining, err := app.bytesToLines(retBlockNameRemaining)
	if err != nil {
		return nil, nil, err
	}

	remaining := retLinesRemaining
	builder := app.blockBuilder.Create().WithName(blockName).WithLines(retLines)
	retSuites, retSuitesRemaining, err := app.bytesToSuites(retLinesRemaining)
	if err == nil {
		builder.WithSuites(retSuites)
		remaining = retSuitesRemaining
	}

	if len(remaining) <= 0 {
		return nil, nil, errors.New("the block was expected to contain at least 1 byte at the end of its definition")
	}

	if remaining[0] != app.blockSuffix {
		return nil, nil, errors.New("the block was expected to contain the blockSuffix byte at its suffix")
	}

	retIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, remaining[1:], nil
}

func (app *application) bytesToSuites(input []byte) (suites.Suites, []byte, error) {
	if !bytes.HasPrefix(input, app.suiteSeparatorPrefix) {
		return nil, nil, errors.New("the suite was expecting the suite prefix bytes as its prefix")
	}

	remaining := input[len(app.suiteSeparatorPrefix):]
	list := []suites.Suite{}
	for {
		retSuite, retRemaining, err := app.bytesToSuite(remaining)
		if err != nil {
			break
		}

		list = append(list, retSuite)
		remaining = retRemaining
	}

	ins, err := app.suitesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) bytesToSuite(input []byte) (suites.Suite, []byte, error) {
	testName, retBlockNameRemaining, err := app.bytesToBlockDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	remaining := retBlockNameRemaining
	builder := app.suiteBuilder.Create().WithName(testName)
	if len(retBlockNameRemaining) != 0 && retBlockNameRemaining[0] == app.failSeparator {
		builder.IsFail()
		remaining = retBlockNameRemaining[1:]
	}

	element, retElementRemaining, err := app.bytesToElementReference(remaining)
	if err != nil {
		return nil, nil, err
	}

	retIns, err := builder.WithElement(element).Now()
	if err != nil {
		return nil, nil, err
	}

	if len(retElementRemaining) <= 0 {
		return nil, nil, errors.New("the suite was expected to contain at least 1 byte at the end of its instruction")
	}

	if retElementRemaining[0] != app.suiteLineSuffix {
		return nil, nil, errors.New("the suite was expected to contain the suiteLineSuffix byte at its suffix")
	}

	return retIns, retElementRemaining[1:], nil
}

func (app *application) bytesToBlockDefinition(input []byte) (string, []byte, error) {
	blockName, retBlockRemaining, err := app.bytesToBlockName(input)
	if err != nil {
		return "", nil, err
	}

	if len(retBlockRemaining) <= 0 {
		return "", nil, errors.New("the blockDefinition was expected to contain at least 1 byte after fetching its name")
	}

	if retBlockRemaining[0] != app.blockDefinitionSeparator {
		return "", nil, errors.New("the blockDefinition was expected to contain the blockDefinitionSeparator byte at its suffix")
	}

	return blockName, retBlockRemaining[1:], nil
}

func (app *application) bytesToLines(input []byte) (lines.Lines, []byte, error) {
	remaining := input
	list := []lines.Line{}
	cpt := 0
	for {

		isFirst := cpt <= 0
		if !isFirst && remaining[0] != app.linesSeparator {
			break
		}

		if !isFirst {
			remaining = remaining[1:]
		}

		retLine, retRemaining, err := app.bytesToLine(remaining)
		if err != nil {
			break
		}

		list = append(list, retLine)
		remaining = retRemaining
		cpt++
	}

	ins, err := app.linesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) bytesToLine(input []byte) (lines.Line, []byte, error) {
	retTokens, retRemaining, err := app.bytesToTokens(input)
	if err != nil {
		return nil, nil, err
	}

	builder := app.lineBuilder.Create().WithTokens(retTokens)
	for i := 0; i < 2; i++ {
		retExecution, retElement, retRemainingAfterExexOrToken, err := app.bytesToExecutionOrReplacement(retRemaining)
		if err != nil {
			break
		}

		if retExecution != nil {
			builder.WithExecution(retExecution)
		}

		if retElement != nil {
			builder.WithReplacement(retElement)
		}

		retRemaining = retRemainingAfterExexOrToken
	}

	line, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return line, retRemaining, nil
}

func (app *application) bytesToExecutionOrReplacement(input []byte) (executions.Execution, elements.Element, []byte, error) {
	if len(input) <= 0 {
		return nil, nil, nil, errors.New("the execution or replacement was expected to contain at least 1 byte for its separator")
	}

	if input[0] != app.lineSeparator {
		return nil, nil, nil, errors.New("the execution or replacement was expected to contain its separator")
	}

	retExecution, retRemaining, err := app.bytesToExecution(input[1:])
	if err != nil {
		retElement, retElementRemaining, err := app.bytesToElementReference(input[1:])
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, retElement, retElementRemaining, nil
	}

	return retExecution, nil, retRemaining, nil
}

func (app *application) bytesToExecution(input []byte) (executions.Execution, []byte, error) {
	funcName, retRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleFuncNameCharacters)
	if err != nil {
		return nil, nil, err
	}

	builder := app.executionBuilder.Create().WithFuncName(string(funcName))
	elements, retElementsRemaining, err := app.bytesToElementReferences(retRemaining)
	if err == nil {
		builder.WithElements(elements)
		retRemaining = retElementsRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

func (app *application) bytesToTokens(input []byte) (tokens.Tokens, []byte, error) {
	list, retRemaining, err := app.bytesToTokenList(input)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.tokensBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

func (app *application) bytesToTokenList(input []byte) ([]tokens.Token, []byte, error) {
	list := []tokens.Token{}
	remaining := input
	for {
		retToken, retRemaining, err := app.bytesToToken(remaining)
		if err != nil {
			break
		}

		list = append(list, retToken)
		remaining = retRemaining
	}

	return list, remaining, nil
}

func (app *application) bytesToToken(input []byte) (tokens.Token, []byte, error) {
	element, retRemaining, err := app.bytesToElementReference(input)
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

func (app *application) bytesToElementReferences(input []byte) (elements.Elements, []byte, error) {
	list := []elements.Element{}
	remaining := input
	for {
		retElement, retRemaining, err := app.bytesToElementReference(remaining)
		if err != nil {
			break
		}

		list = append(list, retElement)
		remaining = retRemaining
	}

	ins, err := app.elementsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) bytesToElementReference(input []byte) (elements.Element, []byte, error) {
	if len(input) <= 0 {
		return nil, nil, errors.New("the token was expected to contain at least 1 byte")
	}

	if input[0] != app.tokenReferenceSeparator {
		return nil, nil, errors.New("the token was expected to contain the tokenReference byte at its prefix")
	}

	return app.bytesToElement(input[1:])
}

func (app *application) bytesToElement(input []byte) (elements.Element, []byte, error) {
	// try to match a rule
	elementBuilder := app.elementBuilder.Create()
	ruleName, retRemaining, err := app.bytesToRuleName(input)
	if err != nil {
		// there is no rule, so try to match a block
		blockName, retBlockRemaining, err := app.bytesToBlockName(input)
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

	return element, retRemaining, nil
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

func (app *application) bytesToBlockName(input []byte) (string, []byte, error) {
	blockName, retBlockRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleLetters)
	if err != nil {
		return "", nil, err
	}

	return string(blockName), retBlockRemaining, nil
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
