package interpreters

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestExecute_Success(t *testing.T) {
	script := [][]byte{
		{NFive, NSix, SPlus, NZero, NSix, NOne},
	}

	application := NewApplication()
	output, err := application.Execute(bytes.Join(script, []byte{}))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	fmt.Printf("\n%v\n", output)
	panic(errors.New("stop"))
}
