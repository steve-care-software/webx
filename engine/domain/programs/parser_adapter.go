package programs

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
)

type parserAdapter struct {
	grammarAdapter      grammars.ParserAdapter
	builder             Builder
	instructionsBuilder instructions.Builder
	instructionBuilder  instructions.InstructionBuilder
	tokensBuilder       instructions.TokensBuilder
	tokenBuilder        instructions.TokenBuilder
	elementsBuilder     instructions.ElementsBuilder
	elementBuilder      instructions.ElementBuilder
	syscallsBuilder     instructions.SyscallsBuilder
	syscallBuilder      instructions.SyscallBuilder
	parametersBuilder   instructions.ParametersBuilder
	parameterBuilder    instructions.ParameterBuilder
}

func createParserAdapter(
	grammarAdapter grammars.ParserAdapter,
	builder Builder,
	instructionsBuilder instructions.Builder,
	instructionBuilder instructions.InstructionBuilder,
	tokensBuilder instructions.TokensBuilder,
	tokenBuilder instructions.TokenBuilder,
	elementsBuilder instructions.ElementsBuilder,
	elementBuilder instructions.ElementBuilder,
	syscallsBuilder instructions.SyscallsBuilder,
	syscallBuilder instructions.SyscallBuilder,
	parametersBuilder instructions.ParametersBuilder,
	parameterBuilder instructions.ParameterBuilder,
) ParserAdapter {
	out := parserAdapter{
		grammarAdapter:      grammarAdapter,
		builder:             builder,
		instructionsBuilder: instructionsBuilder,
		instructionBuilder:  instructionBuilder,
		tokensBuilder:       tokensBuilder,
		tokenBuilder:        tokenBuilder,
		elementsBuilder:     elementsBuilder,
		elementBuilder:      elementBuilder,
		syscallsBuilder:     syscallsBuilder,
		syscallBuilder:      syscallBuilder,
		parametersBuilder:   parametersBuilder,
		parameterBuilder:    parameterBuilder,
	}

	return &out
}

// ToProgram takes the grammar and input and converts them to a program instance and the remaining data
func (app *parserAdapter) ToProgram(grammar grammars.Grammar, input []byte) (Program, []byte, error) {
	root := grammar.Root()
	retElement, retRemaining, err := app.toElement(grammar, root, input)
	if err != nil {
		return nil, nil, err
	}

	program, err := app.builder.Create().
		WithGrammar(grammar).
		WithRoot(retElement).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return program, retRemaining, nil
}

func (app *parserAdapter) toInstruction(
	grammar grammars.Grammar,
	block blocks.Block,
	input []byte,
) (instructions.Instruction, []byte, error) {
	name := block.Name()
	if block.HasLine() {
		line := block.Line()
		retTokens, retRemaining, err := app.toTokens(
			grammar,
			line,
			input,
		)

		if err != nil {
			return nil, nil, err
		}

		retIns, err := app.instructionBuilder.Create().
			WithBlock(name).
			WithLine(uint(0)).
			WithTokens(retTokens).
			Now()

		if err != nil {
			return nil, nil, err
		}

		return retIns, retRemaining, nil
	}

	lines := block.Lines().List()
	for idx, oneLine := range lines {
		retTokens, retRemaining, err := app.toTokens(
			grammar,
			oneLine,
			input,
		)

		if err != nil {
			continue
		}

		retIns, err := app.instructionBuilder.Create().
			WithBlock(name).
			WithLine(uint(idx)).
			WithTokens(retTokens).
			Now()

		if err != nil {
			return nil, nil, err
		}

		return retIns, retRemaining, nil

	}

	str := fmt.Sprintf("the provided input could not match any line of the block (name: %s)", name)
	return nil, nil, errors.New(str)
}

func (app *parserAdapter) toTokens(
	grammar grammars.Grammar,
	line lines.Line,
	input []byte,
) (instructions.Tokens, []byte, error) {
	output := []instructions.Token{}
	remaining := app.filterOmissions(
		grammar,
		input,
	)

	list := line.Tokens().List()
	for idx, oneToken := range list {
		retToken, retRemaining, err := app.toToken(
			grammar,
			oneToken,
			remaining,
		)

		if err != nil {
			name := oneToken.Name()
			str := fmt.Sprintf("the token (name: %s, index: %d) could not be matched using the provided input", name, idx)
			return nil, nil, errors.New(str)
		}

		output = append(output, retToken)
		remaining = retRemaining
	}

	retTokens, err := app.tokensBuilder.Create().WithList(output).Now()
	if err != nil {
		return nil, nil, err
	}

	return retTokens, remaining, nil
}

func (app *parserAdapter) toToken(
	grammar grammars.Grammar,
	token tokens.Token,
	input []byte,
) (instructions.Token, []byte, error) {
	remaining := app.filterOmissions(
		grammar,
		input,
	)

	cardinality := token.Cardinality()
	hasMax := cardinality.HasMax()
	pMax := cardinality.Max()
	elementsList := []instructions.Element{}
	cpt := uint(0)
	for {
		// max has been reached
		if hasMax {
			max := *pMax
			if cpt >= max {
				break
			}
		}

		retBytes := app.filterOmissions(
			grammar,
			remaining,
		)

		element := token.Element()
		retElement, retRemaining, err := app.toElement(
			grammar,
			element,
			retBytes,
		)

		if err != nil {
			break
		}

		elementsList = append(elementsList, retElement)
		remaining = retRemaining
		cpt++
	}

	min := cardinality.Min()
	length := uint(len(elementsList))
	if length < min {
		str := fmt.Sprintf("the token was expected a minimum of %d elements, %d returned", min, length)
		return nil, nil, errors.New(str)
	}

	elements, err := app.elementsBuilder.Create().WithList(elementsList).Now()
	if err != nil {
		return nil, nil, err
	}

	name := token.Name()
	retToken, err := app.tokenBuilder.Create().WithName(name).WithElements(elements).Now()
	if err != nil {
		return nil, nil, err
	}

	return retToken, app.filterOmissions(grammar, remaining), nil
}

func (app *parserAdapter) toElement(
	grammar grammars.Grammar,
	element elements.Element,
	input []byte,
) (instructions.Element, []byte, error) {
	remaining := input
	builder := app.elementBuilder.Create()
	if element.IsRule() {
		ruleName := element.Rule()
		rule, err := grammar.Rules().Fetch(ruleName)
		if err != nil {
			return nil, nil, err
		}

		ruleBytes := rule.Bytes()
		if !bytes.HasPrefix(input, ruleBytes) {
			str := fmt.Sprintf("the rule (name: %s) could not be found in the input bytes", ruleName)
			return nil, nil, errors.New(str)
		}

		builder.WithRule(rule)
		remaining = input[len(ruleBytes):]
	}

	if element.IsBlock() {
		blockName := element.Block()
		block, err := grammar.Blocks().Fetch(blockName)
		if err != nil {
			return nil, nil, err
		}

		retInstruction, retInstructionRemaining, err := app.toInstruction(
			grammar,
			block,
			input,
		)

		if err != nil {
			return nil, nil, err
		}

		builder.WithInstruction(retInstruction)
		remaining = retInstructionRemaining
	}

	if element.IsSyscall() {
		syscallName := element.Syscall()
		sysCall, err := grammar.Syscalls().Fetch(syscallName)
		if err != nil {
			return nil, nil, err
		}

		retSysCall, retRemaining, err := app.toSyscall(
			grammar,
			sysCall,
			remaining,
		)
		if err != nil {
			return nil, nil, err
		}

		builder.WithSyscall(retSysCall)
		remaining = retRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) toSyscall(
	grammar grammars.Grammar,
	syscall syscalls.Syscall,
	input []byte,
) (instructions.Syscall, []byte, error) {
	remaining := input
	name := syscall.Name()
	funcName := syscall.FuncName()
	builder := app.syscallBuilder.Create().
		WithFuncName(funcName).
		WithName(name)

	if syscall.HasParameters() {
		parameters := syscall.Parameters()
		retParameters, retRemaining, err := app.toParameters(
			grammar,
			parameters,
			remaining,
		)

		if err != nil {
			return nil, nil, err
		}

		builder.WithParameters(retParameters)
		remaining = retRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) toParameters(
	grammar grammars.Grammar,
	parameters parameters.Parameters,
	input []byte,
) (instructions.Parameters, []byte, error) {
	remaining := input
	list := parameters.List()
	output := []instructions.Parameter{}
	for _, oneParameter := range list {
		retParameter, retRemaining, err := app.toParameter(
			grammar,
			oneParameter,
			remaining,
		)

		if err != nil {
			return nil, nil, err
		}

		output = append(output, retParameter)
		remaining = retRemaining
	}

	ins, err := app.parametersBuilder.Create().
		WithList(output).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) toParameter(
	grammar grammars.Grammar,
	parameter parameters.Parameter,
	input []byte,
) (instructions.Parameter, []byte, error) {
	element := parameter.Element()
	name := parameter.Name()
	index := parameter.Index()
	ins, err := app.parameterBuilder.Create().
		WithElement(element.Name()).
		WithName(name).
		WithIndex(index).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, input, nil
}

func (app *parserAdapter) filterOmissions(
	grammar grammars.Grammar,
	input []byte,
) []byte {
	if !grammar.HasOmissions() {
		return input
	}

	omissionsList := grammar.Omissions().List()
	for _, oneOmission := range omissionsList {
		_, retRemaining, err := app.toElement(
			grammar,
			oneOmission,
			input,
		)

		if err != nil {
			continue
		}

		return app.filterOmissions(
			grammar,
			retRemaining,
		)
	}

	return input
}
