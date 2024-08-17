package grammars

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytesToRuleName_Success(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedValue := []byte("MY_RULE")
	expectedRemainong := []byte("!this is some value")
	input := []byte(fmt.Sprintf(`%s%s`, string(expectedValue), string(expectedRemainong)))
	retName, retRemaining, err := bytesToRuleName(input, possibleCharacters, []byte(ruleNameSeparator)[0])
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedValue, retName) {
		t.Errorf("the expected output was (%s), returned (%s)", expectedValue, retName)
		return
	}

	if !bytes.Equal(expectedRemainong, retRemaining) {
		t.Errorf("the remaining output was (%s), returned (%s)", expectedRemainong, retRemaining)
		return
	}
}

func TestBytesToRuleName_separatorAtEndOfRuleNameIsNotTaken_Success(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedValue := []byte("MY_RULE")
	expectedRemainong := []byte("_this is some value")
	input := []byte(fmt.Sprintf(`%s%s`, string(expectedValue), string(expectedRemainong)))
	retName, retRemaining, err := bytesToRuleName(input, possibleCharacters, []byte(ruleNameSeparator)[0])
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedValue, retName) {
		t.Errorf("the expected output was (%s), returned (%s)", expectedValue, retName)
		return
	}

	if !bytes.Equal(expectedRemainong, retRemaining) {
		t.Errorf("the remaining output was (%s), returned (%s)", expectedRemainong, retRemaining)
		return
	}
}

func TestBytesToRuleName_ruleNameIsOnlyASeparator_returnsError(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedValue := []byte("_")
	expectedRemainong := []byte("this is some value")
	input := []byte(fmt.Sprintf(`%s%s`, string(expectedValue), string(expectedRemainong)))
	_, _, err := bytesToRuleName(input, possibleCharacters, []byte(ruleNameSeparator)[0])
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToRuleName_firstCharacterIsSeparator_returnsError(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedValue := []byte("_MY_RULE")
	expectedRemainong := []byte("!this is some value")
	input := []byte(fmt.Sprintf(`%s%s`, string(expectedValue), string(expectedRemainong)))
	_, _, err := bytesToRuleName(input, possibleCharacters, []byte(ruleNameSeparator)[0])
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
