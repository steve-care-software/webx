package vms

import (
	"testing"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

func TestModule_engineGrammarSuite_withValid_withString_Success(t *testing.T) {
	script := `
        <- $suite;;

        // suite app
        module @newGrammarSuite:12;;
		@newGrammarSuite $suiteApp;;
		$valid = 157;;
		attach $valid:0 $suiteApp;;
        $suite = execute $suiteApp;;

	`

	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{}, []byte(script))

	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Suite); ok {
		if !ins.IsValid() {
			t.Errorf("the suite was expected to be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}

func TestModule_engineGrammarSuite_withValid_withByte_Success(t *testing.T) {
	script := `
        -> $input;;
        <- $suite;;

        // suite app
        module @newGrammarSuite:12;;
		@newGrammarSuite $suiteApp;;
		attach $input:0 $suiteApp;;
        $suite = execute $suiteApp;;

	`

	input := []byte("this is some data")
	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{
		input,
	}, []byte(script))

	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Suite); ok {
		if !ins.IsValid() {
			t.Errorf("the suite was expected to be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}

func TestModule_engineGrammarSuite_withInvalid_withString_Success(t *testing.T) {
	script := `
        <- $suite;;

        // suite app
        module @newGrammarSuite:12;;
		@newGrammarSuite $suiteApp;;
		$invalid = 157;;
		attach $invalid:1 $suiteApp;;
        $suite = execute $suiteApp;;

	`

	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{}, []byte(script))

	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Suite); ok {
		if ins.IsValid() {
			t.Errorf("the suite was expected to be NOT be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}

func TestModule_engineGrammarSuite_withInvalid_withByte_Success(t *testing.T) {
	script := `
        -> $input;;
        <- $suite;;

        // suite app
        module @newGrammarSuite:12;;
		@newGrammarSuite $suiteApp;;
		attach $input:1 $suiteApp;;
        $suite = execute $suiteApp;;

	`

	input := []byte("this is some data")
	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{
		input,
	}, []byte(script))

	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Suite); ok {
		if ins.IsValid() {
			t.Errorf("the suite was expected to be NOT be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}
