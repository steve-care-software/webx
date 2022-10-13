package defaults

import (
	"strings"
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
)

func TestModule_engineGrammarValue_Success(t *testing.T) {
	script := `
		module @engineGrammarValue;;
		module @castToUint;;

		-> $myName;;
		<- $myValue;;

		// cast to uint:
		$myNumberStr = 157;;
		@castToUint $castToUintApp;;
		attach $myNumberStr:$value $castToUintApp;;
		$myNumber = execute $castToUintApp;;

		// value app:
		@engineGrammarValue $valueApp;;
		attach $myNumber:$number $valueApp;;
		attach $myName:$name $valueApp;;

		$myValue = execute $valueApp;;
		invalid
	`

	name := "roger"
	output, remaining, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{
		"myName": name,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if strings.TrimSpace(string(remaining)) != "invalid" {
		t.Errorf("the remaining space-trimmed command was expected: '%s', '%s' returned", "invalid", strings.TrimSpace(string(remaining)))
		return
	}

	if ins, ok := output["myValue"].(values.Value); ok {
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
