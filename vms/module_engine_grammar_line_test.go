package vms

import (
	"testing"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

func TestModule_engineGrammarLine_Success(t *testing.T) {
	script := `
		-> $name;;
		<- $line;;

		// cast to uint app:
		module @castToUint:1;;
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:0 $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
		module @newGrammarValue:7;;
		@newGrammarValue $valueApp;;
        $name = myName;;
		attach $number:0 $valueApp;;
		attach $name:1 $valueApp;;
        $value = execute $valueApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:0 $castToUintApp;;
		$myMin = execute $castToUintApp;;

        // cardinality:
        module @newGrammarCardinality:8;;
		@newGrammarCardinality $cardinalityApp;;
		attach $myMin:0 $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
        module @newGrammarElement:9;;
		@newGrammarElement $elementApp;;
        attach $cardinality:0 $elementApp;;
        attach $value:1 $elementApp;;
        $element = execute $elementApp;;

        // elements:
        module @containerList:6;;
        @containerList $listApp;;
        attach $element:0 $listApp;;
        $elements = execute $listApp;;

        // line:
        module @newGrammarLine:10;;
		@newGrammarLine $lineApp;;
        attach $elements:0 $lineApp;;
        $line = execute $lineApp;;

	`

	name := "roger"
	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{
		name,
	}, []byte(script))

	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Line); ok {
		list := ins.Elements()
		if len(list) != 1 {
			t.Errorf("%d elements were expected, %d returned", 1, len(list))
			return
		}

		return
	}

	t.Errorf("the line output was expected to contain a Line instance")
	return
}
