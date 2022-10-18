package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_newGrammarSuite_withValid_withString_Success(t *testing.T) {
	script := `
        <- $suite;;

        // suite app
        module @newGrammarSuite;;
		@newGrammarSuite $suiteApp;;
		$valid = 157;;
		attach $valid:$valid $suiteApp;;
        $suite = execute $suiteApp;;

	`

	output, _, err := engines.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["suite"].(grammars.Suite); ok {
		if !ins.IsValid() {
			t.Errorf("the suite was expected to be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}

func TestModule_newGrammarSuite_withValid_withByte_Success(t *testing.T) {
	script := `
        -> $input;;
        <- $suite;;

        // suite app
        module @newGrammarSuite;;
		@newGrammarSuite $suiteApp;;
		attach $input:$valid $suiteApp;;
        $suite = execute $suiteApp;;

	`

	output, _, err := engines.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{
		"input": []byte("this is some data"),
	}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["suite"].(grammars.Suite); ok {
		if !ins.IsValid() {
			t.Errorf("the suite was expected to be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}

func TestModule_newGrammarSuite_withInvalid_withString_Success(t *testing.T) {
	script := `
        <- $suite;;

        // suite app
        module @newGrammarSuite;;
		@newGrammarSuite $suiteApp;;
		$invalid = 157;;
		attach $valid:$invalid $suiteApp;;
        $suite = execute $suiteApp;;

	`

	output, _, err := engines.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["suite"].(grammars.Suite); ok {
		if ins.IsValid() {
			t.Errorf("the suite was expected to be NOT be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}

func TestModule_newGrammarSuite_withInvalid_withByte_Success(t *testing.T) {
	script := `
        -> $input;;
        <- $suite;;

        // suite app
        module @newGrammarSuite;;
		@newGrammarSuite $suiteApp;;
		attach $input:$invalid $suiteApp;;
        $suite = execute $suiteApp;;

	`

	output, _, err := engines.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{
		"input": []byte("this is some data"),
	}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["suite"].(grammars.Suite); ok {
		if ins.IsValid() {
			t.Errorf("the suite was expected to be NOT be valid")
			return
		}

		return
	}

	t.Errorf("the suite output was expected to contain a Suite instance")
	return
}
