package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarBlock_Success(t *testing.T) {
	script := `
		-> $name;;
		<- $block;;

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

        // elements:
        module @containerList;;
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
        module @engineGrammarLine;;
		@engineGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
        module @engineGrammarBlock;;
		@engineGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

	`

	name := "roger"
	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{
		"name": name,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["block"].(grammars.Block); ok {
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
