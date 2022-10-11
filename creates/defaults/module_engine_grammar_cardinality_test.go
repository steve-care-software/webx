package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/cardinalities"
)

func TestModule_engineGrammarCardinality_withMin_withoutMax_Success(t *testing.T) {
	script := `
		module @engineGrammarCardinality;;
		@engineGrammarCardinality $cardinalityApp;;
		<- $output;;

		$myMin = 1;;
		attach $myMin:$min $cardinalityApp;;

		$output = execute $cardinalityApp;;
	`

	name := "roger"
	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{
		"myName": name,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["output"].(cardinalities.Cardinality); ok {
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
		module @engineGrammarCardinality;;
		@engineGrammarCardinality $cardinalityApp;;
		<- $output;;

		$myMin = 1;;
        $myMax = 2;;
		attach $myMin:$min $cardinalityApp;;
        attach $myMax:$max $cardinalityApp;;

		$output = execute $cardinalityApp;;
	`

	name := "roger"
	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{
		"myName": name,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["output"].(cardinalities.Cardinality); ok {
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
