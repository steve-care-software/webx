package applications

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/steve-care-software/webx/engine/applications/stackframes"
	"github.com/steve-care-software/webx/engine/applications/stackframes/cursors"
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

// SyscallFn represents the syscall func
type SyscallFn func(map[string][]byte) error

// NewApplication creates a new application
func NewApplication() Application {
	cursorApp := cursors.NewApplication()
	stackframeApp, err := stackframes.NewFactory().Create()
	if err != nil {
		panic(err)
	}

	elementsAdapter := instructions.NewElementsAdapter()
	grammarParserAdapter := grammars.NewParserAdapter()
	grammarNFTAdapter := grammars.NewNFTAdapter()
	grammarComposeAdapter := grammars.NewComposeAdapter()
	programParserAdapter := programs.NewParserAdapter()
	return createApplication(
		stackframeApp,
		elementsAdapter,
		grammarParserAdapter,
		grammarNFTAdapter,
		grammarComposeAdapter,
		programParserAdapter,
		map[string]SyscallFn{
			"math_operation_arithmetic_add": func(params map[string][]byte) error {
				if firstBytes, ok := params["first"]; ok {
					if secondBytes, ok := params["second"]; ok {
						first, _ := big.NewInt(0).SetString(string(firstBytes), 0)
						if first == nil {
							return errors.New("the values passed to the first paramter could not be casted to an int")
						}

						second, _ := big.NewInt(0).SetString(string(secondBytes), 0)
						if second == nil {
							return errors.New("the values passed to the second paramter could not be casted to an int")
						}

						value := first.Add(first, second)
						fmt.Printf("\n%s, %s, %d\n", params["first"], params["second"], value.Int64())
						return nil
					}

					return errors.New("the second parameter could not be found")
				}

				return errors.New("the first parameter could not be found")
			},
			"cursor_push": func(params map[string][]byte) error {
				if valueStrBytes, ok := params["value"]; ok {
					if kindStrBytes, ok := params["kind"]; ok {
						kind, err := strconv.Atoi(string(kindStrBytes))
						if err != nil {
							return err
						}

						return cursorApp.PushAsStringBytes(valueStrBytes, uint8(kind))
					}

					return errors.New("the kind parameter could not be found")
				}

				return errors.New("the value parameter could not be found")
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
	Suites(grammar grammars.Grammar) error
}
