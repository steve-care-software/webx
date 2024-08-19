package grammars

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/nfts"
)

type application struct {
	grammarBuilder             grammars.Builder
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
}

func createApplication(
	grammarBuilder grammars.Builder,
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
) Application {
	out := application{
		grammarBuilder:             grammarBuilder,
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
	}

	return &out
}

// Parse parses the input and returns a grammar instance
func (app *application) Parse(input []byte) (grammars.Grammar, []byte, error) {
	return app.bytesToGrammar(input)
}

// Compile compiles a grammar to an NFT
func (app *application) Compile(grammar grammars.Grammar) (nfts.NFT, error) {
	//return app.grammarToAST(grammar)
	return nil, nil
}

// Decompile decompiles an NFT to a grammar instance
func (app *application) Decompile(ast nfts.NFT) (grammars.Grammar, error) {
	return nil, nil
}

// Compose composes a grammar instance to a grammar input
func (app *application) Compose(grammar grammars.Grammar) ([]byte, error) {
	return nil, nil
}

/*
func (app *application) grammarToAST(grammar grammars.Grammar) (asts.AST, error) {
	root := grammar.Root()
	if root.IsBlock() {
		blockName := root.Block()
		return app.fetchBlockThenConvertToAST(blockName, grammar)
	}

	ruleName := root.Rule()
	return app.fetchRuleThenConvertToAST(ruleName, grammar)
}

func (app *application) fetchBlockThenConvertToAST(blockName string, grammar grammars.Grammar) (asts.AST, error) {
	nft, list, err := app.fetchBlockThenConvertToNFT(blockName, grammar)
	if err != nil {
		return nil, err
	}

	return app.createAST(list, nft.Hash())
}

func (app *application) fetchRuleThenConvertToAST(ruleName string, grammar grammars.Grammar) (asts.AST, error) {
	nft, err := app.fetchRuleThenConvertToNFT(ruleName, grammar)
	if err != nil {
		return nil, err
	}

	return app.createAST([]asts.NFT{nft}, nft.Hash())
}

func (app *application) blocksToNFT(blocks blocks.Blocks) (asts.NFT, []asts.NFT, error) {
	return nil, nil, nil
}

func (app *application) blockToNFT(block blocks.Block) (asts.NFT, []asts.NFT, error) {
	return nil, nil, nil
}*/

/*
	func (app *application) linesToNFT(lines lines.Lines, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		list := lines.List()
		nftHashes := []hash.Hash{}
		nftList := []asts.NFT{}
		for _, oneLine := range list {
			retNFT, retList, err := app.lineToNFT(oneLine, grammar)
			if err != nil {
				return nil, nil, err
			}

			nftHashes = append(nftHashes, retNFT.Hash())
			nftList = append(nftList, retNFT)
			nftList = append(nftList, retList...)
		}

		nft, err := app.nftBuilder.Create().WithNFTs(nftHashes).Now()
		if err != nil {
			return nil, nil, err
		}

		return nft, nftList, nil
	}

	func (app *application) lineToNFT(line lines.Line, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		tokens := line.Tokens()
		tokensNFT, retList, err := app.tokensToNFT(tokens, grammar)
		if err != nil {
			return nil, nil, err
		}

		nftHashes := []hash.Hash{
			tokensNFT.Hash(),
		}

		nftList := retList
		if line.HasExecution() {
			execution := line.Execution()
			execNFT, retExecList, err := app.executionToNFT(execution, grammar)
			if err != nil {
				return nil, nil, err
			}

			nftHashes = append(nftHashes, execNFT.Hash())
			nftList = append(nftList, retExecList...)
		}

		if line.HasReplacement() {
			replacement := line.Replacement()
			elNFT, retElList, err := app.elementToNFT(replacement, grammar)
			if err != nil {
				return nil, nil, err
			}

			nftHashes = append(nftHashes, elNFT.Hash())
			nftList = append(nftList, retElList...)
		}

		nft, err := app.nftBuilder.Create().WithNFTs(nftHashes).Now()
		if err != nil {
			return nil, nil, err
		}

		return nft, nftList, nil
	}

	func (app *application) executionToNFT(execution executions.Execution, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		fnNameBytes := []byte(execution.FuncName())
		fnNameNFT, err := app.nftBuilder.Create().WithBytes(fnNameBytes).Now()
		if err != nil {
			return nil, nil, err
		}

		elements := execution.Elements()
		elementsNFT, retList, err := app.elementsToNFT(elements, grammar)
		if err != nil {
			return nil, nil, err
		}

		nft, err := app.nftBuilder.Create().WithNFTs([]hash.Hash{
			fnNameNFT.Hash(),
			elementsNFT.Hash(),
		}).Now()

		return nft, retList, nil
	}

	func (app *application) tokensToNFT(tokens tokens.Tokens, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		list := tokens.List()
		nftHashes := []hash.Hash{}
		nftList := []asts.NFT{}
		for _, oneToken := range list {
			retNFT, retList, err := app.tokenToNFT(oneToken, grammar)
			if err != nil {
				return nil, nil, err
			}

			nftHashes = append(nftHashes, retNFT.Hash())
			nftList = append(nftList, retNFT)
			nftList = append(nftList, retList...)
		}

		nft, err := app.nftBuilder.Create().WithNFTs(nftHashes).Now()
		if err != nil {
			return nil, nil, err
		}

		return nft, nftList, nil
	}

	func (app *application) cardinalityToNFT(cardinality cardinalities.Cardinality) (asts.NFT, error) {
		min := cardinality.Min()
		output := []byte{
			BytesCardinalityPrefix,
		}

		output = append(output, uintToBytes(uint64(min))...)
		if cardinality.HasMax() {
			output = append(output, BytesCardinalitySeparator)
			pMax := cardinality.Max()
			output = append(output, uintToBytes(uint64(*pMax))...)
		}

		output = append(output, BytesCardinalitySuffix)
		return app.nftBuilder.Create().
			WithBytes(output).
			Now()
	}

	func (app *application) tokenToNFT(token tokens.Token, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		element := token.Element()
		elementNFT, retList, err := app.elementToNFT(element, grammar)
		if err != nil {
			return nil, nil, err
		}

		cardinality := token.Cardinality()
		cardinalityNFT, err := app.cardinalityToNFT(cardinality)
		if err != nil {
			return nil, nil, err
		}

		nft, err := app.nftBuilder.Create().WithNFTs([]hash.Hash{
			elementNFT.Hash(),
			cardinalityNFT.Hash(),
		}).Now()

		if err != nil {
			return nil, nil, err
		}

		list := append(retList, elementNFT)
		list = append(list, cardinalityNFT)
		return nft, list, nil
	}

	func (app *application) elementsToNFT(elements elements.Elements, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		list := elements.List()
		nftHashes := []hash.Hash{}
		nftList := []asts.NFT{}
		for _, oneElement := range list {
			retNFT, retList, err := app.elementToNFT(oneElement, grammar)
			if err != nil {
				return nil, nil, err
			}

			nftHashes = append(nftHashes, retNFT.Hash())
			nftList = append(nftList, retNFT)
			nftList = append(nftList, retList...)
		}

		nft, err := app.nftBuilder.Create().WithNFTs(nftHashes).Now()
		if err != nil {
			return nil, nil, err
		}

		return nft, nftList, nil
	}

	func (app *application) elementToNFT(element elements.Element, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		if element.IsRule() {
			ruleName := element.Rule()
			nft, err := app.fetchRuleThenConvertToNFT(ruleName, grammar)
			if err != nil {
				return nil, nil, err
			}

			return nft, []asts.NFT{}, nil
		}

		blockName := element.Block()
		return app.fetchBlockThenConvertToNFT(blockName, grammar)
	}

	func (app *application) fetchBlockThenConvertToNFT(blockName string, grammar grammars.Grammar) (asts.NFT, []asts.NFT, error) {
		block, err := grammar.Blocks().Fetch(blockName)
		if err != nil {
			return nil, nil, err
		}

		return app.blockToNFT(block)
	}

	func (app *application) fetchRuleThenConvertToNFT(ruleName string, grammar grammars.Grammar) (asts.NFT, error) {
		rule, err := grammar.Rules().Fetch(ruleName)
		if err != nil {
			return nil, err
		}

		return app.ruleToNFT(rule)
	}

	func (app *application) createAST(list []asts.NFT, entry hash.Hash) (asts.AST, error) {
		library, err := app.nftListToNFTs(list)
		if err != nil {
			return nil, err
		}

		return app.astBuilder.Create().
			WithEntry(entry).
			WithLibrary(library).
			Now()
	}

	func (app *application) nftListToNFTs(list []asts.NFT) (asts.NFTs, error) {
		return app.nftsBuilder.Create().
			WithList(list).
			Now()
	}

	func (app *application) ruleToNFT(rule rules.Rule) (asts.NFT, error) {
		bytes := rule.Bytes()
		return app.nftBuilder.Create().
			WithBytes(bytes).
			Now()
	}
*/
func (app *application) bytesToGrammar(input []byte) (grammars.Grammar, []byte, error) {
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

	retRules, retRemaining, err := app.bytesToRules(retBlocksRemaining)
	if err != nil {
		return nil, nil, err
	}

	ins, err := builder.WithBlocks(retBlocks).
		WithRules(retRules).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
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

	return retIns, filterPrefix(remaining[1:], app.filterBytes), nil
}

func (app *application) bytesToSuites(input []byte) (suites.Suites, []byte, error) {
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

	return blockName, filterPrefix(retBlockRemaining[1:], app.filterBytes), nil
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

func (app *application) bytesToExecution(input []byte) (executions.Execution, []byte, error) {
	funcName, retRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleFuncNameCharacters, app.filterBytes)
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

func (app *application) bytesToRules(input []byte) (rules.Rules, []byte, error) {
	remaining := input
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

	return ins, remaining, nil
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
		app.filterBytes,
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
	blockName, retBlockRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.possibleLetters, app.filterBytes)
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
		app.filterBytes,
	)

	if err != nil {
		return "", nil, err
	}

	return string(retRuleName), retRemaining, nil
}
