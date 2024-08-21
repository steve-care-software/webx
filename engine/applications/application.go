package applications

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

type application struct {
	grammarParserAdapter  grammars.ParserAdapter
	grammarNFTAdapter     grammars.NFTAdapter
	grammarComposeAdapter grammars.ComposeAdapter
}

func createApplication(
	grammarParserAdapter grammars.ParserAdapter,
	grammarNFTAdapter grammars.NFTAdapter,
	grammarComposeAdapter grammars.ComposeAdapter,
) Application {
	out := application{
		grammarParserAdapter:  grammarParserAdapter,
		grammarNFTAdapter:     grammarNFTAdapter,
		grammarComposeAdapter: grammarComposeAdapter,
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
	return app.grammarComposeAdapter.ToBytes(grammar, blockName)
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
