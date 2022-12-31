package vms

import (
	"testing"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

func TestModule_engineGrammarEverything_withEscape_Success(t *testing.T) {
	script := `
        module @castToUint:1;;
        module @newGrammarSuite:12;;
        module @containerList:6;;
        module @newGrammarSuites:13;;
        module @newGrammarValue:7;;
        module @newGrammarCardinality:8;;
        module @newGrammarElement:9;;
        module @newGrammarLine:10;;
        module @newGrammarBlock:11;;
        module @newGrammarToken:14;;
        module @newGrammarEverything:15;;

        <- $everything;;

        // suite app:
		@newGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:0 $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $valid:1 $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        @containerList $suitesListApp;;
        attach $first:0 $suitesListApp;;
        attach $second:1 $suitesListApp;;
        $list = execute $suitesListApp;;

        // suites:
        @newGrammarSuites $suitesApp;;
        attach $list:0 $suitesApp;;
        $suites = execute $suitesApp;;

		// cast to uint app:
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:0 $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
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
		@newGrammarCardinality $cardinalityApp;;
		attach $myMin:0 $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@newGrammarElement $elementApp;;
        attach $cardinality:0 $elementApp;;
        attach $value:1 $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@newGrammarLine $lineApp;;
        attach $elements:0 $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@newGrammarBlock $blockApp;;
        attach $lines:0 $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName =myToken;;
		@newGrammarToken $tokenApp;;
        attach $tokenName:0 $tokenApp;;
        attach $block:1 $tokenApp;;
        attach $suites:2 $tokenApp;;
        $token = execute $tokenApp;;

        // everything:
        $everythingName = myEverything;;
		@newGrammarEverything $everythingApp;;
        attach $everythingName:0 $everythingApp;;
        attach $token:1 $everythingApp;;
        attach $token:2 $everythingApp;;
        $everything = execute $everythingApp;;

	`

	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{}, []byte(script))
	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Everything); ok {
		if !ins.HasEscape() {
			t.Errorf("the escape was expected to be valid")
			return
		}

		return
	}

	t.Errorf("the everything output was expected to contain an Everything instance")
	return
}

func TestModule_engineGrammarEverything_withoutEscape_Success(t *testing.T) {
	script := `
        module @castToUint:1;;
        module @newGrammarSuite:12;;
        module @containerList:6;;
        module @newGrammarSuites:13;;
        module @newGrammarValue:7;;
        module @newGrammarCardinality:8;;
        module @newGrammarElement:9;;
        module @newGrammarLine:10;;
        module @newGrammarBlock:11;;
        module @newGrammarToken:14;;
        module @newGrammarEverything:15;;

        <- $everything;;

        // suite app:
		@newGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:0 $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $valid:1 $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        @containerList $suitesListApp;;
        attach $first:0 $suitesListApp;;
        attach $second:1 $suitesListApp;;
        $list = execute $suitesListApp;;

        // suites:
        @newGrammarSuites $suitesApp;;
        attach $list:0 $suitesApp;;
        $suites = execute $suitesApp;;

		// cast to uint app:
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:0 $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
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
		@newGrammarCardinality $cardinalityApp;;
		attach $myMin:0 $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@newGrammarElement $elementApp;;
        attach $cardinality:0 $elementApp;;
        attach $value:1 $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@newGrammarLine $lineApp;;
        attach $elements:0 $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@newGrammarBlock $blockApp;;
        attach $lines:0 $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName =myToken;;
		@newGrammarToken $tokenApp;;
        attach $tokenName:0 $tokenApp;;
        attach $block:1 $tokenApp;;
        attach $suites:2 $tokenApp;;
        $token = execute $tokenApp;;

        // everything:
        $everythingName = myEverything;;
		@newGrammarEverything $everythingApp;;
        attach $everythingName:0 $everythingApp;;
        attach $token:1 $everythingApp;;
        $everything = execute $everythingApp;;

	`

	virtualMachine := NewApplication()
	output, remaining, err := virtualMachine.ParseThenInterpret([]interface{}{}, []byte(script))
	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty: %s", remaining)
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output[0].(grammars.Everything); ok {
		if ins.HasEscape() {
			t.Errorf("the escape was expected to NOT be valid")
			return
		}

		return
	}

	t.Errorf("the everything output was expected to contain an Everything instance")
	return
}
