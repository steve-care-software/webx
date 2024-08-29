package grammars

import (
	"bytes"
	"fmt"
	"testing"
)

func TestExtractBetween_Success(t *testing.T) {
	expectedValue := []byte("this is a value")
	expectedRemaining := []byte("this is some remaining")
	input := []byte(fmt.Sprintf(`"%s"%s`, string(expectedValue), string(expectedRemaining)))
	escapeByte := []byte(ruleValueEscape)[0]
	retValue, retRemaining, err := extractBetween(input, []byte(ruleValuePrefix)[0], []byte(ruleValueSuffix)[0], &escapeByte)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedValue, retValue) {
		t.Errorf("the expected output was expected to be (%s), returned (%s)", expectedValue, retValue)
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining was expected to be (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}
}

func TestExtractBetween_withoutNotEnoughCharacters_returnsError(t *testing.T) {
	input := []byte(string("\""))
	escapeByte := []byte(ruleValueEscape)[0]
	_, _, err := extractBetween(input, []byte(ruleValuePrefix)[0], []byte(ruleValueSuffix)[0], &escapeByte)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestExtractBetween_withoutPrefix_returnsError(t *testing.T) {
	expectedValue := []byte("this is a value")
	expectedRemaining := []byte("this is some remaining")
	escapeByte := []byte(ruleValueEscape)[0]
	input := []byte(fmt.Sprintf(`%s"%s`, string(expectedValue), string(expectedRemaining)))
	_, _, err := extractBetween(input, []byte(ruleValuePrefix)[0], []byte(ruleValueSuffix)[0], &escapeByte)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestExtractBetween_withoutSuffix_returnsError(t *testing.T) {
	expectedValue := []byte("this is a value")
	expectedRemaining := []byte("this is some remaining")
	escapeByte := []byte(ruleValueEscape)[0]
	input := []byte(fmt.Sprintf(`"%s%s`, string(expectedValue), string(expectedRemaining)))
	_, _, err := extractBetween(input, []byte(ruleValuePrefix)[0], []byte(ruleValueSuffix)[0], &escapeByte)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestExtractBetween_withEscape_Success(t *testing.T) {
	valueWithEscape := []byte(`this \" with escape`)
	expectedValue := []byte(`this " with escape`)
	expectedRemaining := []byte("this is some remaining")
	escapeByte := []byte(ruleValueEscape)[0]
	input := []byte(fmt.Sprintf(`"%s"%s`, string(valueWithEscape), string(expectedRemaining)))
	retValue, retRemaining, err := extractBetween(input, []byte(ruleValuePrefix)[0], []byte(ruleValueSuffix)[0], &escapeByte)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedValue, retValue) {
		t.Errorf("the expected output was expected to be (%s), returned (%s)", expectedValue, retValue)
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining was expected to be (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}
}

func TestExtractBetween_withOnlyEscape_Success(t *testing.T) {
	valueWithEscape := []byte(`\"`)
	expectedValue := []byte(`"`)
	expectedRemaining := []byte("this is some remaining")
	escapeByte := []byte(ruleValueEscape)[0]
	input := []byte(fmt.Sprintf(`"%s"%s`, string(valueWithEscape), string(expectedRemaining)))
	retValue, retRemaining, err := extractBetween(input, []byte(ruleValuePrefix)[0], []byte(ruleValueSuffix)[0], &escapeByte)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedValue, retValue) {
		t.Errorf("the expected output was expected to be (%s), returned (%s)", expectedValue, retValue)
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining was expected to be (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}
}
