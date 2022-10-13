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

        // value app
        module @engineGrammarValue;;
		@engineGrammarValue $valueApp;;
		$number = 157;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

        // cardinality:
        module @engineGrammarCardinality;;
		@engineGrammarCardinality $cardinalityApp;;
        $myMin = 1;;
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
