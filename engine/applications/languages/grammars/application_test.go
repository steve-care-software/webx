package grammars

import (
	"bytes"
	"testing"
)

func TestApplication_rule_Success(t *testing.T) {
	expectedName := "MY_RULE"
	expectedValue := []byte(`this " with escape`)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`MY_RULE:"this \" with escape"this is some remaining`)

	application := NewApplication().(*application)
	retRule, retRemaining, err := application.bytesToRule(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retRule.Name() != expectedName {
		t.Errorf("the name was expected to be %s, %s returned", expectedName, retRule.Name())
		return
	}

	if !bytes.Equal(expectedValue, retRule.Bytes()) {
		t.Errorf("the expected value was (%s), returned (%s)", expectedValue, retRule.Bytes())
		return
	}
}

func TestApplication_rule_withInvalidName_returnsError(t *testing.T) {
	input := []byte(`_MY_RULE:"this \" with escape"this is some remaining`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_rule_withoutValue_returnsError(t *testing.T) {
	input := []byte(`MY_RULE:""this is some remaining`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_cardinality_withoutMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[1,]this is some remaining`)

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestApplication_cardinality_withMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedMax := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[1,1]this is some remaining`)

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if !retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pRetMax := retCardinality.Max()
	if *pRetMax != expectedMax {
		t.Errorf("the max was expected to be %d, %d returned", expectedMax, *pRetMax)
		return
	}
}

func TestApplication_cardinality_withZeroPlus_Success(t *testing.T) {
	expectedMin := uint(0)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`*this is some remaining`)

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestApplication_cardinality_withOnePlus_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`+this is some remaining`)

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestApplication_cardinality_withInvalidInput_returnsError(t *testing.T) {
	input := []byte(`this is some invalid input`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToCardinality(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}