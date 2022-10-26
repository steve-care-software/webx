package defaults

import (
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

	grammarApp := grammar_applications.NewApplication()
	selectionApp := selection_applications.NewApplication()
	tree, err := grammarApp.Execute(grammarIns, []byte(fullScriptForTests))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	converted, err := selectionApp.Convert(tree, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// root
	rootCriteria := createCommandApp.Root()
	rootSelection, err := selectionApp.Search(converted, rootCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// instructions:
	instructionsCriteria := createCommandApp.Instructions()
	instructionsSelection, err := selectionApp.Search(rootSelection, instructionsCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructionsList := instructionsSelection.List()
	if len(instructionsList) != 15 {
		t.Errorf("%d instructions were expected, %d returned", 15, len(instructionsList))
		return
	}

	// assignments:
	amount := uint(12)
	assignmentsCriteria := createCommandApp.Assignment(0, &amount)
	assignmentSelection, err := selectionApp.Search(instructionsSelection, assignmentsCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	assignmentList := assignmentSelection.List()
	if len(assignmentList) != 5 {
		t.Errorf("%d instructions were expected, %d returned", 5, len(assignmentList))
		return
	}
}
