package applications

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

// SyscallFn represents the syscall func
type SyscallFn func(map[string][]byte) (any, error)

// NewApplication creates a new application
func NewApplication() Application {
	elementsAdapter := instructions.NewElementsAdapter()
	grammarParserAdapter := grammars.NewParserAdapter()
	grammarNFTAdapter := grammars.NewNFTAdapter()
	grammarComposeAdapter := grammars.NewComposeAdapter()
	return createApplication(
		elementsAdapter,
		grammarParserAdapter,
		grammarNFTAdapter,
		grammarComposeAdapter,
		map[string]SyscallFn{
			"math_operation_arithmetic_add": func(params map[string][]byte) (any, error) {
				if firstBytes, ok := params["first"]; ok {
					if secondBytes, ok := params["second"]; ok {
						first, _ := big.NewInt(0).SetString(string(firstBytes), 0)
						if first == nil {
							return nil, errors.New("the values passed to the first paramter could not be casted to an int")
						}

						second, _ := big.NewInt(0).SetString(string(secondBytes), 0)
						if second == nil {
							return nil, errors.New("the values passed to the second paramter could not be casted to an int")
						}

						value := first.Add(first, second)
						fmt.Printf("\n%s, %s\n", params["first"], params["second"])
						return value.Int64(), nil
					}

					return nil, errors.New("the first parameter could not be found")
				}

				return nil, errors.New("the second parameter could not be found")
			},
		},
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

	// Interpret interprets the input and returns the stack
	Interpret(program programs.Program) (stacks.Stack, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) ([]byte, error)

	// Suite executes the test suite of the provided blockName in the grammar
	Suite(grammar grammars.Grammar, blockName string) ([]byte, error)
}
