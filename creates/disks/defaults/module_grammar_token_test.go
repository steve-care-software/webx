package defaults

import (
	"testing"

	engines "github.com/steve-care-software/webx/applications"
	"github.com/steve-care-software/webx/domain/grammars"
)

func TestModule_newGrammarToken_Success(t *testing.T) {
	script := `
        module @newGrammarSuite;;
        module @containerList;;
        module @newGrammarSuites;;
        module @newGrammarValue;;
        module @newGrammarCardinality;;
        module @newGrammarElement;;
        module @newGrammarLine;;
        module @newGrammarBlock;;
        module @newGrammarToken;;

        <- $token;;

        // suite app:
		@newGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:$valid $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $valid:$invalid $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        @containerList $suitesListApp;;
        attach $first:$0 $suitesListApp;;
        attach $second:$1 $suitesListApp;;
        $list = execute $suitesListApp;;

        // suites:
        @newGrammarSuites $suitesApp;;
        attach $list:$suites $suitesApp;;
        $suites = execute $suitesApp;;

		// cast to uint app:
		module @castToUint;;
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:$value $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
		@newGrammarValue $valueApp;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:$value $castToUintApp;;
		$myMin = execute $castToUintApp;;

        // cardinality:
		@newGrammarCardinality $cardinalityApp;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@newGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@newGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@newGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName = myToken;;
		@newGrammarToken $tokenApp;;
        attach $tokenName:$name $tokenApp;;
        attach $suites:$suites $tokenApp;;
        attach $block:$block $tokenApp;;
        $token = execute $tokenApp;;

	`

	output, _, err := engines.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["token"].(grammars.Token); ok {
		name := ins.Name()
		if name != "myToken" {
			t.Errorf("the name was expected to be '%s', '%s' returned", "myToken", name)
			return
		}

		return
	}

	t.Errorf("the token output was expected to contain a Token instance")
	return
}
