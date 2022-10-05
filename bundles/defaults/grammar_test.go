package defaults

import (
	"fmt"
	"testing"
)

func TestCreateGrammar_Success(t *testing.T) {
	grammarIns, err := NewGrammarCreateApplication().Execute()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", grammarIns)
}
