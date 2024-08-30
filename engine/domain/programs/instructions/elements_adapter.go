package instructions

type elementsAdapter struct {
}

func createElementsAdapter() ElementsAdapter {
	out := elementsAdapter{}
	return &out
}

// ToBytes takes an elements and returns its bytes
func (app *elementsAdapter) ToBytes(elements Elements) ([]byte, error) {
	return app.elementsToBytes(
		elements,
	)
}

func (app *elementsAdapter) instructionToBytes(
	instruction Instruction,
) ([]byte, error) {
	tokens := instruction.Tokens()
	return app.tokensToBytes(
		tokens,
	)
}

func (app *elementsAdapter) tokensToBytes(
	tokens Tokens,
) ([]byte, error) {
	output := []byte{}
	list := tokens.List()
	for _, oneToken := range list {
		retBytes, err := app.tokenToBytes(
			oneToken,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

func (app *elementsAdapter) tokenToBytes(
	token Token,
) ([]byte, error) {
	elements := token.Elements()
	return app.elementsToBytes(
		elements,
	)
}

func (app *elementsAdapter) elementsToBytes(
	elements Elements,
) ([]byte, error) {
	output := []byte{}
	list := elements.List()
	for _, oneElement := range list {
		retBytes, err := app.elementToBytes(
			oneElement,
		)

		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

func (app *elementsAdapter) elementToBytes(
	element Element,
) ([]byte, error) {
	if element.IsRule() {
		rule := element.Rule()
		return rule.Bytes(), nil
	}

	instruction := element.Instruction()
	return app.instructionToBytes(
		instruction,
	)
}
