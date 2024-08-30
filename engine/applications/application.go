package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/applications/stackframes"
	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

type application struct {
	stackFrameApp         stackframes.Application
	elementsAdapter       instructions.ElementsAdapter
	grammarParserAdapter  grammars.ParserAdapter
	grammarComposeAdapter grammars.ComposeAdapter
	programParserAdapter  programs.ParserAdapter
	syscalls              map[string]SyscallFn
}

func createApplication(
	stackFrameApp stackframes.Application,
	elementsAdapter instructions.ElementsAdapter,
	grammarParserAdapter grammars.ParserAdapter,
	grammarComposeAdapter grammars.ComposeAdapter,
	programParserAdapter programs.ParserAdapter,
	syscalls map[string]SyscallFn,
) Application {
	out := application{
		stackFrameApp:         stackFrameApp,
		elementsAdapter:       elementsAdapter,
		grammarParserAdapter:  grammarParserAdapter,
		grammarComposeAdapter: grammarComposeAdapter,
		programParserAdapter:  programParserAdapter,
		syscalls:              syscalls,
	}

	return &out
}

// ParseGrammar parses an input and creates a Grammar instance
func (app *application) ParseGrammar(input []byte) (grammars.Grammar, []byte, error) {
	return app.grammarParserAdapter.ToGrammar(input)
}

// ComposeBlock fetches a blockName from the grammar and composes an output
func (app *application) ComposeBlock(grammar grammars.Grammar, blockName string) ([]byte, error) {
	return app.grammarComposeAdapter.ToBytes(grammar, blockName)
}

// ParseProgram takes a grammar and an input, parses it and returns the program
func (app *application) ParseProgram(grammar grammars.Grammar, input []byte) (programs.Program, error) {
	return nil, nil
}

// Interpret interprets the input and returns the stack
func (app *application) Interpret(program programs.Program) (stacks.Stack, error) {
	root := program.Root()
	app.interpretElement(
		nil,
		root,
	)

	return app.stackFrameApp.Root().Fetch()
}

// Suites executes all the test suites of the grammar
func (app *application) Suites(grammar grammars.Grammar) error {
	blocksList := grammar.Blocks().List()
	for _, oneBlock := range blocksList {
		if !oneBlock.HasSuites() {
			continue
		}

		blockName := oneBlock.Name()
		suitesList := oneBlock.Suites().List()
		for idx, oneSuite := range suitesList {
			err := app.interpretSuite(
				grammar,
				blockName,
				oneSuite,
			)

			prefix := fmt.Sprintf("block (name: %s) index (%d) suite (%s)", blockName, idx, oneSuite.Name())
			if oneSuite.IsFail() {
				if err == nil {
					str := fmt.Sprintf("%s: the suite was expected to FAIL but succeeded!", prefix)
					return errors.New(str)
				}

				continue
			}

			if err != nil {
				str := fmt.Sprintf("%s the suite was expected to SUCCEED but failed --- error: %s", prefix, err.Error())
				return errors.New(str)
			}
		}
	}
	return nil
}

func (app *application) interpretSuite(
	grammar grammars.Grammar,
	blockName string,
	suite suites.Suite,
) error {
	program, retRemaining, err := app.programParserAdapter.ToProgramWithRoot(
		grammar,
		blockName,
		suite.Input(),
	)

	if err != nil {
		return err
	}

	if len(retRemaining) != 0 {
		str := fmt.Sprintf("the bytes (%s) were remaining", retRemaining)
		return errors.New(str)
	}

	_, err = app.Interpret(program)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) interpretInstruction(
	instruction instructions.Instruction,
) error {
	tokens := instruction.Tokens()
	if instruction.HasSyscall() {
		syscall := instruction.Syscall()
		err := app.interpretSyscall(
			tokens,
			syscall,
		)

		if err != nil {
			return err
		}
	}

	return app.interpretTokens(
		tokens,
	)
}

func (app *application) interpretTokens(
	tokens instructions.Tokens,
) error {
	list := tokens.List()
	for _, oneToken := range list {
		err := app.interpretToken(
			tokens,
			oneToken,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) interpretToken(
	currentTokens instructions.Tokens,
	token instructions.Token,
) error {
	elements := token.Elements()
	return app.interpretElements(
		currentTokens,
		elements,
	)
}

func (app *application) interpretElements(
	currentTokens instructions.Tokens,
	elements instructions.Elements,
) error {
	list := elements.List()
	for _, oneElement := range list {
		err := app.interpretElement(
			currentTokens,
			oneElement,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) interpretElement(
	currentTokens instructions.Tokens,
	element instructions.Element,
) error {
	if element.IsRule() {
		return nil
	}

	instruction := element.Instruction()
	return app.interpretInstruction(
		instruction,
	)
}

func (app *application) interpretSyscall(
	currentTokens instructions.Tokens,
	sysCall instructions.Syscall,
) error {
	fnName := sysCall.FuncName()
	mpParams := map[string][]byte{}
	if sysCall.HasParameters() {
		parameters := sysCall.Parameters()
		retMapParams, err := app.fetchParameters(
			currentTokens,
			parameters,
		)

		if err != nil {
			str := fmt.Sprintf("there was an error while fetching the syscall (sysCallFn: %s) parameters: %s", fnName, err.Error())
			return errors.New(str)
		}

		mpParams = retMapParams
	}

	if fn, ok := app.syscalls[fnName]; ok {
		err := fn(mpParams)
		if err != nil {
			return err
		}
	}

	str := fmt.Sprintf("the sysCall (sysCallFn: %s) does not exists", fnName)
	return errors.New(str)
}

func (app *application) fetchParameters(
	currentTokens instructions.Tokens,
	parameters instructions.Parameters,
) (map[string][]byte, error) {
	output := map[string][]byte{}
	list := parameters.List()
	for _, oneParameter := range list {
		name, value, err := app.fetchParameter(
			currentTokens,
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
	currentTokens instructions.Tokens,
	parameter instructions.Parameter,
) (string, []byte, error) {
	value := parameter.Value()
	retBytes, err := app.fetchValue(
		currentTokens,
		value,
	)

	if err != nil {
		return "", nil, err
	}

	return parameter.Name(), retBytes, nil
}

func (app *application) fetchValue(
	currentTokens instructions.Tokens,
	value instructions.Value,
) ([]byte, error) {
	if value.IsBytes() {
		return value.Bytes(), nil
	}

	reference := value.Reference()
	element := reference.Element()
	index := reference.Index()
	retToken, err := currentTokens.Fetch(element, index)
	if err != nil {
		return nil, err
	}

	elements := retToken.Elements()
	return app.elementsAdapter.ToBytes(elements)
}
