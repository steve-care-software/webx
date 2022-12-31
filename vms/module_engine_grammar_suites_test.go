package vms

import (
	"testing"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

func TestModule_engineGrammarSuites_Success(t *testing.T) {
	script := `
        <- $suites;;

        // suite app:
        module @newGrammarSuite:12;;
		@newGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:0 $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $invalid:1 $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        module @containerList:6;;
        @containerList $listApp;;
        attach $first:0 $listApp;;
        attach $second:1 $listApp;;
        $list = execute $listApp;;

        // suites:
        module @newGrammarSuites:13;;
        @newGrammarSuites $suitesApp;;
        attach $list:0 $suitesApp;;
        $suites = execute $suitesApp;;

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

	if ins, ok := output[0].(grammars.Suites); ok {
		list := ins.List()
		if len(list) != 2 {
			t.Errorf("%d Suite instances were expected, %d returned", 2, len(list))
			return
		}

		return
	}

	t.Errorf("the suites output was expected to contain a Suites instance")
	return
}
