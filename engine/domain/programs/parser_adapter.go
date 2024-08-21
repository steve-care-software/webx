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
	instructions_tokens "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	instructions_elements "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	instructions_syscalls "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls"
	instructions_syscalls_values_parameters "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls/parameters"
)

type parserAdapter struct {
	grammarAdapter       grammars.ParserAdapter
	builder              Builder
	instructionsBuilder  instructions.Builder
	instructionBuilder   instructions.InstructionBuilder
	tokensBuilder        instructions_tokens.Builder
	tokenBuilder         instructions_tokens.TokenBuilder
	elementsBuilder      instructions_elements.Builder
	elementBuilder       instructions_elements.ElementBuilder
	syscallsBuilder      instructions_syscalls.Builder
	syscallBuilder       instructions_syscalls.SyscallBuilder
	parametersBuilder    instructions_syscalls_values_parameters.Builder
	parameterBuilder     instructions_syscalls_values_parameters.ParameterBuilder
	currrentInstructions []instructions.Instruction
}

func createParserAdapter(
	grammarAdapter grammars.ParserAdapter,
	builder Builder,
	instructionsBuilder instructions.Builder,
	instructionBuilder instructions.InstructionBuilder,
	tokensBuilder instructions_tokens.Builder,
	tokenBuilder instructions_tokens.TokenBuilder,
	elementsBuilder instructions_elements.Builder,
	elementBuilder instructions_elements.ElementBuilder,
	syscallsBuilder instructions_syscalls.Builder,
	syscallBuilder instructions_syscalls.SyscallBuilder,
	parametersBuilder instructions_syscalls_values_parameters.Builder,
	parameterBuilder instructions_syscalls_values_parameters.ParameterBuilder,
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

	out.init()
	return &out
}

func (app *parserAdapter) init() {
	app.currrentInstructions = []instructions.Instruction{}
}

// ToProgram takes the grammar and input and converts them to a program instance and the remaining data
func (app *parserAdapter) ToProgram(grammar grammars.Grammar, input []byte) (Program, []byte, error) {
	root := grammar.Root()
	retElement, retRemaining, err := app.toElement(grammar, root, input)
	if err != nil {
		return nil, nil, err
	}

	instructions, err := app.instructionsBuilder.Create().
		WithList(app.currrentInstructions).
		Now()

	if err != nil {
		return nil, nil, err
	}

	program, err := app.builder.Create().
		WithGrammar(grammar).
		WithInstructions(instructions).
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
) (instructions_tokens.Tokens, []byte, error) {
	output := []instructions_tokens.Token{}
	remaining := input
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
) (instructions_tokens.Token, []byte, error) {
	remaining := input
	cardinality := token.Cardinality()
	hasMax := cardinality.HasMax()
	pMax := cardinality.Max()
	elementsList := []instructions_elements.Element{}
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

	return retToken, remaining, nil
}

func (app *parserAdapter) toElement(
	grammar grammars.Grammar,
	element elements.Element,
	input []byte,
) (instructions_elements.Element, []byte, error) {
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

		builder.WithRule(ruleName)
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

		name := retInstruction.Block()
		builder.WithInstruction(name)
		remaining = retInstructionRemaining

		// add instruction to the list:
		app.currrentInstructions = append(app.currrentInstructions, retInstruction)
	}

	if element.IsSyscall() {
		syscallName := element.Syscall()
		sysCall, err := grammar.Syscalls().Fetch(syscallName)
		if err != nil {
			return nil, nil, err
		}

		retSysCall, err := app.toSyscall(
			sysCall,
		)
		if err != nil {
			return nil, nil, err
		}

		builder.WithSyscall(retSysCall)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) toSyscall(
	syscall syscalls.Syscall,
) (instructions_syscalls.Syscall, error) {
	name := syscall.Name()
	funcName := syscall.FuncName()
	builder := app.syscallBuilder.Create().WithFuncName(funcName).WithName(name)
	if syscall.HasParameters() {
		parameters := syscall.Parameters()
		retParameters, err := app.toParameters(
			parameters,
		)

		if err != nil {
			return nil, err
		}

		builder.WithParameters(retParameters)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (app *parserAdapter) toParameters(
	parameters parameters.Parameters,
) (instructions_syscalls_values_parameters.Parameters, error) {
	list := parameters.List()
	output := []instructions_syscalls_values_parameters.Parameter{}
	for _, oneParameter := range list {
		element := oneParameter.Element()
		name := oneParameter.Name()
		index := oneParameter.Index()
		ins, err := app.parameterBuilder.Create().
			WithToken(element.Name()).
			WithName(name).
			WithIndex(index).
			Now()

		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.parametersBuilder.Create().
		WithList(output).
		Now()
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
