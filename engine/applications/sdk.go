package applications

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	grammarParserAdapter := grammars.NewParserAdapter()
	grammarNFTAdapter := grammars.NewNFTAdapter()
	grammarComposeAdapter := grammars.NewComposeAdapter()
	return createApplication(
		grammarParserAdapter,
		grammarNFTAdapter,
		grammarComposeAdapter,
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
