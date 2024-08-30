package grammars

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values/references"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/processors"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/reverses"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
)

type parserAdapter struct {
	grammarBuilder                    Builder
	blocksBuilder                     blocks.Builder
	blockBuilder                      blocks.BlockBuilder
	suitesBuilder                     suites.Builder
	suiteBuilder                      suites.SuiteBuilder
	linesBuilder                      lines.Builder
	lineBuilder                       lines.LineBuilder
	processorBuilder                  processors.Builder
	executionBuilder                  executions.Builder
	parametersBuilder                 parameters.Builder
	parameterBuilder                  parameters.ParameterBuilder
	valueBuilder                      values.Builder
	referenceBuilder                  references.Builder
	tokensBuilder                     tokens.Builder
	tokenBuilder                      tokens.TokenBuilder
	reverseBuilder                    reverses.Builder
	elementsBuilder                   elements.Builder
	elementBuilder                    elements.ElementBuilder
	rulesBuilder                      rules.Builder
	ruleBuilder                       rules.RuleBuilder
	cardinalityBuilder                cardinalities.Builder
	filterBytes                       []byte
	suiteSeparatorPrefix              []byte
	blockNameAfterFirstByteCharacters []byte
	possibleLowerCaseLetters          []byte
	possibleUpperCaseLetters          []byte
	possibleNumbers                   []byte
	possibleFuncNameCharacters        []byte
	omissionPrefix                    byte
	omissionSuffix                    byte
	versionPrefix                     byte
	versionSuffix                     byte
	rootPrefix                        byte
	rootSuffix                        byte
	blockSuffix                       byte
	suiteLineSuffix                   byte
	failSeparator                     byte
	blockDefinitionSeparator          byte
	linesSeparator                    byte
	lineSeparator                     byte
	tokenReversePrefix                byte
	tokenReverseEscapePrefix          byte
	tokenReverseEscapeSuffix          byte
	tokenReferenceSeparator           byte
	ruleNameSeparator                 byte
	ruleNameValueSeparator            byte
	ruleValuePrefix                   byte
	ruleValueSuffix                   byte
	ruleValueEscape                   byte
	cardinalityOpen                   byte
	cardinalityClose                  byte
	cardinalitySeparator              byte
	cardinalityZeroPlus               byte
	cardinalityOnePlus                byte
	cardinalityOptional               byte
	indexOpen                         byte
	indexClose                        byte
	parameterSeparator                byte
	syscallDefinitionSeparator        byte
	sysCallNamePrefix                 byte
	sysCallFuncNamePrefix             byte
	sysCallPrefix                     byte
	sysCallSuffix                     byte
}

func createParserAdapter(
	grammarBuilder Builder,
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	linesBuilder lines.Builder,
	lineBuilder lines.LineBuilder,
	processorBuilder processors.Builder,
	executionBuilder executions.Builder,
	parametersBuilder parameters.Builder,
	parameterBuilder parameters.ParameterBuilder,
	valueBuilder values.Builder,
	referenceBuilder references.Builder,
	tokensBuilder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
	reverseBuilder reverses.Builder,
	elementsBuilder elements.Builder,
	elementBuilder elements.ElementBuilder,
	rulesBuilder rules.Builder,
	ruleBuilder rules.RuleBuilder,
	cardinalityBuilder cardinalities.Builder,
	filterBytes []byte,
	suiteSeparatorPrefix []byte,
	blockNameAfterFirstByteCharacters []byte,
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
	tokenReversePrefix byte,
	tokenReverseEscapePrefix byte,
	tokenReverseEscapeSuffix byte,
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
	cardinalityOptional byte,
	indexOpen byte,
	indexClose byte,
	parameterSeparator byte,
	syscallDefinitionSeparator byte,
	sysCallNamePrefix byte,
	sysCallFuncNamePrefix byte,
	sysCallPrefix byte,
	sysCallSuffix byte,
) ParserAdapter {
	out := parserAdapter{
		grammarBuilder:                    grammarBuilder,
		blocksBuilder:                     blocksBuilder,
		blockBuilder:                      blockBuilder,
		suitesBuilder:                     suitesBuilder,
		suiteBuilder:                      suiteBuilder,
		linesBuilder:                      linesBuilder,
		lineBuilder:                       lineBuilder,
		processorBuilder:                  processorBuilder,
		executionBuilder:                  executionBuilder,
		parametersBuilder:                 parametersBuilder,
		parameterBuilder:                  parameterBuilder,
		valueBuilder:                      valueBuilder,
		referenceBuilder:                  referenceBuilder,
		tokensBuilder:                     tokensBuilder,
		tokenBuilder:                      tokenBuilder,
		reverseBuilder:                    reverseBuilder,
		elementsBuilder:                   elementsBuilder,
		elementBuilder:                    elementBuilder,
		rulesBuilder:                      rulesBuilder,
		ruleBuilder:                       ruleBuilder,
		cardinalityBuilder:                cardinalityBuilder,
		filterBytes:                       filterBytes,
		suiteSeparatorPrefix:              suiteSeparatorPrefix,
		blockNameAfterFirstByteCharacters: blockNameAfterFirstByteCharacters,
		possibleLowerCaseLetters:          possibleLowerCaseLetters,
		possibleUpperCaseLetters:          possibleUpperCaseLetters,
		possibleNumbers:                   possibleNumbers,
		possibleFuncNameCharacters:        possibleFuncNameCharacters,
		omissionPrefix:                    omissionPrefix,
		omissionSuffix:                    omissionSuffix,
		versionPrefix:                     versionPrefix,
		versionSuffix:                     versionSuffix,
		rootPrefix:                        rootPrefix,
		rootSuffix:                        rootSuffix,
		suiteLineSuffix:                   suiteLineSuffix,
		failSeparator:                     failSeparator,
		blockDefinitionSeparator:          blockDefinitionSeparator,
		blockSuffix:                       blockSuffix,
		linesSeparator:                    linesSeparator,
		lineSeparator:                     lineSeparator,
		tokenReversePrefix:                tokenReversePrefix,
		tokenReverseEscapePrefix:          tokenReverseEscapePrefix,
		tokenReverseEscapeSuffix:          tokenReverseEscapeSuffix,
		tokenReferenceSeparator:           tokenReferenceSeparator,
		ruleNameSeparator:                 ruleNameSeparator,
		ruleNameValueSeparator:            ruleNameValueSeparator,
		ruleValuePrefix:                   ruleValuePrefix,
		ruleValueSuffix:                   ruleValueSuffix,
		ruleValueEscape:                   ruleValueEscape,
		cardinalityOpen:                   cardinalityOpen,
		cardinalityClose:                  cardinalityClose,
		cardinalitySeparator:              cardinalitySeparator,
		cardinalityZeroPlus:               cardinalityZeroPlus,
		cardinalityOnePlus:                cardinalityOnePlus,
		cardinalityOptional:               cardinalityOptional,
		indexOpen:                         indexOpen,
		indexClose:                        indexClose,
		parameterSeparator:                parameterSeparator,
		syscallDefinitionSeparator:        syscallDefinitionSeparator,
		sysCallNamePrefix:                 sysCallNamePrefix,
		sysCallFuncNamePrefix:             sysCallFuncNamePrefix,
		sysCallPrefix:                     sysCallPrefix,
		sysCallSuffix:                     sysCallSuffix,
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

	builder = builder.WithBlocks(retBlocks)
	retRules, retRemaining, err := app.bytesToRules(retBlocksRemaining)
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
	cpt := 0
	remaining := input
	list := []blocks.Block{}
	for {
		retBlock, retRemaining, err := app.bytesToBlock(remaining)
		if err != nil {
			log.Printf("there was an error while creating the block (idx: %d): %s", cpt, err.Error())
			break
		}

		list = append(list, retBlock)
		remaining = retRemaining
		cpt++
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
		str := fmt.Sprintf("the block was expected to contain the blockSuffix byte at its suffix, data: \n%s\n", string(remaining))
		return nil, nil, errors.New(str)
	}

	retIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, filterPrefix(remaining[1:], app.filterBytes), nil
}

func (app *parserAdapter) bytesToParameterOrToken(input []byte) (parameters.Parameter, tokens.Token, []byte, error) {
	retParameter, retRemaining, err := app.bytesToParameter(input)
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

	retSuiteValue, retRemainingAfterBetween, err := extractBetween(remaining, app.ruleValuePrefix, app.ruleValueSuffix, &app.ruleValueEscape)
	if err != nil {
		return nil, nil, err
	}

	retIns, err := builder.WithValue(retSuiteValue).Now()
	if err != nil {
		return nil, nil, err
	}

	if len(retRemainingAfterBetween) <= 0 {
		return nil, nil, errors.New("the suite was expected to contain at least 1 byte at the end of its instruction")
	}

	if retRemainingAfterBetween[0] != app.suiteLineSuffix {
		return nil, nil, errors.New("the suite was expected to contain the suiteLineSuffix byte at its suffix")
	}

	return retIns, retRemainingAfterBetween[1:], nil
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
	remaining := input
	builder := app.lineBuilder.Create()
	retSyscall, retRemainingAfterSyscall, err := app.bytesToSyscall(remaining)
	if err == nil {
		builder.WithSyscall(retSyscall)
		remaining = retRemainingAfterSyscall
	}

	retTokens, retRemaining, err := app.bytesToTokens(remaining)
	if err != nil {
		return nil, nil, err
	}

	remaining = retRemaining
	builder.WithTokens(retTokens)
	retProcessor, retRemainingAfterProcessor, err := app.bytesToProcessor(remaining)
	if err == nil {
		builder.WithProcessor(retProcessor)
		remaining = retRemainingAfterProcessor
	}

	line, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return line, remaining, nil
}

func (app *parserAdapter) bytesToSyscall(input []byte) (executions.Execution, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if len(input) <= 0 {
		return nil, nil, errors.New("the syscall was expected to contain at least 1 byte for its prefix")
	}

	if input[0] != app.sysCallPrefix {
		return nil, nil, errors.New("the syscall was expected to contain the sysCallPrefix as its first byte")
	}

	remaining := input[1:]
	retExecution, retRemaining, err := app.bytesToExecution(remaining)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) <= 0 {
		return nil, nil, errors.New("the syscall was expected to contain at least 1 byte for its prefix")
	}

	if retRemaining[0] != app.sysCallSuffix {
		return nil, nil, errors.New("the syscall was expected to contain the sysCallSuffix as its last byte")
	}

	return retExecution, retRemaining[1:], nil
}

func (app *parserAdapter) bytesToProcessor(input []byte) (processors.Processor, []byte, error) {
	remaining := input
	builder := app.processorBuilder.Create()
	retExecution, retElement, retRemainingAfterExexOrToken, err := app.bytesToExecutionOrReplacement(remaining)
	if err == nil {
		remaining = retRemainingAfterExexOrToken
	}

	if retExecution != nil {
		builder.WithExecution(retExecution)
	}

	if retElement != nil {
		builder.WithReplacement(retElement)
	}

	processor, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return processor, remaining, nil
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
	parameters, retParametersRemaining, err := app.bytesToParameters(retRemaining)
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

func (app *parserAdapter) bytesToParameters(input []byte) (parameters.Parameters, []byte, error) {
	list := []parameters.Parameter{}
	remaining := input
	for {
		retParameter, retRemaining, err := app.bytesToParameter(remaining)
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

func (app *parserAdapter) bytesToValue(input []byte) (values.Value, []byte, error) {
	builder := app.valueBuilder.Create()
	retReference, retRemaining, err := app.bytesToReference(input)
	if err != nil {
		retValue, retRemainingAfterValue, err := extractBetween(input, app.ruleValuePrefix, app.ruleValueSuffix, &app.ruleValueEscape)
		if err != nil {
			return nil, nil, err
		}

		builder.WithBytes(retValue)
		retRemaining = retRemainingAfterValue
	}

	if retReference != nil {
		builder.WithReference(retReference)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

func (app *parserAdapter) bytesToReference(input []byte) (references.Reference, []byte, error) {
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

	builder := app.referenceBuilder.Create().WithIndex(retValue).WithElement(element)
	if err != nil {
		builder.WithIndex(0)
		retValueRemaining = retElementRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retValueRemaining, nil
}

func (app *parserAdapter) bytesToParameter(input []byte) (parameters.Parameter, []byte, error) {
	value, retRemainingAfterValue, err := app.bytesToValue(input)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemainingAfterValue) <= 0 {
		return nil, nil, errors.New("the parameter was expected to contain at least 1 byte before its name")
	}

	if retRemainingAfterValue[0] != app.parameterSeparator {
		return nil, nil, errors.New("the parameter was expected to contain the parameterSeparator byte before its name")
	}

	name, retNameRemaining, err := app.bytesToBlockName(retRemainingAfterValue[1:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.parameterBuilder.Create().
		WithValue(value).
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

	ins, err := app.tokensBuilder.Create().
		WithList(list).
		Now()

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
	remaining := filterPrefix(input, app.filterBytes)
	builder := app.tokenBuilder.Create()
	retReverse, retRemainingAfterReverse, err := app.bytesToTokenReverse(remaining)
	if err == nil {
		builder.WithReverse(retReverse)
		remaining = retRemainingAfterReverse
	}

	element, retRemaining, err := app.bytesToElementReference(remaining)
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

	ins, err := builder.
		WithCardinality(cardinalityIns).
		WithElement(element).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}

func (app *parserAdapter) bytesToTokenReverse(input []byte) (reverses.Reverse, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the tokenReverse was expected to contain at least 1 byte")
	}

	if remaining[0] != app.tokenReversePrefix {
		return nil, nil, errors.New("the tokenReverse was expected to contain the tokenReversePrefix byte at its prefix")
	}

	remaining = remaining[1:]
	builder := app.reverseBuilder.Create()
	retEscape, retRemaining, err := app.bytesToTokenReverseEscape(remaining)
	if err == nil {
		remaining = retRemaining
		builder.WithEscape(retEscape)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) bytesToTokenReverseEscape(input []byte) (elements.Element, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain at least 1 byte at its prefix")
	}

	if remaining[0] != app.tokenReverseEscapePrefix {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain the tokenReverseEscapePrefix byte at its prefix")
	}

	remaining = remaining[1:]
	retElement, retRemaining, err := app.bytesToElementReference(remaining)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) <= 0 {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain at least 1 byte at its suffix")
	}

	if retRemaining[0] != app.tokenReverseEscapeSuffix {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain the tokenReverseEscapeSuffix byte at its suffix")
	}

	return retElement, retRemaining[1:], nil
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
			return nil, nil, err
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
		app.cardinalityOptional,
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
	blockName, retBlockRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.blockNameAfterFirstByteCharacters, app.filterBytes)
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
