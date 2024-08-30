package grammars

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values/references"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/processors"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities/uints"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
)

type nftAdapter struct {
	ruleAdapter        rules.Adapter
	cardinalityAdapter cardinalities.Adapter
	uintAdapter        uints.Adapter
	nftsBuilder        nfts.Builder
	nftBuilder         nfts.NFTBuilder
}

func createNFTAdapter(
	ruleAdapter rules.Adapter,
	cardinalityAdapter cardinalities.Adapter,
	uintAdapter uints.Adapter,
	nftsBuilder nfts.Builder,
	nftBuilder nfts.NFTBuilder,
) NFTAdapter {
	out := nftAdapter{
		ruleAdapter:        ruleAdapter,
		cardinalityAdapter: cardinalityAdapter,
		uintAdapter:        uintAdapter,
		nftsBuilder:        nftsBuilder,
		nftBuilder:         nftBuilder,
	}

	return &out
}

// ToNFT converts a grammar instance to an NFT
func (app *nftAdapter) ToNFT(grammar Grammar) (nfts.NFT, error) {
	return app.grammarToNFT(grammar)
}

// ToGrammar converts an NFT to a grammar instance
func (app *nftAdapter) ToGrammar(nft nfts.NFT) (Grammar, error) {
	return nil, nil
}

func (app *nftAdapter) grammarToNFT(
	grammar Grammar,
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

func (app *nftAdapter) blocksToNFT(
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

func (app *nftAdapter) blockToNFT(
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

func (app *nftAdapter) suitesToNFT(
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

func (app *nftAdapter) suiteToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	suite suites.Suite,
) (nfts.NFT, error) {
	nftsList := []nfts.NFT{}
	inputAsBytes := suite.Input()
	for _, oneByte := range inputAsBytes {
		nft, err := app.nftBuilder.Create().WithByte(oneByte).Now()
		if err != nil {
			return nil, err
		}

		nftsList = append(nftsList, nft)
	}

	suiteNFTs, err := app.nftsBuilder.Create().WithList(nftsList).Now()
	if err != nil {
		return nil, err
	}

	suiteNFT, err := app.nftBuilder.Create().WithNFTs(suiteNFTs).Now()
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
		suiteNFT,
		isFailNFT,
	}, name)

	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *nftAdapter) linesToNFT(
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

func (app *nftAdapter) lineToNFT(
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

	if line.HasProcessor() {
		processor := line.Processor()
		processNFT, err := app.processorToNFT(
			parentBlockNames,
			rules,
			blocks,
			processor,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, processNFT)
	}

	if line.HasSyscall() {
		syscall := line.Syscall()
		syscallNFT, err := app.executionToNFT(
			parentBlockNames,
			rules,
			blocks,
			syscall,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, syscallNFT)
	}

	nft, err := app.nftsListToNFT(list, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *nftAdapter) processorToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	processor processors.Processor,
) (nfts.NFT, error) {
	if processor.IsExecution() {
		execution := processor.Execution()
		return app.executionToNFT(
			parentBlockNames,
			rules,
			blocks,
			execution,
		)
	}

	replacement := processor.Replacement()
	return app.elementToNFT(
		parentBlockNames,
		rules,
		blocks,
		replacement,
	)
}

func (app *nftAdapter) tokensToNFT(
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

func (app *nftAdapter) tokenToNFT(
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

func (app *nftAdapter) executionToNFT(
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

func (app *nftAdapter) parametersToNFT(
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

func (app *nftAdapter) parameterToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	parameter parameters.Parameter,
) (nfts.NFT, error) {
	value := parameter.Value()
	retNFT, err := app.valueToNFT(
		parentBlockNames,
		rules,
		blocks,
		value,
	)

	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		retNFT,
	}

	name := parameter.Name()
	nft, err := app.nftsListToNFT(list, name)
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *nftAdapter) valueToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	value values.Value,
) (nfts.NFT, error) {
	if value.IsReference() {
		reference := value.Reference()
		return app.referenceToNFT(
			parentBlockNames,
			rules,
			blocks,
			reference,
		)
	}

	bytes := value.Bytes()
	return app.bytesToNFT(bytes)
}

func (app *nftAdapter) referenceToNFT(
	parentBlockNames []string,
	rules nfts.NFT,
	blocks nfts.NFTs,
	reference references.Reference,
) (nfts.NFT, error) {
	element := reference.Element()
	elementNFT, err := app.elementToNFT(
		parentBlockNames,
		rules,
		blocks,
		element,
	)

	if err != nil {
		return nil, err
	}

	index := reference.Index()
	indexNFT, err := app.uintAdapter.ToNFT(uint64(index))
	if err != nil {
		return nil, err
	}

	list := []nfts.NFT{
		elementNFT,
		indexNFT,
	}

	nft, err := app.nftsListToNFT(list, "")
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (app *nftAdapter) stringToNFT(
	value string,
) (nfts.NFT, error) {
	data := []byte(value)
	return app.bytesToNFT(data)
}

func (app *nftAdapter) bytesToNFT(
	data []byte,
) (nfts.NFT, error) {
	list := []nfts.NFT{}
	for _, oneByte := range data {
		nft, err := app.nftBuilder.Create().WithByte(oneByte).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, nft)
	}

	return app.nftsListToNFT(list, "")
}

func (app *nftAdapter) elementsToNFT(
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

func (app *nftAdapter) elementToNFT(
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

func (app *nftAdapter) nftsListToNFT(list []nfts.NFT, name string) (nfts.NFT, error) {
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
