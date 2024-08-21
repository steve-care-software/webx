package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls/parameters"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

type application struct {
	programComposer       programs.ComposeAdapter
	grammarParserAdapter  grammars.ParserAdapter
	grammarNFTAdapter     grammars.NFTAdapter
	grammarComposeAdapter grammars.ComposeAdapter
	syscalls              map[string]SyscallFn
}

func createApplication(
	programComposer programs.ComposeAdapter,
	grammarParserAdapter grammars.ParserAdapter,
	grammarNFTAdapter grammars.NFTAdapter,
	grammarComposeAdapter grammars.ComposeAdapter,
	syscalls map[string]SyscallFn,
) Application {
	out := application{
		programComposer:       programComposer,
		grammarParserAdapter:  grammarParserAdapter,
		grammarNFTAdapter:     grammarNFTAdapter,
		grammarComposeAdapter: grammarComposeAdapter,
		syscalls:              syscalls,
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
	root := program.Root()
	app.interpretElement(
		program,
		root,
	)

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

func (app *application) interpretInstruction(
	program programs.Program,
	instruction instructions.Instruction,
) error {
	tokens := instruction.Tokens()
	return app.interpretTokens(
		program,
		tokens,
	)
}

func (app *application) interpretTokens(
	program programs.Program,
	tokens tokens.Tokens,
) error {
	list := tokens.List()
	for _, oneToken := range list {
		err := app.interpretToken(
			program,
			oneToken,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) interpretToken(
	program programs.Program,
	token tokens.Token,
) error {
	elements := token.Elements()
	return app.interpretElements(
		program,
		elements,
	)
}

func (app *application) interpretElements(
	program programs.Program,
	elements elements.Elements,
) error {
	list := elements.List()
	for _, oneElement := range list {
		err := app.interpretElement(
			program,
			oneElement,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) interpretElement(
	program programs.Program,
	element elements.Element,
) error {
	if element.IsRule() {
		return nil
	}

	if element.IsSyscall() {
		syscall := element.Syscall()
		return app.interpretSyscall(
			program,
			syscall,
		)
	}

	insName := element.Instruction()
	instruction, err := program.Instructions().Fetch(insName)
	if err != nil {
		return err
	}

	return app.interpretInstruction(
		program,
		instruction,
	)
}

func (app *application) interpretSyscall(
	program programs.Program,
	sysCall syscalls.Syscall,
) error {
	name := sysCall.Name()
	fnName := sysCall.FuncName()
	mpParams := map[string][]byte{}
	if sysCall.HasParameters() {
		parameters := sysCall.Parameters()
		retMapParams, err := app.fetchParameters(program, parameters)
		if err != nil {
			str := fmt.Sprintf("there was an error while fetching the syscall (blockName: %s, sysCallFn: %s) parameters: %s", name, fnName, err.Error())
			return errors.New(str)
		}

		mpParams = retMapParams
	}

	if fn, ok := app.syscalls[fnName]; ok {
		_, err := fn(mpParams)
		if err != nil {
			return err
		}
	}

	str := fmt.Sprintf("the sysCall (sysCallFn: %s) declared in block (name: %s) does not exists", fnName, name)
	return errors.New(str)
}

func (app *application) fetchParameters(
	program programs.Program,
	parameters parameters.Parameters,
) (map[string][]byte, error) {
	output := map[string][]byte{}
	list := parameters.List()
	for _, oneParameter := range list {
		name, value, err := app.fetchParameter(
			program,
			oneParameter,
		)

		if err != nil {
			return nil, err
		}

		output[name] = value
	}

	return output, nil
}

func (app *application) fetchParameter(
	program programs.Program,
	parameter parameters.Parameter,
) (string, []byte, error) {
	token := parameter.Token()
	retBytes, err := app.programComposer.ToBytes(program, token)
	if err != nil {
		return "", nil, err
	}

	return parameter.Name(), retBytes, nil
}
