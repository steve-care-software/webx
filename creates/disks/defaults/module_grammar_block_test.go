package defaults

import (
	"testing"

	"github.com/steve-care-software/webx/applications"
	"github.com/steve-care-software/webx/domain/grammars"
)

func TestModule_newGrammarBlock_Success(t *testing.T) {
	script := `
		-> $name;;
		<- $block;;

		// cast to uint app:
		module @castToUint;;
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:$value $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
        module @newGrammarValue;;
		@newGrammarValue $valueApp;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:$value $castToUintApp;;
		$myMin = execute $castToUintApp;;

        // cardinality:
        module @newGrammarCardinality;;
		@newGrammarCardinality $cardinalityApp;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
        module @newGrammarElement;;
		@newGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        module @containerList;;
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
        module @newGrammarLine;;
		@newGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
        module @newGrammarBlock;;
		@newGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

	`

	name := "roger"
	output, _, err := applications.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{
		valueToHashStringForTests("name"): name,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[valueToHashStringForTests("block")].(grammars.Block); ok {
		list := ins.Lines()
		if len(list) != 1 {
			t.Errorf("%d lines were expected, %d returned", 1, len(list))
			return
		}

		return
	}

	t.Errorf("the block output was expected to contain a Block instance")
	return
}
