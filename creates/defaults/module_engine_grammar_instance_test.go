package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarInstance_withEverything_Success(t *testing.T) {
	script := `
        module @engineGrammarSuite;;
        module @containerList;;
        module @engineGrammarSuites;;
        module @engineGrammarValue;;
        module @engineGrammarCardinality;;
        module @engineGrammarElement;;
        module @engineGrammarLine;;
        module @engineGrammarBlock;;
        module @engineGrammarToken;;
        module @engineGrammarEverything;;
        module @engineGrammarInstance;;

        <- $instance;;

        // suite app:
		@engineGrammarSuite $suiteApp;;

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
        @engineGrammarSuites $suitesApp;;
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
		@engineGrammarCardinality $cardinalityApp;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@engineGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@engineGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@engineGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName = myToken;;
		@engineGrammarToken $tokenApp;;
        attach $tokenName:$name $tokenApp;;
        attach $suites:$suites $tokenApp;;
        attach $block:$block $tokenApp;;
        $token = execute $tokenApp;;

        // everything:
        $everythingName = myEverything;;
		@engineGrammarEverything $everythingApp;;
        attach $everythingName:$name $everythingApp;;
        attach $token:$exception $everythingApp;;
        $everything = execute $everythingApp;;

        // instance:
		@engineGrammarInstance $instanceApp;;
        attach $everything:$everything $instanceApp;;
        $instance = execute $instanceApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["instance"].(grammars.Instance); ok {
		if !ins.IsEverything() {
			t.Errorf("the instance was expected to contain an Everything instance")
			return
		}

		return
	}

	t.Errorf("the instance output was expected to contain an Instance instance")
	return
}

func TestModule_engineGrammarInstance_withToken_Success(t *testing.T) {
	script := `
        module @engineGrammarSuite;;
        module @containerList;;
        module @engineGrammarSuites;;
        module @engineGrammarValue;;
        module @engineGrammarCardinality;;
        module @engineGrammarElement;;
        module @engineGrammarLine;;
        module @engineGrammarBlock;;
        module @engineGrammarToken;;
        module @engineGrammarEverything;;
        module @engineGrammarInstance;;

        <- $instance;;

        // suite app:
		@engineGrammarSuite $suiteApp;;

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
        @engineGrammarSuites $suitesApp;;
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
		@engineGrammarCardinality $cardinalityApp;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@engineGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@engineGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@engineGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName = myToken;;
		@engineGrammarToken $tokenApp;;
        attach $tokenName:$name $tokenApp;;
        attach $suites:$suites $tokenApp;;
        attach $block:$block $tokenApp;;
        $token = execute $tokenApp;;

        // instance:
		@engineGrammarInstance $instanceApp;;
        attach $token:$token $instanceApp;;
        $instance = execute $instanceApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["instance"].(grammars.Instance); ok {
		if !ins.IsToken() {
			t.Errorf("the instance was expected to contain a Token instance")
			return
		}

		return
	}

	t.Errorf("the instance output was expected to contain an Instance instance")
	return
}
