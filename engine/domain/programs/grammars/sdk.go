package grammars

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/webx/engine/domain/nfts"
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
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls"
)

// CoreFn represents a core fn
type CoreFn func(input map[string][]byte) ([]byte, error)

const llA = "a"
const llB = "b"
const llC = "c"
const llD = "d"
const llE = "e"
const llF = "f"
const llG = "g"
const llH = "h"
const llI = "i"
const llJ = "j"
const llK = "k"
const llL = "l"
const llM = "m"
const llN = "n"
const llO = "o"
const llP = "p"
const llQ = "q"
const llR = "r"
const llS = "s"
const llT = "t"
const llU = "u"
const llV = "v"
const llW = "w"
const llX = "x"
const llY = "y"
const llZ = "z"

const ulA = "A"
const ulB = "B"
const ulC = "C"
const ulD = "D"
const ulE = "E"
const ulF = "F"
const ulG = "G"
const ulH = "H"
const ulI = "I"
const ulJ = "J"
const ulK = "K"
const ulL = "L"
const ulM = "M"
const ulN = "N"
const ulO = "O"
const ulP = "P"
const ulQ = "Q"
const ulR = "R"
const ulS = "S"
const ulT = "T"
const ulU = "U"
const ulV = "V"
const ulW = "W"
const ulX = "X"
const ulY = "Y"
const ulZ = "Z"

const nZero = "0"
const nOne = "1"
const nTwo = "2"
const nTree = "3"
const nFour = "4"
const nFive = "5"
const nSix = "6"
const nSeven = "7"
const nHeight = "8"
const nNine = "9"

const ruleValueEscape = "\\"
const ruleValuePrefix = "\""
const ruleValueSuffix = "\""
const ruleNameSeparator = "_"
const ruleNameValueSeparator = ":"
const cardinalityOpen = "["
const cardinalityClose = "]"
const cardinalitySeparator = ","
const cardinalityZeroPlus = "*"
const cardinalityOnePlus = "+"
const indexOpen = "["
const indexClose = "]"
const parameterSeparator = ":"
const tokenReference = "."
const linesSeparator = "|"
const lineSeparator = "-"
const funcNameSeparator = "_"
const blockDefinitionSeparator = ":"
const failSeparator = "@"
const suiteLineSuffix = "."
const blockSuffix = ";"
const suiteSeparatorPrefix = "---"
const versionPrefix = "v"
const versionSuffix = ";"
const rootPrefix = ">"
const rootSuffix = ";"
const omissionPrefix = "#"
const omissionSuffix = ";"
const filterBytes = " \n\r\t"

// NewComposeAdapter creates a new composer adapter
func NewComposeAdapter() ComposeAdapter {
	return createComposeAdapter(
		map[string]CoreFn{
			"math_operation_arithmetic_addition": func(input map[string][]byte) ([]byte, error) {
				if firstBytes, ok := input["first"]; ok {
					if secondBytes, ok := input["second"]; ok {
						pFirst, _ := big.NewInt(int64(0)).SetString(string(firstBytes), 10)
						if pFirst == nil {
							return nil, errors.New("the first value could not be converted to a number")
						}

						pSecond, _ := big.NewInt(int64(0)).SetString(string(secondBytes), 10)
						if pSecond == nil {
							return nil, errors.New("the second value could not be converted to a number")
						}

						return []byte(pFirst.Add(pFirst, pSecond).String()), nil
					}

					return nil, errors.New("the second value was not defined")
				}

				return nil, errors.New("the first value was not defined")
			},
		},
	)
}

// NewNFTAdapter creates a new nft adapter
func NewNFTAdapter() NFTAdapter {
	ruleAdapter := rules.NewAdapter()
	cardinalityAdapter := cardinalities.NewAdapter()
	uintAdapter := uints.NewAdapter()
	nftsBuilder := nfts.NewBuilder()
	nftBuilder := nfts.NewNFTBuilder()
	return createNFTAdapter(
		ruleAdapter,
		cardinalityAdapter,
		uintAdapter,
		nftsBuilder,
		nftBuilder,
	)
}

// NewParserAdapter creates a new parser adapter
func NewParserAdapter() ParserAdapter {
	grammarBuilder := NewBuilder()
	blocksBuilder := blocks.NewBuilder()
	blockBuilder := blocks.NewBlockBuilder()
	suitesBuilder := suites.NewBuilder()
	suiteBuilder := suites.NewSuiteBuilder()
	linesBuilder := lines.NewBuilder()
	lineBuilder := lines.NewLineBuilder()
	executionBuilder := executions.NewBuilder()
	parametersBuilder := parameters.NewBuilder()
	parameterBuilder := parameters.NewParameterBuilder()
	tokensBuilder := tokens.NewBuilder()
	tokenBuilder := tokens.NewTokenBuilder()
	elementsBuilder := elements.NewBuilder()
	elementBuilder := elements.NewElementBuilder()
	rulesBuilder := rules.NewBuilder()
	ruleBuilder := rules.NewRuleBuilder()
	cardinalityBuilder := cardinalities.NewBuilder()
	possibleLetters := createPossibleLetters()
	possibleLowerCaseLetters := createPossibleLowerCaseLetters()
	possibleUpperCaseLetters := createPossibleUpperCaseLetters()
	possibleNumbers := createPossibleNumbers()
	possibleFuncNameCharacters := createPossibleFuncNameCharacters()
	return createParserAdapter(
		grammarBuilder,
		blocksBuilder,
		blockBuilder,
		suitesBuilder,
		suiteBuilder,
		linesBuilder,
		lineBuilder,
		executionBuilder,
		parametersBuilder,
		parameterBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
		rulesBuilder,
		ruleBuilder,
		cardinalityBuilder,
		[]byte(filterBytes),
		[]byte(suiteSeparatorPrefix),
		possibleLetters,
		possibleLowerCaseLetters,
		possibleUpperCaseLetters,
		possibleNumbers,
		possibleFuncNameCharacters,
		[]byte(omissionPrefix)[0],
		[]byte(omissionSuffix)[0],
		[]byte(versionPrefix)[0],
		[]byte(versionSuffix)[0],
		[]byte(rootPrefix)[0],
		[]byte(rootSuffix)[0],
		[]byte(blockSuffix)[0],
		[]byte(suiteLineSuffix)[0],
		[]byte(failSeparator)[0],
		[]byte(blockDefinitionSeparator)[0],
		[]byte(linesSeparator)[0],
		[]byte(lineSeparator)[0],
		[]byte(tokenReference)[0],
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleNameValueSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(cardinalityZeroPlus)[0],
		[]byte(cardinalityOnePlus)[0],
		[]byte(indexOpen)[0],
		[]byte(indexClose)[0],
		[]byte(parameterSeparator)[0],
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// ParserAdapter represents the grammar parser adapter
type ParserAdapter interface {
	// ToGrammar takes the input and converts it to a grammar instance and the remaining data
	ToGrammar(input []byte) (Grammar, []byte, error)

	// ToBytes takes a grammar and returns the bytes
	ToBytes(grammar Grammar) ([]byte, error)
}

// NFTAdapter represents the grammar nft adapter
type NFTAdapter interface {
	// ToNFT converts a grammar instance to an NFT
	ToNFT(grammar Grammar) (nfts.NFT, error)

	// ToGrammar converts an NFT to a grammar instance
	ToGrammar(nft nfts.NFT) (Grammar, error)
}

// ComposeAdapter represents the grammar compose adapter
type ComposeAdapter interface {
	// ToBytes takes a grammar and a blockname and returns its bytes
	ToBytes(grammar Grammar, blockName string) ([]byte, error)
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithRoot(root elements.Element) Builder
	WithRules(rules rules.Rules) Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithSyscalls(syscalls syscalls.Syscalls) Builder
	WithOmissions(omissions elements.Elements) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Version() uint
	Root() elements.Element
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasSyscalls() bool
	Syscalls() syscalls.Syscalls
	HasOmissions() bool
	Omissions() elements.Elements
}
