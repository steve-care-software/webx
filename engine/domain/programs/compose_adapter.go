package programs

import (
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
)

type composeAdapter struct {
}

func createComposeAdapter() ComposeAdapter {
	out := composeAdapter{}
	return &out
}

// ToBytes takes a program and an element name and returns its bytes
func (app *composeAdapter) ToBytes(program Program, elementName string) ([]byte, error) {
	retInstruction, err := program.Instructions().Fetch(elementName)
	if err != nil {
		retRule, err := program.Grammar().Rules().Fetch(elementName)
		if err != nil {
			return nil, err
		}

		return retRule.Bytes(), nil
	}

	return app.instructionToBytes(
		program,
		retInstruction,
	)
}

func (app *composeAdapter) instructionToBytes(
	program Program,
	instruction instructions.Instruction,
) ([]byte, error) {
	tokens := instruction.Tokens()
	return app.tokensToBytes(
		program,
		tokens,
	)
}

func (app *composeAdapter) tokensToBytes(
	program Program,
	tokens tokens.Tokens,
) ([]byte, error) {
	output := []byte{}
	list := tokens.List()
	for _, oneToken := range list {
		retBytes, err := app.tokenToBytes(
			program,
			oneToken,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

func (app *composeAdapter) tokenToBytes(
	program Program,
	token tokens.Token,
) ([]byte, error) {
	elements := token.Elements()
	return app.elementsToBytes(
		program,
		elements,
	)
}

func (app *composeAdapter) elementsToBytes(
	program Program,
	elements elements.Elements,
) ([]byte, error) {
	output := []byte{}
	list := elements.List()
	for _, oneElement := range list {
		retBytes, err := app.elementToBytes(
			program,
			oneElement,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

func (app *composeAdapter) elementToBytes(
	program Program,
	element elements.Element,
) ([]byte, error) {
	if element.IsRule() {
		ruleName := element.Rule()
		retRule, err := program.Grammar().Rules().Fetch(ruleName)
		if err != nil {
			return nil, err
		}

		return retRule.Bytes(), nil
	}

	if element.IsSyscall() {
		return []byte{}, nil
	}

	instructionName := element.Instruction()
	retInstruction, err := program.Instructions().Fetch(instructionName)
	if err != nil {
		return nil, err
	}

	return app.instructionToBytes(
		program,
		retInstruction,
	)
}
