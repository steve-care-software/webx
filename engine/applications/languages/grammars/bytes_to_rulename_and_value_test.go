package grammars

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytesToRuleNameAndValue_Success(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedName := []byte("MY_RULE")
	expectedValue := []byte(`this " with escape`)
	expectedRemaining := []byte("!this is some value")

	input := []byte(
		fmt.Sprintf(
			`%s%s%s%s%s%s`,
			string(expectedName),
			string([]byte(ruleNameValueSeparator)[0]),
			string([]byte(ruleValuePrefix)[0]),
			string(`this \" with escape`),
			string([]byte(ruleValueSuffix)[0]),
			string(expectedRemaining),
		),
	)

	retName, retValue, retRemaining, err := bytesToRuleNameAndValue(
		input,
		[]byte(ruleNameValueSeparator)[0],
		possibleCharacters,
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(filterBytes),
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedName, retName) {
		t.Errorf("the expected name was (%s), returned (%s)", expectedName, retName)
		return
	}

	if !bytes.Equal(expectedValue, retValue) {
		t.Errorf("the expected value was (%s), returned (%s)", expectedValue, retValue)
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}
}

func TestBytesToRuleNameAndValue_withoutSeparator_returnsError(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedName := []byte("MY_RULE")
	expectedRemaining := []byte("!this is some value")

	input := []byte(
		fmt.Sprintf(
			`%s%s%s%s%s`,
			string(expectedName),
			string([]byte(ruleValuePrefix)[0]),
			string(`this \" with escape`),
			string([]byte(ruleValueSuffix)[0]),
			string(expectedRemaining),
		),
	)

	_, _, _, err := bytesToRuleNameAndValue(
		input,
		[]byte(ruleNameValueSeparator)[0],
		possibleCharacters,
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToRuleNameAndValue_withInvalidName_returnsError(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedName := []byte("_MY_RULE")
	expectedRemaining := []byte("!this is some value")

	input := []byte(
		fmt.Sprintf(
			`%s%s%s%s%s%s`,
			string(expectedName),
			string([]byte(ruleNameValueSeparator)[0]),
			string([]byte(ruleValuePrefix)[0]),
			string(`this \" with escape`),
			string([]byte(ruleValueSuffix)[0]),
			string(expectedRemaining),
		),
	)

	_, _, _, err := bytesToRuleNameAndValue(
		input,
		[]byte(ruleNameValueSeparator)[0],
		possibleCharacters,
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToRuleNameAndValue_withoutRemainingAfterName_returnsError(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedName := []byte("MY_RULE")

	input := []byte(string(expectedName))
	_, _, _, err := bytesToRuleNameAndValue(
		input,
		[]byte(ruleNameValueSeparator)[0],
		possibleCharacters,
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToRuleNameAndValue_withoutValueSuffix_returnsError(t *testing.T) {
	possibleCharacters := createPossibleUpperCaseLetters()
	expectedName := []byte("MY_RULE")
	expectedRemaining := []byte("!this is some value")

	input := []byte(
		fmt.Sprintf(
			`%s%s%s%s%s`,
			string(expectedName),
			string([]byte(ruleNameValueSeparator)[0]),
			string([]byte(ruleValuePrefix)[0]),
			string(`this \" with escape`),
			string(expectedRemaining),
		),
	)

	_, _, _, err := bytesToRuleNameAndValue(
		input,
		[]byte(ruleNameValueSeparator)[0],
		possibleCharacters,
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
