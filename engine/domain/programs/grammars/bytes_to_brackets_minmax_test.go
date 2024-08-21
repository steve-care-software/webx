package grammars

import (
	"bytes"
	"testing"
)

func TestBytesToBracketMinMax_withSpecificValue_Success(t *testing.T) {
	expectedRemaining := []byte("this is some remaining")
	input := append([]byte("[1]"), expectedRemaining...)
	retMin, pRetMax, retRemaining, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retMin != 1 {
		t.Errorf("the min was expected to be %d, %d returned", 1, retMin)
		return
	}

	if pRetMax == nil {
		t.Errorf("the returned max was expected to be valid, nil returned")
		return
	}

	if *pRetMax != 1 {
		t.Errorf("the max was expected to be %d, %d returned", 1, *pRetMax)
		return
	}
}

func TestBytesToBracketMinMax_withMin_withoutMax_Success(t *testing.T) {
	expectedRemaining := []byte("this is some remaining")
	input := append([]byte("[1,]"), expectedRemaining...)
	retMin, pRetMax, retRemaining, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retMin != 1 {
		t.Errorf("the min was expected to be %d, %d returned", 1, retMin)
		return
	}

	if pRetMax != nil {
		t.Errorf("the returned max was expected to be nil, value returned: %d", *pRetMax)
		return
	}
}

func TestBytesToBracketMinMax_withMin_withMax_Success(t *testing.T) {
	expectedRemaining := []byte("this is some remaining")
	input := append([]byte("[1,2]"), expectedRemaining...)
	retMin, pRetMax, retRemaining, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retMin != 1 {
		t.Errorf("the min was expected to be %d, %d returned", 1, retMin)
		return
	}

	if pRetMax == nil {
		t.Errorf("the returned max was expected to be valid, nil returned")
		return
	}

	if *pRetMax != 2 {
		t.Errorf("the max was expected to be %d, %d returned", 2, *pRetMax)
		return
	}
}

func TestBytesToBracketMinMax_withMin_withMax_withRemaining_withoutCloseByte_returnsError(t *testing.T) {
	input := []byte("[1,2 ")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withMin_withMax_withoutCloseByte_returnsError(t *testing.T) {
	input := []byte("[1,2")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withMin_withoutSeparatorByte_returnsError(t *testing.T) {
	input := []byte("[1 ")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withMin_withoutRemainingAfterMin_returnsError(t *testing.T) {
	input := []byte("[1")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withMin_withoutRemainingAfterSeparator_returnsError(t *testing.T) {
	input := []byte("[1,")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withoutOpenByte_returnsError(t *testing.T) {
	input := []byte("111")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withoutMinMatch_returnsError(t *testing.T) {
	input := []byte("[d]")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withEmptyData_returnsError(t *testing.T) {
	input := []byte("")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesToBracketMinMax_withInvalidNumbers_returnsError(t *testing.T) {
	input := []byte("[1]")
	_, _, _, err := bytesToBracketsMinMax(
		input,
		[]byte("invalid numbers"),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
