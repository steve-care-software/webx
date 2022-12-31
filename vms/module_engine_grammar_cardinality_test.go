package vms

import (
	"testing"

	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
)

func TestModule_engineGrammarCardinality_withMin_withoutMax_Success(t *testing.T) {
	script := `
		module @newGrammarCardinality:8;;
		@newGrammarCardinality $cardinalityApp;;
		<- $output;;

		// cast to uint app:
		module @castToUint:1;;
		@castToUint $castToUintApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:0 $castToUintApp;;
		$myMin = execute $castToUintApp;;

		// cardinality app:
		attach $myMin:0 $cardinalityApp;;
		$output = execute $cardinalityApp;;
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

	if ins, ok := output[0].(cardinalities.Cardinality); ok {
		retMin := ins.Min()
		if retMin != 1 {
			t.Errorf("the min was expected to be '%d', '%d' returned", 1, retMin)
			return
		}

		if ins.HasMax() {
			t.Errorf("the cardinality was expected to NOT contain a max")
			return
		}

		return
	}

	t.Errorf("the output variable was expected to contain a Cardinality instance")
	return
}

func TestModule_engineGrammarCardinality_withMin_withMax_Success(t *testing.T) {
	script := `
		module @newGrammarCardinality:8;;
		@newGrammarCardinality $cardinalityApp;;
		<- $output;;

		// cast to uint app:
		module @castToUint:1;;
		@castToUint $castToUintApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:0 $castToUintApp;;
		$myMin = execute $castToUintApp;;

		// max casting to uint:
		$myMaxStr = 2;;
		attach $myMaxStr:0 $castToUintApp;;
		$myMax = execute $castToUintApp;;

		// cardinality app:
		attach $myMin:0 $cardinalityApp;;
        attach $myMax:1 $cardinalityApp;;

		$output = execute $cardinalityApp;;
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

	if ins, ok := output[0].(cardinalities.Cardinality); ok {
		retMin := ins.Min()
		if retMin != 1 {
			t.Errorf("the min was expected to be '%d', '%d' returned", 1, retMin)
			return
		}

		if !ins.HasMax() {
			t.Errorf("the cardinality was expected to contain a max")
			return
		}

		pMax := ins.Max()
		if *pMax != 2 {
			t.Errorf("the max was expected to be '%d', '%d' returned", 2, *pMax)
			return
		}

		return
	}

	t.Errorf("the output variable was expected to contain a Cardinality instance")
	return
}
