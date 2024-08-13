package interpreters

import (
	"errors"
	"fmt"
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
	execMap := map[string][]byte{}
	remaining := input
	for oneVariableName, oneValueName := range line.values {
		if pValue, ok := app.grammar.values[oneValueName]; ok {
			isValid, retRemaining, err := app.value(
				pValue,
				remaining,
			)

			if err != nil {
				return false, nil, err
			}

			if !isValid {
				return false, nil, nil
			}

			length := len(remaining) - len(retRemaining)
			execMap[oneVariableName] = remaining[0:length]
			remaining = retRemaining
			continue
		}

		str := fmt.Sprintf("the value (name: %s) does not exists", oneValueName)
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

func (app *application) value(
	value *value,
	input []byte,
) (bool, []byte, error) {
	if len(input) <= 0 {
		str := fmt.Sprintf("the value (name: %s) cannot be executed because it contains 0 bytes in the input", value.name)
		return false, nil, errors.New(str)
	}

	if value.token != "" {
		if pToken, ok := app.grammar.tokens[value.token]; ok {
			return app.token(
				pToken,
				input,
			)
		}

		str := fmt.Sprintf("the token (name: %s) does not exists", value.token)
		return false, nil, errors.New(str)
	}

	if value.tokenPointer != "" {
		if pTokenPointer, ok := app.grammar.tokenPointers[value.tokenPointer]; ok {
			return app.tokenPointer(
				pTokenPointer,
				input,
			)
		}

		str := fmt.Sprintf("the tokenPointer (name: %s) does not exists", value.tokenPointer)
		return false, nil, errors.New(str)
	}

	if pBlockPointer, ok := app.grammar.blockPointers[value.blockPointer]; ok {
		return app.blockPointer(
			pBlockPointer,
			input,
		)
	}

	str := fmt.Sprintf("the tokenPointer (name: %s) does not exists", value.tokenPointer)
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
