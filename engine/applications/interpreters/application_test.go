package interpreters

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestExecute_Success(t *testing.T) {
	script := [][]byte{
		{LLM, LLY, ULV, LLA, LLR, SEqual, NFive, NSix, SPlus, NZero, NSix, NOne}, // 56 + 061
		//{LLM, LLY, ULV, LLA, LLR}, // myVar
		//{NFive, NSix, SPlus, NZero, NSix, NOne}, // 56 + 061

	}

	application := NewApplication().(*application)
	isValid, retRemaining, _, err := application.entry(bytes.Join(script, []byte{}), "bRoot")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	fmt.Printf("\n%t, %v\n", isValid, retRemaining)
	panic(errors.New("stop"))
}
