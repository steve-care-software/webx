package programs

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
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
	ruleBuilder         rules.RuleBuilder
	syscallBuilder      instructions.SyscallBuilder
	parametersBuilder   instructions.ParametersBuilder
	parameterBuilder    instructions.ParameterBuilder
	valueBuilder        instructions.ValueBuilder
	referenceBuilder    instructions.ReferenceBuilder
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
	ruleBuilder rules.RuleBuilder,
	syscallBuilder instructions.SyscallBuilder,
	parametersBuilder instructions.ParametersBuilder,
	parameterBuilder instructions.ParameterBuilder,
	valueBuilder instructions.ValueBuilder,
	referenceBuilder instructions.ReferenceBuilder,
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
		ruleBuilder:         ruleBuilder,
		syscallBuilder:      syscallBuilder,
		parametersBuilder:   parametersBuilder,
		parameterBuilder:    parameterBuilder,
		valueBuilder:        valueBuilder,
		referenceBuilder:    referenceBuilder,
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

// ToProgramWithRoot creates a program but changes the root block of the grammar
func (app *parserAdapter) ToProgramWithRoot(grammar grammars.Grammar, rootBlockName string, input []byte) (Program, []byte, error) {
	rootBlock, err := grammar.Blocks().Fetch(rootBlockName)
	if err != nil {
		return nil, nil, err
	}

	retInstruction, retInstructionRemaining, err := app.toInstruction(
		grammar,
		rootBlock,
		input,
	)

	if err != nil {
		return nil, nil, err
	}

	element, err := app.elementBuilder.Create().
		WithInstruction(retInstruction).
		Now()

	if err != nil {
		return nil, nil, err
	}

	program, err := app.builder.Create().
		WithGrammar(grammar).
		WithRoot(element).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return program, retInstructionRemaining, nil
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

		builder := app.instructionBuilder.Create().
			WithBlock(name).
			WithLine(uint(0)).
			WithTokens(retTokens)

		if line.HasSyscall() {
			syscall := line.Syscall()
			retSyscall, err := app.toSysCall(syscall)
			if err != nil {
				return nil, nil, err
			}

			builder.WithSyscall(retSyscall)
		}

		retIns, err := builder.Now()
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

		builder := app.instructionBuilder.Create().
			WithBlock(name).
			WithLine(uint(idx)).
			WithTokens(retTokens)

		if oneLine.HasSyscall() {
			syscall := oneLine.Syscall()
			retSyscall, err := app.toSysCall(syscall)
			if err != nil {
				return nil, nil, err
			}

			builder.WithSyscall(retSyscall)
		}

		retIns, err := builder.Now()
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

		if len(remaining) <= 0 {
			break
		}

		retBytes := app.filterOmissions(
			grammar,
			remaining,
		)

		element := token.Element()
		if token.HasReverse() {
			isEscaped := false
			reverse := token.Reverse()
			retRemaining := retBytes
			accumulated := []byte{}
			for _, oneByte := range retBytes {
				if reverse.HasEscape() {
					escapeElement := reverse.Escape()
					_, retRemainingAfterEscape, err := app.toElement(
						grammar,
						escapeElement,
						retRemaining,
					)

					if err == nil {
						retRemaining = retRemainingAfterEscape
						isEscaped = true
						continue
					}
				}

				_, retRemainingAfterElement, err := app.toElement(
					grammar,
					element,
					retRemaining,
				)

				if isEscaped || err != nil {
					accumulated = append(accumulated, oneByte)
				}

				if err != nil {
					// previous character was escape but the next one did not match the element, so reset the escape:
					if isEscaped {
						isEscaped = false
						continue
					}

					retRemaining = retRemaining[1:]
					continue
				}

				// we escape the character so continue and reset it:
				if isEscaped {
					isEscaped = false
					retRemaining = retRemainingAfterElement
					continue
				}

				break
			}

			name := token.Name()
			rule, err := app.ruleBuilder.Create().
				WithBytes(accumulated).
				WithName(name).
				Now()

			if err != nil {
				return nil, nil, err
			}

			retElement, err := app.elementBuilder.Create().
				WithRule(rule).
				Now()

			if err != nil {
				return nil, nil, err
			}

			elementsList = append(elementsList, retElement)
			remaining = retRemaining
			cpt++
			continue
		}

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

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *parserAdapter) toSysCall(
	execution executions.Execution,
) (instructions.Syscall, error) {
	funcName := execution.FuncName()
	builder := app.syscallBuilder.Create().
		WithFuncName(funcName)

	if execution.HasParameters() {
		parameters := execution.Parameters()
		retParameters, err := app.toParameters(
			parameters,
		)

		if err != nil {
			return nil, err
		}

		builder.WithParameters(retParameters)
	}

	return builder.Now()
}

func (app *parserAdapter) toParameters(
	parameters parameters.Parameters,
) (instructions.Parameters, error) {
	list := parameters.List()
	output := []instructions.Parameter{}
	for _, oneParameter := range list {
		retParameter, err := app.toParameter(
			oneParameter,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retParameter)
	}

	ins, err := app.parametersBuilder.Create().
		WithList(output).
		Now()

	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (app *parserAdapter) toParameter(
	parameter parameters.Parameter,
) (instructions.Parameter, error) {
	value := parameter.Value()
	retValue, err := app.toValue(value)
	if err != nil {
		return nil, err
	}

	name := parameter.Name()
	return app.parameterBuilder.Create().
		WithName(name).
		WithValue(retValue).
		Now()
}

func (app *parserAdapter) toValue(
	value values.Value,
) (instructions.Value, error) {
	builder := app.valueBuilder.Create()
	if value.IsBytes() {
		bytes := value.Bytes()
		builder.WithBytes(bytes)
	}

	if value.IsReference() {
		reference := value.Reference()
		element := reference.Element()
		index := reference.Index()
		retRef, err := app.referenceBuilder.Create().
			WithElement(element.Name()).
			WithIndex(index).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithReference(retRef)
	}

	return builder.Now()
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
