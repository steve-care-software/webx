package defaults

import (
	"testing"

	engines "github.com/steve-care-software/webx/applications"
	"github.com/steve-care-software/webx/domain/grammars"
)

func TestModule_newGrammarSuites_Success(t *testing.T) {
	script := `
        <- $suites;;

        // suite app:
        module @newGrammarSuite;;
		@newGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:$valid $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $invalid:$invalid $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        module @containerList;;
        @containerList $listApp;;
        attach $first:$0 $listApp;;
        attach $second:$1 $listApp;;
        $list = execute $listApp;;

        // suites:
        module @newGrammarSuites;;
        @newGrammarSuites $suitesApp;;
        attach $list:$suites $suitesApp;;
        $suites = execute $suitesApp;;

	`

	output, _, err := engines.NewApplication(NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{}, []byte(script))
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
