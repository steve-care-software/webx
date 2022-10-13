package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarSuites_Success(t *testing.T) {
	script := `
        <- $suites;;

        // suite app:
        module @engineGrammarSuite;;
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
        module @containerList;;
        @containerList $listApp;;
        attach $first:$0 $listApp;;
        attach $second:$1 $listApp;;
        $list = execute $listApp;;

        // suites:
        module @engineGrammarSuites;;
        @engineGrammarSuites $suitesApp;;
        attach $list:$suites $suitesApp;;
        $suites = execute $suitesApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["suites"].(grammars.Suites); ok {
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
