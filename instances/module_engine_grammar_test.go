package instances

import (
	"strings"
	"testing"

	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
)

func TestModule_engine_grammar_value_Success(t *testing.T) {
	script := `
		module @newGrammarValue:7;;
		module @castToUint:1;;

		-> $myName;;
		<- $myValue;;

		// cast to uint:
		$myNumberStr = 157;;
		@castToUint $castToUintApp;;
		attach $myNumberStr:0 $castToUintApp;;
		$myNumber = execute $castToUintApp;;

		// value app:
		@newGrammarValue $valueApp;;
		attach $myNumber:0 $valueApp;;
		attach $myName:1 $valueApp;;

		$myValue = execute $valueApp;;
		invalid
	`

	name := "roger"
	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{
		name,
	}, []byte(script))

	if strings.TrimSpace(string(remaining)) != "invalid" {
		t.Errorf("the remaining space-trimmed command was expected: '%s', '%s' returned", "invalid", strings.TrimSpace(string(remaining)))
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(values.Value); ok {
		retName := ins.Name()
		if retName != name {
			t.Errorf("the name was expected to be '%s', '%s' returned", name, retName)
			return
		}

		retNumber := ins.Number()
		if retNumber != 157 {
			t.Errorf("the number was expected to be %d, %d returned", 157, retNumber)
			return
		}
		return
	}

	t.Errorf("the myValue output was expected to contain a Value instance")
	return
}
