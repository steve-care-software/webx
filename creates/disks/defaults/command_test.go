package defaults

import (
	"fmt"
	"testing"

	grammar_applications "github.com/steve-care-software/webx/applications/grammars"
	selection_applications "github.com/steve-care-software/webx/applications/selections"
)

func TestCommand_Success(t *testing.T) {
	createApp := NewApplication(
		bitrateForTests,
		basePathForTests,
		delimiterForTests,
		extensionForTests,
	)

	createGrammarApp := createApp.Grammar()
	grammarIns, err := createGrammarApp.Execute()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	createCommandApp := createApp.Command().(*command)

	variableNameCriteria := createCommandApp.VariableName(0)

	grammarApp := grammar_applications.NewApplication()
	selectionApp := selection_applications.NewApplication()

	tree, err := grammarApp.Execute(grammarIns, []byte("-> $myVariable;;"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	converted, err := selectionApp.Convert(tree, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	searched, err := selectionApp.Search(converted, variableNameCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", searched)

}
