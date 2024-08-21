package grammars

import (
	"bytes"
	"testing"
)

func TestBytesToMinMax_withZeroPlus_Success(t *testing.T) {
	expectedRemaining := []byte("this is some data")
	input := append([]byte(cardinalityZeroPlus), expectedRemaining...)
	retMin, pRetMax, retRemaining, err := bytesToMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(cardinalityZeroPlus)[0],
		[]byte(cardinalityOnePlus)[0],
		[]byte(cardinalityOptional)[0],
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

	if retMin != 0 {
		t.Errorf("the min was expected to be %d, %d returned", 0, retMin)
		return
	}

	if pRetMax != nil {
		t.Errorf("the returned max was expected to be nil, value returned: %d", *pRetMax)
		return
	}
}

func TestBytesToMinMax_withOnePlus_Success(t *testing.T) {
	expectedRemaining := []byte("this is some data")
	input := append([]byte(cardinalityOnePlus), expectedRemaining...)
	retMin, pRetMax, retRemaining, err := bytesToMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(cardinalityZeroPlus)[0],
		[]byte(cardinalityOnePlus)[0],
		[]byte(cardinalityOptional)[0],
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

func TestBytesToMinMax_withBrackets_Success(t *testing.T) {
	expectedRemaining := []byte("this is some data")
	input := append([]byte("[1,]"), expectedRemaining...)
	retMin, pRetMax, retRemaining, err := bytesToMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(cardinalityZeroPlus)[0],
		[]byte(cardinalityOnePlus)[0],
		[]byte(cardinalityOptional)[0],
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

func TestBytesToMinMax_withoutInput_returnsError(t *testing.T) {
	input := []byte("")
	_, _, _, err := bytesToMinMax(
		input,
		createPossibleNumbers(),
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(cardinalityZeroPlus)[0],
		[]byte(cardinalityOnePlus)[0],
		[]byte(cardinalityOptional)[0],
		[]byte(filterBytes),
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
