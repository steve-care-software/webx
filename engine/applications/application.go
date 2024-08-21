package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities/uints"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

type application struct {
	grammarParserAdapter       grammars.ParserAdapter
	grammarNFTAdapter          grammars.NFTAdapter
	ruleAdapter                rules.Adapter
	cardinalityAdapter         cardinalities.Adapter
	uintAdapter                uints.Adapter
	nftsBuilder                nfts.Builder
	nftBuilder                 nfts.NFTBuilder
	grammarBuilder             grammars.Builder
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
	funcsMap                   map[string]CoreFn
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
}

func createApplication(
	grammarParserAdapter grammars.ParserAdapter,
	grammarNFTAdapter grammars.NFTAdapter,
	ruleAdapter rules.Adapter,
	cardinalityAdapter cardinalities.Adapter,
	uintAdapter uints.Adapter,
	nftsBuilder nfts.Builder,
	nftBuilder nfts.NFTBuilder,
	grammarBuilder grammars.Builder,
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
	funcsMap map[string]CoreFn,
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
) Application {
	out := application{
		grammarParserAdapter:       grammarParserAdapter,
		grammarNFTAdapter:          grammarNFTAdapter,
		ruleAdapter:                ruleAdapter,
		cardinalityAdapter:         cardinalityAdapter,
		uintAdapter:                uintAdapter,
		nftsBuilder:                nftsBuilder,
		nftBuilder:                 nftBuilder,
		grammarBuilder:             grammarBuilder,
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
		funcsMap:                   funcsMap,
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
	}

	return &out
}

// ParseGrammar parses an input and creates a Grammar instance
func (app *application) ParseGrammar(input []byte) (grammars.Grammar, []byte, error) {
	return app.grammarParserAdapter.ToGrammar(input)
}

// CompileGrammar compiles a grammar to an NFT
func (app *application) CompileGrammar(grammar grammars.Grammar) (nfts.NFT, error) {
	return app.grammarNFTAdapter.ToNFT(grammar)
}

// DecompileGrammar decompiles an NFT into a grammar instance
func (app *application) DecompileGrammar(ast nfts.NFT) (grammars.Grammar, error) {
	return nil, nil
}

// ComposeBlock fetches a blockName from the grammar and composes an output
func (app *application) ComposeBlock(grammar grammars.Grammar, blockName string) ([]byte, error) {
	block, err := grammar.Blocks().Fetch(blockName)
	if err != nil {
		return nil, err
	}

	return app.writeBlock(grammar, block)
}

// ParseProgram takes a grammar and an input, parses it and returns the program
func (app *application) ParseProgram(grammar grammars.Grammar, input []byte) (programs.Program, error) {
	return nil, nil
}

// CompileProgram compiles a program to an NFT
func (app *application) CompileProgram(program programs.Program) (nfts.NFT, error) {
	return nil, nil
}

// DecompileProgram decompiles an NFT into a program instance
func (app *application) DecompileProgram(nft nfts.NFT) (programs.Program, error) {
	return nil, nil
}

// ComposeProgram takes the program and composes an output
func (app *application) ComposeProgram(program programs.Program) ([]byte, error) {
	return nil, nil
}

// Interpret interprets the input and returns the stack
func (app *application) Interpret(program programs.Program) (stacks.Stack, error) {
	return nil, nil
}

// Suites executes all the test suites of the grammar
func (app *application) Suites(grammar grammars.Grammar) ([]byte, error) {
	return nil, nil
}

// Suite executes the test suite of the provided blockName in the grammar
func (app *application) Suite(grammar grammars.Grammar, blockName string) ([]byte, error) {
	return nil, nil
}

func (app *application) writeBlock(grammar grammars.Grammar, block blocks.Block) ([]byte, error) {
	if !block.HasLine() {
		str := fmt.Sprintf("the block (name: %s) cannot be written because it contains lines instead of a line", block.Name())
		return nil, errors.New(str)
	}

	line := block.Line()
	return app.writeLine(grammar, line)
}

func (app *application) writeLine(grammar grammars.Grammar, line lines.Line) ([]byte, error) {
	tokens := line.Tokens()
	retBytes, tokensMap, err := app.writeTokens(grammar, tokens)
	if err != nil {
		return nil, err
	}

	output := retBytes
	if line.HasExecution() {
		execution := line.Execution()
		fnName := execution.FuncName()
		params := map[string][]byte{}
		if execution.HasParameters() {
			parametersList := execution.Parameters().List()
			for _, oneParameter := range parametersList {
				paramElementName := oneParameter.Element().Name()
				paramElementIndex := oneParameter.Index()
				if _, ok := tokensMap[paramElementName]; !ok {
					str := fmt.Sprintf("the func (name: %s) contains a param (name: %s, index: %d) that is not declared in the line", fnName, paramElementName, paramElementIndex)
					return nil, errors.New(str)
				}

				name := oneParameter.Name()
				params[name] = tokensMap[paramElementName][int(paramElementIndex)]
			}
		}

		if fn, ok := app.funcsMap[fnName]; ok {
			execOutput, err := fn(params)
			if err != nil {
				return nil, err
			}

			output = execOutput
		}
	}

	if line.HasReplacement() {
		replacement := line.Replacement()
		retValueBytes, err := app.writeElement(grammar, replacement)
		if err != nil {
			return nil, err
		}

		output = retValueBytes
	}

	return output, nil
}

func (app *application) writeTokens(grammar grammars.Grammar, tokens tokens.Tokens) ([]byte, map[string][][]byte, error) {
	output := []byte{}
	mp := map[string][][]byte{}
	list := tokens.List()
	for _, oneToken := range list {
		name := oneToken.Name()
		retBytes, retMultiLine, err := app.writeToken(grammar, oneToken)
		if err != nil {
			return nil, nil, err
		}

		mp[name] = retMultiLine
		output = append(output, retBytes...)
	}

	return output, mp, nil
}

func (app *application) writeToken(grammar grammars.Grammar, token tokens.Token) ([]byte, [][]byte, error) {
	name := token.Name()
	cardinality := token.Cardinality()
	if !cardinality.HasMax() {
		str := fmt.Sprintf("the cardinality, in the token (name: %s) must contain a max in order to be written", name)
		return nil, nil, errors.New(str)
	}

	pMax := cardinality.Max()
	min := cardinality.Min()
	if *pMax != min {
		str := fmt.Sprintf("the cardinality, in the token (name: %s) must contain a min (%d) and a max (%d) that are equal in order to be written", name, min, *pMax)
		return nil, nil, errors.New(str)
	}

	element := token.Element()
	elementBytes, err := app.writeElement(grammar, element)
	if err != nil {
		return nil, nil, err
	}

	output := []byte{}
	multiLine := [][]byte{}
	castedMin := int(min)
	for i := 0; i < castedMin; i++ {
		multiLine = append(multiLine, elementBytes)
		output = append(output, elementBytes...)
	}

	return output, multiLine, nil
}

func (app *application) writeElement(grammar grammars.Grammar, element elements.Element) ([]byte, error) {
	if element.IsBlock() {
		blockName := element.Block()
		block, err := grammar.Blocks().Fetch(blockName)
		if err != nil {
			return nil, err
		}

		return app.writeBlock(grammar, block)
	}

	ruleName := element.Rule()
	rule, err := grammar.Rules().Fetch(ruleName)
	if err != nil {
		return nil, err
	}

	return rule.Bytes(), nil
}

func (app *application) grammarToNFT(
	grammar grammars.Grammar,
) (nfts.NFT, error) {
	version := grammar.Version()
	versionNFT, err := app.uintAdapter.ToNFT(uint64(version))
	if err != nil {
		return nil, err
	}

	rules := grammar.Rules()
	rulesNFT, err := app.ruleAdapter.InstancesToNFT(rules)
	if err != nil {
		return nil, err
	}

	blocks := grammar.Blocks()
	blocksNFT, retBlocksNFTs, err := app.blocksToNFT(
		rulesNFT,
		blocks,
	)

	if err != nil {
		return nil, err
	}

	root := grammar.Root()
	rootNFT, err := app.elementToNFT(
		[]string{},
		rulesNFT,
		retBlocksNFTs,
		root,
	)

	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		versionNFT,
		rootNFT,
		rulesNFT,
		blocksNFT,
	}

	if grammar.HasOmissions() {
		omissions := grammar.Omissions()
		omissionsNFT, err := app.elementsToNFT(
			[]string{},
			rulesNFT,
			retBlocksNFTs,
			omissions,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, omissionsNFT)
	}

	return app.nftsListToNFT(list, "")
}

func (app *application) blocksToNFT(
	rules nfts.NFT,
	blocksIns blocks.Blocks,
) (nfts.NFT, nfts.NFTs, error) {
	var blocks nfts.NFTs
	output := []nfts.NFT{}
	list := blocksIns.List()
	for _, oneBlock := range list {
		retNFT, err := app.blockToNFT(
			[]string{},
			rules,
			blocks,
			oneBlock,
		)

		if err != nil {
			return nil, nil, err
		}

		output = append(output, retNFT)
		retBlocks, err := app.nftsBuilder.Create().WithList(output).Now()
		if err != nil {
			return nil, nil, err
		}

		blocks = retBlocks
	}

	nft, err := app.nftsListToNFT(output, "")
	if err != nil {
		return nil, nil, err
	}

	return nft, blocks, nil
}

func (app *application) blockToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	block blocks.Block,
) (nfts.NFT, error) {
	list := []nfts.NFT{}
	name := block.Name()
	parentBlockNames = append(parentBlockNames, name)
	if block.HasLines() {
		lines := block.Lines()
		linesNFT, err := app.linesToNFT(
			parentBlockNames,
			rules,
			blocks,
			lines,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, linesNFT)
	}

	if block.HasLine() {
		line := block.Line()
		lineNFT, err := app.lineToNFT(
			parentBlockNames,
			rules,
			blocks,
			line,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, lineNFT)
	}

	if block.HasSuites() {
		suites := block.Suites()
		suitesNFT, err := app.suitesToNFT(
			parentBlockNames,
			rules,
			blocks,
			suites,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, suitesNFT)
	}

	nft, err := app.nftsListToNFT(list, name)
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) suitesToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	suites suites.Suites,
) (nfts.NFT, error) {
	output := []nfts.NFT{}
	list := suites.List()
	for _, oneSuite := range list {
		retNFT, err := app.suiteToNFT(
			parentBlockNames,
			rules,
			blocks,
			oneSuite,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retNFT)
	}

	nft, err := app.nftsListToNFT(output, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) suiteToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	suite suites.Suite,
) (nfts.NFT, error) {
	element := suite.Element()
	elementNFT, err := app.elementToNFT(
		parentBlockNames,
		rules,
		blocks,
		element,
	)

	if err != nil {
		return nil, err
	}

	isFail := byte(0)
	if suite.IsFail() {
		isFail = byte(1)
	}

	isFailNFT, err := app.nftBuilder.Create().
		WithByte(isFail).
		Now()

	if err != nil {
		return nil, err
	}

	name := suite.Name()
	nft, err := app.nftsListToNFT([]nfts.NFT{
		elementNFT,
		isFailNFT,
	}, name)

	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) linesToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	lines lines.Lines,
) (nfts.NFT, error) {
	output := []nfts.NFT{}
	list := lines.List()
	for _, oneLine := range list {
		retNFT, err := app.lineToNFT(
			parentBlockNames,
			rules,
			blocks,
			oneLine,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retNFT)
	}

	nft, err := app.nftsListToNFT(output, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) lineToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	line lines.Line,
) (nfts.NFT, error) {
	tokens := line.Tokens()
	tokensNFT, err := app.tokensToNFT(
		parentBlockNames,
		rules,
		blocks,
		tokens,
	)

	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		tokensNFT,
	}

	if line.HasExecution() {
		execution := line.Execution()
		execNFT, err := app.executionToNFT(
			parentBlockNames,
			rules,
			blocks,
			execution,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, execNFT)
	}

	if line.HasReplacement() {
		replacement := line.Replacement()
		execNFT, err := app.elementToNFT(
			parentBlockNames,
			rules,
			blocks,
			replacement,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, execNFT)
	}

	nft, err := app.nftsListToNFT(list, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) tokensToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	tokens tokens.Tokens,
) (nfts.NFT, error) {
	output := []nfts.NFT{}
	list := tokens.List()
	for _, oneToken := range list {
		retNFT, err := app.tokenToNFT(
			parentBlockNames,
			rules,
			blocks,
			oneToken,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retNFT)
	}

	nft, err := app.nftsListToNFT(output, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) tokenToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	token tokens.Token,
) (nfts.NFT, error) {
	element := token.Element()
	retElementAST, err := app.elementToNFT(
		parentBlockNames,
		rules,
		blocks,
		element,
	)

	if err != nil {
		return nil, err
	}

	cardinality := token.Cardinality()
	cardinalityNFT, err := app.cardinalityAdapter.ToNFT(cardinality)
	if err != nil {
		return nil, err
	}

	name := token.Name()
	nft, err := app.nftsListToNFT([]nfts.NFT{
		retElementAST,
		cardinalityNFT,
	}, name)

	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) executionToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	execution executions.Execution,
) (nfts.NFT, error) {
	funcName := execution.FuncName()
	funcNameNFT, err := app.stringToNFT(funcName)
	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		funcNameNFT,
	}

	if execution.HasParameters() {
		parameters := execution.Parameters()
		parameterNFT, err := app.parametersToNFT(
			parentBlockNames,
			rules,
			blocks,
			parameters,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, parameterNFT)
	}

	nft, err := app.nftsListToNFT(list, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) parametersToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	parameters parameters.Parameters,
) (nfts.NFT, error) {
	output := []nfts.NFT{}
	list := parameters.List()
	for _, oneParameter := range list {
		retNFT, err := app.parameterToNFT(
			parentBlockNames,
			rules,
			blocks,
			oneParameter,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retNFT)
	}

	nft, err := app.nftsListToNFT(output, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) parameterToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	parameter parameters.Parameter,
) (nfts.NFT, error) {
	element := parameter.Element()
	elementNFT, err := app.elementToNFT(
		parentBlockNames,
		rules,
		blocks,
		element,
	)

	if err != nil {
		return nil, err
	}

	index := parameter.Index()
	indexNFT, err := app.uintAdapter.ToNFT(uint64(index))
	if err != nil {
		return nil, err
	}

	name := parameter.Name()
	nameNFT, err := app.stringToNFT(name)
	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		elementNFT,
		indexNFT,
		nameNFT,
	}

	nft, err := app.nftsListToNFT(list, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) stringToNFT(
	value string,
) (nfts.NFT, error) {
	list := []nfts.NFT{}
	data := []byte(value)
	for _, oneByte := range data {
		nft, err := app.nftBuilder.Create().WithByte(oneByte).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, nft)
	}

	return app.nftsListToNFT(list, "")
}

func (app *application) elementsToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	elements elements.Elements,
) (nfts.NFT, error) {
	output := []nfts.NFT{}
	updatedBlocks := blocks
	list := elements.List()
	for _, oneElement := range list {
		retNFT, err := app.elementToNFT(
			parentBlockNames,
			rules,
			updatedBlocks,
			oneElement,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retNFT)
	}

	nft, err := app.nftsListToNFT(output, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *application) elementToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	element elements.Element,
) (nfts.NFT, error) {
	if element.IsBlock() {
		name := element.Block()
		for _, oneParentBlockName := range parentBlockNames {
			if name != oneParentBlockName {
				continue
			}

			level := uint(len(parentBlockNames))
			return app.nftBuilder.Create().
				WithRecursive(level).
				Now()
		}

		if blocks == nil {
			str := fmt.Sprintf("the block reference (name: %s) could not be found in element", name)
			return nil, errors.New(str)
		}

		retIns, err := blocks.Fetch(name)
		if err != nil {
			return nil, err
		}

		return retIns, nil
	}

	name := element.Rule()
	retNFT, err := rules.Fetch(name)
	if err != nil {
		return nil, err
	}

	return retNFT, nil
}

func (app *application) nftsListToNFT(list []nfts.NFT, name string) (nfts.NFT, error) {
	builder := app.nftsBuilder.Create().WithList(list)
	if name != "" {
		builder.WithName(name)
	}

	nfts, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return app.nftBuilder.Create().
		WithNFTs(nfts).
		Now()
}
