package instances

import (
	"testing"

	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
)

func TestGrammar_Success(t *testing.T) {
	grammarApp := grammar_applications.NewApplication()
	ins := NewGrammar()
	script := `
		module @myModule:0;;
		@myModule $myApp;;

		-> $myInput;;
		<- $myOutput;;

		$constantVariable = this is a constant;;
		$assignedVariable = $myInput;;
		$instructions = {
			-> $first;;
			<- $output;;

			$output = $first;;
		};;

		attach $myInput:0 $myApp;;

		execute $myApp;;
		$appExec = execute $myApp;;
		`

	treeIns, err := grammarApp.Execute(ins, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if treeIns.HasRemaining() {
		t.Errorf("the tree was expected to NOT contain remaining data")
		return
	}

}
