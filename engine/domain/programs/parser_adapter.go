package programs

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls/values"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	instructions_tokens "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	instructions_elements "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	instructions_syscalls "github.com/steve-care-software/webx/engine/domain/programs/syscalls"
)

type parserAdapter struct {
	grammarAdapter      grammars.ParserAdapter
	builder             Builder
	instructionsBuilder instructions.Builder
	instructionBuilder  instructions.InstructionBuilder
	tokensBuilder       instructions_tokens.Builder
	tokenBuilder        instructions_tokens.TokenBuilder
	elementsBuilder     instructions_elements.Builder
	elementBuilder      instructions_elements.ElementBuilder
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
	}

	return &out
}

// ToProgram takes the grammar and input and converts them to a program instance and the remaining data
func (app *parserAdapter) ToProgram(grammar grammars.Grammar, input []byte) (Program, []byte, error) {
	root := grammar.Root()
	retElement, retInstList, retRemaining, err := app.toElement(grammar, root, []instructions.Instruction{}, input)
	if err != nil {
		return nil, nil, err
	}

	instructions, err := app.instructionsBuilder.Create().
		WithList(retInstList).
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

// ToBytes takes a program and returns the bytes for the grammar and program
func (app *parserAdapter) ToBytes(program Program) ([]byte, []byte, error) {
	return nil, nil, nil
}

func (app *parserAdapter) toInstruction(
	grammar grammars.Grammar,
	block blocks.Block,
	currentInsList []instructions.Instruction,
	input []byte,
) (instructions.Instruction, []instructions.Instruction, []byte, error) {
	name := block.Name()
	if block.HasLine() {
		line := block.Line()
		retTokens, retInsList, retRemaining, err := app.toTokens(
			grammar,
			line,
			currentInsList,
			input,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		retIns, err := app.instructionBuilder.Create().
			WithBlock(name).
			WithLine(uint(0)).
			WithTokens(retTokens).
			Now()

		if err != nil {
			return nil, nil, nil, err
		}

		return retIns, retInsList, retRemaining, nil
	}

	lines := block.Lines().List()
	for idx, oneLine := range lines {
		retTokens, retInsList, retRemaining, err := app.toTokens(
			grammar,
			oneLine,
			currentInsList,
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
			return nil, nil, nil, err
		}

		return retIns, retInsList, retRemaining, nil
	}

	str := fmt.Sprintf("the provided input could not match any line of the block (name: %s)", name)
	return nil, nil, nil, errors.New(str)
}

func (app *parserAdapter) toTokens(
	grammar grammars.Grammar,
	line lines.Line,
	currentInsList []instructions.Instruction,
	input []byte,
) (instructions_tokens.Tokens, []instructions.Instruction, []byte, error) {
	output := []instructions_tokens.Token{}
	remaining := input
	retInsList := currentInsList
	list := line.Tokens().List()
	for idx, oneToken := range list {
		retToken, retInsListAfterToken, retRemaining, err := app.toToken(
			grammar,
			oneToken,
			retInsList,
			remaining,
		)

		if err != nil {
			name := oneToken.Name()
			str := fmt.Sprintf("the token (name: %s, index: %d) could not be matched using the provided input", name, idx)
			return nil, nil, nil, errors.New(str)
		}

		output = append(output, retToken)
		retInsList = retInsListAfterToken
		remaining = retRemaining
	}

	retTokens, err := app.tokensBuilder.Create().WithList(output).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return retTokens, retInsList, remaining, nil
}

func (app *parserAdapter) toToken(
	grammar grammars.Grammar,
	token tokens.Token,
	currentInsList []instructions.Instruction,
	input []byte,
) (instructions_tokens.Token, []instructions.Instruction, []byte, error) {
	remaining := input
	retInsList := currentInsList
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
			currentInsList,
			remaining,
		)

		element := token.Element()
		retElement, retInsListAfterElement, retRemaining, err := app.toElement(
			grammar,
			element,
			retInsList,
			retBytes,
		)

		if err != nil {
			break
		}

		elementsList = append(elementsList, retElement)
		retInsList = retInsListAfterElement
		remaining = retRemaining
		cpt++
	}

	min := cardinality.Min()
	length := uint(len(elementsList))
	if length < min {
		str := fmt.Sprintf("the token was expected a minimum of %d elements, %d returned", min, length)
		return nil, nil, nil, errors.New(str)
	}

	elements, err := app.elementsBuilder.Create().WithList(elementsList).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	name := token.Name()
	retToken, err := app.tokenBuilder.Create().WithName(name).WithElements(elements).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return retToken, retInsList, remaining, nil
}

func (app *parserAdapter) filterOmissions(
	grammar grammars.Grammar,
	currentInsList []instructions.Instruction,
	input []byte,
) []byte {
	if !grammar.HasOmissions() {
		return input
	}

	omissionsList := grammar.Omissions().List()
	for _, oneOmission := range omissionsList {
		_, _, retRemaining, err := app.toElement(
			grammar,
			oneOmission,
			currentInsList,
			input,
		)

		if err != nil {
			continue
		}

		return app.filterOmissions(
			grammar,
			currentInsList,
			retRemaining,
		)
	}

	return input
}

func (app *parserAdapter) toElement(
	grammar grammars.Grammar,
	element elements.Element,
	currentInsList []instructions.Instruction,
	input []byte,
) (instructions_elements.Element, []instructions.Instruction, []byte, error) {
	remaining := input
	retInsList := currentInsList
	builder := app.elementBuilder.Create()
	if element.IsRule() {
		ruleName := element.Rule()
		rule, err := grammar.Rules().Fetch(ruleName)
		if err != nil {
			return nil, nil, nil, err
		}

		ruleBytes := rule.Bytes()
		if !bytes.HasPrefix(input, ruleBytes) {
			str := fmt.Sprintf("the rule (name: %s) could not be found in the input bytes", ruleName)
			return nil, nil, nil, errors.New(str)
		}

		builder.WithRule(ruleName)
		remaining = input[len(ruleBytes):]
	}

	if element.IsBlock() {
		blockName := element.Block()
		block, err := grammar.Blocks().Fetch(blockName)
		if err != nil {
			return nil, nil, nil, err
		}

		retInstruction, retInsListAfterIns, retInstructionRemaining, err := app.toInstruction(grammar, block, currentInsList, input)
		if err != nil {
			return nil, nil, nil, err
		}

		name := retInstruction.Block()
		builder.WithInstruction(name)
		retInsList = append(retInsList, retInsListAfterIns...)
		retInsList = append(retInsList, retInstruction)
		remaining = retInstructionRemaining
	}

	if element.IsSyscall() {
		/*syscall := element.Syscall()
		fmt.Printf("\n%s\n", syscall)
		fmt.Printf("\n%s\n", input)
		syscallBytes := []byte(syscall)
		if !bytes.HasPrefix(input, syscallBytes) {
			str := fmt.Sprintf("the syscall (name: %s) could not be found in the input bytes", syscall)
			return nil, nil, nil, errors.New(str)
		}

		builder.WithSyscall(syscall)
		remaining = input[len(syscallBytes):]*/
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return ins, retInsList, remaining, nil
}

func (app *parserAdapter) toSyscall(
	grammar grammars.Grammar,
	syscall syscalls.Syscall,
	currentInsList []instructions.Instruction,
	input []byte,
) (instructions_syscalls.Syscall, []instructions.Instruction, []byte, error) {
	return nil, nil, nil, nil
}

func (app *parserAdapter) toValues(
	grammar grammars.Grammar,
	values values.Values,
	currentInsList []instructions.Instruction,
	input []byte,
) ([]instructions.Instruction, []byte, error) {
	return nil, nil, nil
}
