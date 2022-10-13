package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarElement_withValue_Success(t *testing.T) {
	script := `
		-> $name;;
		<- $element;;

		// cast to uint app:
		module @castToUint;;
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:$value $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
		module @engineGrammarValue;;
		@engineGrammarValue $valueApp;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:$value $castToUintApp;;
		$myMin = execute $castToUintApp;;

        // cardinality:
        module @engineGrammarCardinality;;
		@engineGrammarCardinality $cardinalityApp;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
        module @engineGrammarElement;;
		@engineGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

	`

	name := "roger"
	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{
		"name": name,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["element"].(grammars.Element); ok {
		if !ins.Content().IsValue() {
			t.Errorf("the element was expecting a value instance")
			return
		}

		return
	}

	t.Errorf("the element output was expected to contain an Element instance")
	return
}
