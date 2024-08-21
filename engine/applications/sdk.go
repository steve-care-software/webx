package applications

import (
	"errors"
	"math/big"

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

// NewApplication creates a new application
func NewApplication() Application {
	grammarParserAdapter := grammars.NewParserAdapter()
	grammarNFTAdapter := grammars.NewNFTAdapter()
	ruleAdapter := rules.NewAdapter()
	cardinalityAdapter := cardinalities.NewAdapter()
	uintAdapter := uints.NewAdapter()
	nftsBuilder := nfts.NewBuilder()
	nftBuilder := nfts.NewNFTBuilder()
	grammarBuilder := grammars.NewBuilder()
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
	return createApplication(
		grammarParserAdapter,
		grammarNFTAdapter,
		ruleAdapter,
		cardinalityAdapter,
		uintAdapter,
		nftsBuilder,
		nftBuilder,
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

// Application represents the grammar application
type Application interface {
	// ParseGrammar parses an input and creates a Grammar instance
	ParseGrammar(input []byte) (grammars.Grammar, []byte, error)

	// CompileGrammar compiles a grammar to an NFT
	CompileGrammar(grammar grammars.Grammar) (nfts.NFT, error)

	// DecompileGrammar decompiles an NFT into a grammar instance
	DecompileGrammar(nft nfts.NFT) (grammars.Grammar, error)

	// ComposeBlock fetches a blockName from the grammar and composes an output
	ComposeBlock(grammar grammars.Grammar, blockName string) ([]byte, error)

	// ParseProgram takes a grammar and an input, parses it and returns the program
	ParseProgram(grammar grammars.Grammar, input []byte) (programs.Program, error)

	// CompileProgram compiles a program to an NFT
	CompileProgram(program programs.Program) (nfts.NFT, error)

	// DecompileProgram decompiles an NFT into a program instance
	DecompileProgram(nft nfts.NFT) (programs.Program, error)

	// ComposeProgram takes the program and composes an output
	ComposeProgram(program programs.Program) ([]byte, error)

	// Interpret interprets the input and returns the stack
	Interpret(program programs.Program) (stacks.Stack, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) ([]byte, error)

	// Suite executes the test suite of the provided blockName in the grammar
	Suite(grammar grammars.Grammar, blockName string) ([]byte, error)
}
