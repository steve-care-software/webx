package vms

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

func TestRodan_grammar_Success(t *testing.T) {
	// read the script:
	script, err := ioutil.ReadFile("./../rodan/index.rodan")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	virtualMachine := NewApplication(func() (modules.Modules, error) {
		return newModules(newInterpreterModulesFuncs()), nil
	})

	result, err := virtualMachine.ParseThenInterpret([]interface{}{}, script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if result.HasRemaining() {
		t.Errorf("the result was expected to NOT contain remaining: %s", result.Remaining())
		return
	}

	if !result.IsValid() {
		t.Errorf("the result was expected to be valid")
		return
	}

	if !result.HasValues() {
		t.Errorf("the result was expected to contain values")
		return
	}

	output := result.Values()

	fmt.Printf("\n%v\n", output)
}
