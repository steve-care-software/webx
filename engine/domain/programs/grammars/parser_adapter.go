package grammars

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls/values"
)

type parserAdapter struct {
	grammarBuilder             Builder
	syscallsBuilder            syscalls.Builder
	syscallBuilder             syscalls.SyscallBuilder
	syscallValuesBuilder       values.Builder
	syscallValueBuilder        values.ValueBuilder
	blocksBuilder              blocks.Builder
	blockBuilder               blocks.BlockBuilder
	suitesBuilder              suites.Builder
	suiteBuilder               suites.SuiteBuilder
	linesBuilder               lines.Builder
	lineBuilder                lines.LineBuilder
	executionBuilder           executions.Builder
	parametersBuilder          parameters.Builder
	parameterBuilder           parameters.ParameterBuilder
	tokensBuilder              tokens.Builder
	tokenBuilder               tokens.TokenBuilder
	elementsBuilder            elements.Builder
	elementBuilder             elements.ElementBuilder
	rulesBuilder               rules.Builder
	ruleBuilder                rules.RuleBuilder
	cardinalityBuilder         cardinalities.Builder
	filterBytes                []byte
	suiteSeparatorPrefix       []byte
	possibleLetters            []byte
	possibleLowerCaseLetters   []byte
	possibleUpperCaseLetters   []byte
	possibleNumbers            []byte
	possibleFuncNameCharacters []byte
	omissionPrefix             byte
	omissionSuffix             byte
	versionPrefix              byte
	versionSuffix              byte
	rootPrefix                 byte
	rootSuffix                 byte
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
	indexOpen                  byte
	indexClose                 byte
	parameterSeparator         byte
	sysCallNamePrefix          byte
	sysCallFuncNamePrefix      byte
}

func createParserAdapter(
	grammarBuilder Builder,
	syscallsBuilder syscalls.Builder,
	syscallBuilder syscalls.SyscallBuilder,
	syscallValuesBuilder values.Builder,
	syscallValueBuilder values.ValueBuilder,
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	linesBuilder lines.Builder,
	lineBuilder lines.LineBuilder,
	executionBuilder executions.Builder,
	parametersBuilder parameters.Builder,
	parameterBuilder parameters.ParameterBuilder,
	tokensBuilder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
	elementsBuilder elements.Builder,
	elementBuilder elements.ElementBuilder,
	rulesBuilder rules.Builder,
	ruleBuilder rules.RuleBuilder,
	cardinalityBuilder cardinalities.Builder,
	filterBytes []byte,
	suiteSeparatorPrefix []byte,
	possibleLetters []byte,
	possibleLowerCaseLetters []byte,
	possibleUpperCaseLetters []byte,
	possibleNumbers []byte,
	possibleFuncNameCharacters []byte,
	omissionPrefix byte,
	omissionSuffix byte,
	versionPrefix byte,
	versionSuffix byte,
	rootPrefix byte,
	rootSuffix byte,
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
	indexOpen byte,
	indexClose byte,
	parameterSeparator byte,
	sysCallNamePrefix byte,
	sysCallFuncNamePrefix byte,
) ParserAdapter {
	out := parserAdapter{
		grammarBuilder:             grammarBuilder,
		syscallsBuilder:            syscallsBuilder,
		syscallBuilder:             syscallBuilder,
		syscallValuesBuilder:       syscallValuesBuilder,
		syscallValueBuilder:        syscallValueBuilder,
		blocksBuilder:              blocksBuilder,
		blockBuilder:               blockBuilder,
		suitesBuilder:              suitesBuilder,
		suiteBuilder:               suiteBuilder,
		linesBuilder:               linesBuilder,
		lineBuilder:                lineBuilder,
		executionBuilder:           executionBuilder,
		parametersBuilder:          parametersBuilder,
		parameterBuilder:           parameterBuilder,
		tokensBuilder:              tokensBuilder,
		tokenBuilder:               tokenBuilder,
		elementsBuilder:            elementsBuilder,
		elementBuilder:             elementBuilder,
		rulesBuilder:               rulesBuilder,
		ruleBuilder:                ruleBuilder,
		cardinalityBuilder:         cardinalityBuilder,
		filterBytes:                filterBytes,
		suiteSeparatorPrefix:       suiteSeparatorPrefix,
		possibleLetters:            possibleLetters,
		possibleLowerCaseLetters:   possibleLowerCaseLetters,
		possibleUpperCaseLetters:   possibleUpperCaseLetters,
		possibleNumbers:            possibleNumbers,
		possibleFuncNameCharacters: possibleFuncNameCharacters,
		omissionPrefix:             omissionPrefix,
		omissionSuffix:             omissionSuffix,
		versionPrefix:              versionPrefix,
		versionSuffix:              versionSuffix,
		rootPrefix:                 rootPrefix,
		rootSuffix:                 rootSuffix,
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
		indexOpen:                  indexOpen,
		indexClose:                 indexClose,
		parameterSeparator:         parameterSeparator,
		sysCallNamePrefix:          sysCallNamePrefix,
		sysCallFuncNamePrefix:      sysCallFuncNamePrefix,
	}

	return &out
}

// ToGrammar takes the input and converts it to a grammar instance and the remaining data
func (app *parserAdapter) ToGrammar(input []byte) (Grammar, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	retVersion, retVersionRemaining, err := extractBetween(input, app.versionPrefix, app.versionSuffix, nil)
	if err != nil {
		return nil, nil, err
	}

	version, err := strconv.Atoi(string(retVersion))
	if err != nil {
		return nil, nil, err
	}

	retVersionRemaining = filterPrefix(retVersionRemaining, app.filterBytes)
	retRootBytes, retRootRemaining, err := extractBetween(retVersionRemaining, app.rootPrefix, app.rootSuffix, nil)
	if err != nil {
		return nil, nil, err
	}

	retRoot, _, err := app.bytesToElementReference(retRootBytes)
	if err != nil {
		return nil, nil, err
	}

	retRootRemaining = filterPrefix(retRootRemaining, app.filterBytes)
	remaining := retRootRemaining
	builder := app.grammarBuilder.Create().WithVersion(uint(version)).WithRoot(retRoot)
	retOmissionBytes, retOmissionRemaining, err := extractBetween(retRootRemaining, app.omissionPrefix, app.omissionSuffix, nil)
	if err == nil {
		retOmissions, _, err := app.bytesToElementReferences(retOmissionBytes)
		if err != nil {
			return nil, nil, err
		}

		builder.WithOmissions(retOmissions)
		remaining = retOmissionRemaining
	}

	retBlocks, retBlocksRemaining, err := app.bytesToBlocks(remaining)
	if err != nil {
		return nil, nil, err
	}

	retRemaining := retBlocksRemaining
	builder = builder.WithBlocks(retBlocks)
	retSyscalls, retSyscallsRemaining, err := app.bytesToSyscalls(retBlocksRemaining)
	if err == nil {
		builder.WithSyscalls(retSyscalls)
		retRemaining = retSyscallsRemaining
	}

	retRules, retRemaining, err := app.bytesToRules(retRemaining)
	if err != nil {
		return nil, nil, err
	}

	ins, err := builder.
		WithRules(retRules).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

// ToBytes takes a grammar and returns the bytes
func (app *parserAdapter) ToBytes(grammar Grammar) ([]byte, error) {
	return nil, nil
}

func (app *parserAdapter) bytesToBlocks(input []byte) (blocks.Blocks, []byte, error) {
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

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToBlock(input []byte) (blocks.Block, []byte, error) {
	blockName, retBlockNameRemaining, err := app.bytesToBlockDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	retLines, retLinesRemaining, err := app.bytesToLines(retBlockNameRemaining)
	if err != nil {
		return nil, nil, err
	}

	remaining := retLinesRemaining
	builder := app.blockBuilder.Create().WithName(blockName)
	linesList := retLines.List()
	listLength := len(linesList)
	if listLength == 1 {
		builder.WithLine(linesList[0])
	}

	if listLength > 1 {
		builder.WithLines(retLines)
	}

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

	return retIns, filterPrefix(remaining[1:], app.filterBytes), nil
}

func (app *parserAdapter) bytesToSyscalls(input []byte) (syscalls.Syscalls, []byte, error) {
	remaining := input
	list := []syscalls.Syscall{}
	for {
		retSyscall, retRemaining, err := app.bytesToSyscall(remaining)
		if err != nil {
			break
		}

		list = append(list, retSyscall)
		remaining = retRemaining
	}

	ins, err := app.syscallsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToSyscall(input []byte) (syscalls.Syscall, []byte, error) {
	blockName, retBlockNameRemaining, err := app.bytesToSyscallDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	if retBlockNameRemaining[0] != app.sysCallFuncNamePrefix {
		return nil, nil, errors.New("the sysCallFunc was expecting the sysCallFuncNamePrefix bytes as its prefix")
	}

	funcName, retFuncNameRemaining, err := app.bytesToFuncName(retBlockNameRemaining[1:])
	if err != nil {
		return nil, nil, err
	}

	builder := app.syscallBuilder.Create().
		WithFuncName(funcName).
		WithName(blockName)

	retRemaining := retFuncNameRemaining
	retValues, retValuesRemaining, err := app.bytesToValues(retFuncNameRemaining)
	if err == nil {
		builder.WithValues(retValues)
		retRemaining = retValuesRemaining
	}

	if len(retRemaining) <= 0 {
		return nil, nil, errors.New("the sysCall was expected to contain at least 1 byte at the end of its definition")
	}

	if retRemaining[0] != app.blockSuffix {
		return nil, nil, errors.New("the sysCall was expected to contain the blockSuffix byte at its suffix")
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(retRemaining[1:], app.filterBytes), nil
}

func (app *parserAdapter) bytesToSyscallDefinition(input []byte) (string, []byte, error) {
	if input[0] != app.sysCallNamePrefix {
		return "", nil, errors.New("the sysCallDefinition was expecting the sysCallNamePrefix bytes as its prefix")
	}

	remaining := filterPrefix(input[1:], app.filterBytes)
	blockName, retBlockNameRemaining, err := app.bytesToBlockDefinition(remaining)
	if err != nil {
		return "", nil, err
	}

	return string(blockName), filterPrefix(retBlockNameRemaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToSyscallName(input []byte) (string, []byte, error) {
	if input[0] != app.sysCallNamePrefix {
		return "", nil, errors.New("the sysCallDefinition was expecting the sysCallNamePrefix bytes as its prefix")
	}

	remaining := filterPrefix(input[1:], app.filterBytes)
	blockName, retBlockNameRemaining, err := app.bytesToBlockName(remaining)
	if err != nil {
		return "", nil, err
	}

	return string(blockName), filterPrefix(retBlockNameRemaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToValues(input []byte) (values.Values, []byte, error) {
	remaining := input
	list := []values.Value{}
	for {
		retValue, retRemaining, err := app.bytesToValue(remaining)
		if err != nil {
			break
		}

		list = append(list, retValue)
		remaining = retRemaining
	}

	ins, err := app.syscallValuesBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToValue(input []byte) (values.Value, []byte, error) {
	retParameter, retToken, retRemaining, err := app.bytesToParametereOrToken(input)
	if err != nil {
		return nil, nil, err
	}

	builder := app.syscallValueBuilder.Create()
	if retParameter != nil {
		builder.WithParameter(retParameter)
	}

	if retToken != nil {
		builder.WithToken(retToken)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

func (app *parserAdapter) bytesToParametereOrToken(input []byte) (parameters.Parameter, tokens.Token, []byte, error) {
	retParameter, retRemaining, err := app.bytesToParametere(input)
	if err == nil {
		return retParameter, nil, retRemaining, err
	}

	retToken, retRemaining, err := app.bytesToToken(input)
	if err != nil {
		return nil, nil, nil, err
	}

	return nil, retToken, retRemaining, nil
}

func (app *parserAdapter) bytesToSuites(input []byte) (suites.Suites, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if !bytes.HasPrefix(input, app.suiteSeparatorPrefix) {
		return nil, nil, errors.New("the suite was expecting the suite prefix bytes as its prefix")
	}

	remaining := filterPrefix(input[len(app.suiteSeparatorPrefix):], app.filterBytes)
	list := []suites.Suite{}
	for {
		retSuite, retRemaining, err := app.bytesToSuite(remaining)
		if err != nil {
			break
		}

		list = append(list, retSuite)
		remaining = filterPrefix(retRemaining, app.filterBytes)
	}

	ins, err := app.suitesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToSuite(input []byte) (suites.Suite, []byte, error) {
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

func (app *parserAdapter) bytesToBlockDefinition(input []byte) (string, []byte, error) {
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

	return blockName, filterPrefix(retBlockRemaining[1:], app.filterBytes), nil
}

func (app *parserAdapter) bytesToLines(input []byte) (lines.Lines, []byte, error) {
	remaining := input
	list := []lines.Line{}
	cpt := 0
	for {

		isFirst := cpt <= 0
		if !isFirst && remaining[0] != app.linesSeparator {
			break
		}

		if !isFirst {
			remaining = filterPrefix(remaining[1:], app.filterBytes)
		}

		retLine, retRemaining, err := app.bytesToLine(remaining)
		if err != nil {
			break
		}

		list = append(list, retLine)
		remaining = filterPrefix(retRemaining, app.filterBytes)
		cpt++
	}

	ins, err := app.linesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) bytesToLine(input []byte) (lines.Line, []byte, error) {
	retTokens, retRemaining, err := app.bytesToTokens(input)
	if err != nil {
		return nil, nil, err
	}

	remaining := retRemaining
	builder := app.lineBuilder.Create().WithTokens(retTokens)
	retExecution, retElement, retRemainingAfterExexOrToken, err := app.bytesToExecutionOrReplacement(retRemaining)
	if err == nil {
		remaining = retRemainingAfterExexOrToken
	}

	if retExecution != nil {
		builder.WithExecution(retExecution)
	}

	if retElement != nil {
		builder.WithReplacement(retElement)
	}

	line, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return line, remaining, nil
}

func (app *parserAdapter) bytesToExecutionOrReplacement(input []byte) (executions.Execution, elements.Element, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
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

func (app *parserAdapter) bytesToExecution(input []byte) (executions.Execution, []byte, error) {
	funcName, retRemaining, err := app.bytesToFuncName(input)
	if err != nil {
		return nil, nil, err
	}

	builder := app.executionBuilder.Create().WithFuncName(string(funcName))
	parameters, retParametersRemaining, err := app.bytesToParameteres(retRemaining)
	if err == nil {
		builder.WithParameters(parameters)
		retRemaining = retParametersRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

func (app *parserAdapter) bytesToFuncName(input []byte) (string, []byte, error) {
	funcName, retRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleFuncNameCharacters, app.filterBytes)
	if err != nil {
		return "", nil, err
	}

	return string(funcName), retRemaining, nil
}

func (app *parserAdapter) bytesToParameteres(input []byte) (parameters.Parameters, []byte, error) {
	list := []parameters.Parameter{}
	remaining := input
	for {
		retParameter, retRemaining, err := app.bytesToParametere(remaining)
		if err != nil {
			break
		}

		list = append(list, retParameter)
		remaining = retRemaining
	}

	ins, err := app.parametersBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) bytesToParametere(input []byte) (parameters.Parameter, []byte, error) {
	element, retElementRemaining, err := app.bytesToElementReference(input)
	if err != nil {
		return nil, nil, err
	}

	retValue, retValueRemaining, err := bytesToBracketsIndex(
		retElementRemaining,
		app.possibleNumbers,
		app.indexOpen,
		app.indexClose,
		app.filterBytes,
	)

	if err != nil {
		retValue = 0
		retValueRemaining = retElementRemaining
	}

	if len(retValueRemaining) <= 0 {
		return nil, nil, errors.New("the parameter was expected to contain at least 1 byte before its name")
	}

	if retValueRemaining[0] != app.parameterSeparator {
		return nil, nil, errors.New("the parameter was expected to contain the parameterSeparator byte before its name")
	}

	name, retNameRemaining, err := app.bytesToBlockName(retValueRemaining[1:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.parameterBuilder.Create().
		WithElement(element).
		WithIndex(uint(retValue)).
		WithName(name).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retNameRemaining, nil
}

func (app *parserAdapter) bytesToTokens(input []byte) (tokens.Tokens, []byte, error) {
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

func (app *parserAdapter) bytesToTokenList(input []byte) ([]tokens.Token, []byte, error) {
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

func (app *parserAdapter) bytesToToken(input []byte) (tokens.Token, []byte, error) {
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

func (app *parserAdapter) bytesToElementReferences(input []byte) (elements.Elements, []byte, error) {
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

func (app *parserAdapter) bytesToElementReference(input []byte) (elements.Element, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if len(input) <= 0 {
		return nil, nil, errors.New("the token was expected to contain at least 1 byte")
	}

	if input[0] != app.tokenReferenceSeparator {
		return nil, nil, errors.New("the token was expected to contain the tokenReference byte at its prefix")
	}

	input = filterPrefix(input[1:], app.filterBytes)
	return app.bytesToElement(input)
}

func (app *parserAdapter) bytesToElement(input []byte) (elements.Element, []byte, error) {
	// try to match a rule
	elementBuilder := app.elementBuilder.Create()
	ruleName, retRemaining, err := app.bytesToRuleName(input)
	if err != nil {
		// there is no rule, so try to match a block
		blockName, retBlockRemaining, err := app.bytesToBlockName(input)
		if err != nil {
			// there is no block, so try to match a syscall:
			sysCallName, retSysCallRemaining, err := app.bytesToSyscallName(input)
			if err != nil {
				return nil, nil, err
			}

			elementBuilder.WithSyscall(string(sysCallName))
			retRemaining = retSysCallRemaining
		}

		if err == nil {
			elementBuilder.WithBlock(string(blockName))
			retRemaining = retBlockRemaining
		}
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

func (app *parserAdapter) bytesToCardinality(input []byte) (cardinalities.Cardinality, []byte, error) {
	retMin, pRetMax, retRemaining, err := bytesToMinMax(
		input,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.cardinalitySeparator,
		app.cardinalityZeroPlus,
		app.cardinalityOnePlus,
		app.filterBytes,
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

func (app *parserAdapter) bytesToRules(input []byte) (rules.Rules, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	list := []rules.Rule{}
	for {
		retRule, retRemaining, err := app.bytesToRule(remaining)
		if err != nil {
			break
		}

		list = append(list, retRule)
		remaining = filterPrefix(retRemaining, app.filterBytes)
	}

	ins, err := app.rulesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *parserAdapter) bytesToRule(input []byte) (rules.Rule, []byte, error) {
	name, value, remaining, err := bytesToRuleNameAndValue(
		input,
		app.ruleNameValueSeparator,
		app.possibleUpperCaseLetters,
		app.ruleNameSeparator,
		app.ruleValuePrefix,
		app.ruleValueSuffix,
		app.ruleValueEscape,
		app.filterBytes,
	)

	if err != nil {
		return nil, nil, err
	}

	if len(remaining) <= 0 {
		return nil, nil, errors.New("the rule was expected to contain at least 1 byte at the end of its definition")
	}

	if remaining[0] != app.blockSuffix {
		return nil, nil, errors.New("the rule was expected to contain the blockSuffix byte at its suffix")
	}

	ins, err := app.ruleBuilder.Create().
		WithName(string(name)).
		WithBytes(value).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining[1:], app.filterBytes), nil
}

func (app *parserAdapter) bytesToBlockName(input []byte) (string, []byte, error) {
	blockName, retBlockRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleLetters, app.filterBytes)
	if err != nil {
		return "", nil, err
	}

	return string(blockName), retBlockRemaining, nil
}

func (app *parserAdapter) bytesToRuleName(input []byte) (string, []byte, error) {
	retRuleName, retRemaining, err := bytesToRuleName(
		input,
		app.possibleUpperCaseLetters,
		app.ruleNameSeparator,
		app.filterBytes,
	)

	if err != nil {
		return "", nil, err
	}

	return string(retRuleName), retRemaining, nil
}
