package grammars

import (
	"bytes"
	"testing"
)

func TestBlockName_Success(t *testing.T) {
	lowerCaseLetters := createPossibleLowerCaseLetters()
	anyLetters := createBlockNameCharacters()
	expectedRemaining := []byte{}
	expectedBlockName := []byte("myBlockName")
	input := []byte(append(expectedBlockName, expectedRemaining...))
	retBlockName, retRemaining, err := blockName(input, lowerCaseLetters, anyLetters, []byte(filterBytes))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedBlockName, retBlockName) {
		t.Errorf("the returned block name is invalid, expected (%s), returned (%s)", string(expectedBlockName), string(retBlockName))
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid, expected (%s), returned (%s)", string(expectedRemaining), string(retRemaining))
		return
	}
}

func TestBlockName_withRemaining_Success(t *testing.T) {
	lowerCaseLetters := createPossibleLowerCaseLetters()
	anyLetters := createBlockNameCharacters()
	expectedRemaining := []byte("!this is some remaining")
	expectedBlockName := []byte("myBlockName")
	input := []byte(append(expectedBlockName, expectedRemaining...))
	retBlockName, retRemaining, err := blockName(input, lowerCaseLetters, anyLetters, []byte(filterBytes))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedBlockName, retBlockName) {
		t.Errorf("the returned block name is invalid, expected (%s), returned (%s)", string(expectedBlockName), string(retBlockName))
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid, expected (%s), returned (%s)", string(expectedRemaining), string(retRemaining))
		return
	}
}

func TestBlockName_withoutMatch_returnsError(t *testing.T) {
	lowerCaseLetters := createPossibleLowerCaseLetters()
	anyLetters := createBlockNameCharacters()
	expectedRemaining := []byte("!this is some remaining")
	expectedBlockName := []byte("INVALID")
	input := []byte(append(expectedBlockName, expectedRemaining...))
	_, _, err := blockName(input, lowerCaseLetters, anyLetters, []byte(filterBytes))
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
