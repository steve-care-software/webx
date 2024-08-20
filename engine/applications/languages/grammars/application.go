package grammars

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/cardinalities/uints"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/nfts"
)

type application struct {
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
	return app.grammarToNFT(grammar)
}

// Decompile decompiles an NFT to a grammar instance
func (app *application) Decompile(ast nfts.NFT) (grammars.Grammar, error) {
	return nil, nil
}

// Compose composes a grammar instance to a grammar input
func (app *application) Compose(grammar grammars.Grammar) ([]byte, error) {
	return nil, nil
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
	lines := block.Lines()
	name := block.Name()
	parentBlockNames = append(parentBlockNames, name)
	linesNFT, err := app.linesToNFT(
		parentBlockNames,
		rules,
		blocks,
		lines,
	)

	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		linesNFT,
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

	if execution.HasElements() {
		elements := execution.Elements()
		elementNFT, err := app.elementsToNFT(
			parentBlockNames,
			rules,
			blocks,
			elements,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, elementNFT)
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
