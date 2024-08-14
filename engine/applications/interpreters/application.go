package interpreters

import (
	"errors"
	"fmt"
	"strings"
)

type application struct {
	grammar *grammar
}

func createApplication(
	grammar *grammar,
) Application {
	out := application{
		grammar: grammar,
	}

	return &out
}

// Execute executes bytecode
func (app *application) Execute(input []byte) ([]byte, error) {
	isValid, retRemaining, err := app.entry(input)
	if err != nil {
		return nil, err
	}

	if len(retRemaining) > 0 {
		str := fmt.Sprintf("there was %d additional bytes that could not be executed in the script", len(retRemaining))
		return nil, errors.New(str)
	}

	fmt.Printf("\n%t\n%v\n", isValid, retRemaining)
	return nil, nil
}

func (app *application) entry(
	input []byte,
) (bool, []byte, error) {
	if pBlock, ok := app.grammar.blocks[app.grammar.blockEntry]; ok {
		return app.block(
			pBlock,
			input,
		)
	}

	str := fmt.Sprintf("the block (name: %s) referred to in blockEntry does not exists", app.grammar.blockEntry)
	return false, nil, errors.New(str)
}

func (app *application) block(
	block *block,
	input []byte,
) (bool, []byte, error) {
	if len(input) <= 0 {
		str := fmt.Sprintf("the block (name: %s) cannot be executed because it contains 0 bytes in the input", block.name)
		return false, nil, errors.New(str)
	}

	for _, oneLine := range block.lines {
		isValid, retRemaining, err := app.line(
			oneLine,
			input,
		)

		if err != nil {
			continue
		}

		if !isValid {
			continue
		}

		return true, retRemaining, nil
	}

	return false, nil, nil
}

func (app *application) line(
	line *line,
	input []byte,
) (bool, []byte, error) {
	execMap := map[string]string{}
	remaining := input
	for oneVariableName, oneValueName := range line.elements {
		if pElement, ok := app.grammar.elements[oneValueName]; ok {
			isValid, retRemaining, err := app.element(
				pElement,
				remaining,
			)

			if err != nil {
				return false, nil, err
			}

			if !isValid {
				return false, nil, nil
			}

			length := len(remaining) - len(retRemaining)
			strValue, err := app.bytesToString(remaining[0:length])
			if err != nil {
				return false, nil, err
			}

			execMap[oneVariableName] = strValue
			remaining = retRemaining
			continue
		}

		str := fmt.Sprintf("the element (name: %s) does not exists", oneValueName)
		return false, nil, errors.New(str)
	}

	// execute the line, if any:
	if line.execFn != nil {
		err := line.execFn(execMap)
		if err != nil {
			return false, nil, err
		}
	}

	return true, remaining, nil
}

func (app *application) element(
	element *element,
	input []byte,
) (bool, []byte, error) {
	if len(input) <= 0 {
		str := fmt.Sprintf("the element (name: %s) cannot be executed because it contains 0 bytes in the input", element.name)
		return false, nil, errors.New(str)
	}

	if element.token != "" {
		if pToken, ok := app.grammar.tokens[element.token]; ok {
			return app.token(
				pToken,
				input,
			)
		}

		str := fmt.Sprintf("the token (name: %s) does not exists", element.token)
		return false, nil, errors.New(str)
	}

	if element.tokenPointer != "" {
		if pTokenPointer, ok := app.grammar.tokenPointers[element.tokenPointer]; ok {
			return app.tokenPointer(
				pTokenPointer,
				input,
			)
		}

		str := fmt.Sprintf("the tokenPointer (name: %s) does not exists", element.tokenPointer)
		return false, nil, errors.New(str)
	}

	if pBlockPointer, ok := app.grammar.blockPointers[element.blockPointer]; ok {
		return app.blockPointer(
			pBlockPointer,
			input,
		)
	}

	str := fmt.Sprintf("the tokenPointer (name: %s) does not exists", element.tokenPointer)
	return false, nil, errors.New(str)
}

func (app *application) blockPointer(
	blockPointer *blockPointer,
	input []byte,
) (bool, []byte, error) {
	if pBlock, ok := app.grammar.blocks[blockPointer.block]; ok {
		if pCardinality, ok := app.grammar.cardinalities[blockPointer.cardinality]; ok {
			return app.match(
				nil,
				nil,
				pBlock,
				input,
				pCardinality,
			)
		}

		str := fmt.Sprintf("the cardinality (name: %s) does not exists", blockPointer.cardinality)
		return false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the block (name: %s) does not exists", blockPointer.name)
	return false, nil, errors.New(str)
}

func (app *application) token(
	token *token,
	input []byte,
) (bool, []byte, error) {
	if pCardinality, ok := app.grammar.cardinalities[token.cardinality]; ok {
		return app.match(
			token.characters,
			nil,
			nil,
			input,
			pCardinality,
		)
	}

	str := fmt.Sprintf("the cardinality (name: %s) does not exists", token.cardinality)
	return false, nil, errors.New(str)
}

func (app *application) tokenPointer(
	tokenPointer *tokenPointer,
	input []byte,
) (bool, []byte, error) {
	if pToken, ok := app.grammar.tokens[tokenPointer.token]; ok {
		if pCardinality, ok := app.grammar.cardinalities[tokenPointer.cardinality]; ok {
			return app.match(
				nil,
				pToken,
				nil,
				input,
				pCardinality,
			)
		}

		str := fmt.Sprintf("the cardinality (name: %s) does not exists", tokenPointer.cardinality)
		return false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the token (name: %s) does not exists", tokenPointer.name)
	return false, nil, errors.New(str)
}

func (app *application) match(
	possibleCharacters []byte,
	token *token,
	block *block,
	input []byte,
	cardinality *cardinality,
) (bool, []byte, error) {
	remaining := input
	matches := []byte{}
	cpt := uint(0)
	for {

		if len(remaining) <= 0 {
			break
		}

		if cardinality.pAmount != nil {
			expectedAmount := *cardinality.pAmount
			if cpt >= expectedAmount {
				break
			}
		}

		// if there is expected characters:
		if possibleCharacters != nil {
			idxMatch := 0
			isMatch := false
			for idx, oneCharacter := range possibleCharacters {
				if oneCharacter == remaining[0] {
					isMatch = true
					idxMatch = idx
					break
				}
			}

			// no match:
			if !isMatch {
				break
			}

			// match:
			matches = append(matches, possibleCharacters[idxMatch])
			remaining = remaining[1:]
			cpt++
			continue
		}

		// if there is a valid token:
		if token != nil {
			isValid, retRemaining, err := app.token(
				token,
				remaining,
			)

			if err != nil {
				break
			}

			if !isValid {
				break
			}

			remaining = retRemaining
			cpt++
			continue
		}

		// if there is a valid block:
		isValid, retRemaining, err := app.block(
			block,
			remaining,
		)

		if err != nil {
			break
		}

		if !isValid {
			break
		}

		remaining = retRemaining
		cpt++

	}

	// no match
	if cpt < cardinality.min {
		str := fmt.Sprintf("the expected mininmum (%d) was not reached, amount: %d", cardinality.min, cpt)
		return false, nil, errors.New(str)
	}

	return true, remaining, nil
}

func (app *application) bytesToString(input []byte) (string, error) {
	output := []string{}
	for _, oneByte := range input {
		if str, ok := app.grammar.bytesMapping[oneByte]; ok {
			output = append(output, str)
			continue
		}

		str := fmt.Sprintf("the byte (%d) is undefined in the bytes mapping of the provided grammar", oneByte)
		return "", errors.New(str)
	}

	return strings.Join(output, ""), nil
}
