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
	if len(instructionsList) != 17 {
		t.Errorf("%d instructions were expected, %d returned", 17, len(instructionsList))
		return
	}

	// assignment constant:
	constantAmount := uint(12)
	constantCriteria := createCommandApp.AssignmentConstant(0, &constantAmount)
	constantSelection, err := selectionApp.Search(instructionsSelection, constantCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	constantsList := constantSelection.List()
	if len(constantsList) != 2 {
		t.Errorf("%d instructions were expected, %d returned", 2, len(constantsList))
		return
	}

	// assignment variable:
	variableAmount := uint(12)
	variableCriteria := createCommandApp.AssignmentVariable(0, &variableAmount)
	variableSelection, err := selectionApp.Search(instructionsSelection, variableCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	variablesList := variableSelection.List()
	if len(variablesList) != 1 {
		t.Errorf("%d instructions were expected, %d returned", 1, len(variablesList))
		return
	}

	// assignment code:
	codeAmount := uint(12)
	codeCriteria := createCommandApp.AssignmentCode(0, &codeAmount)
	codeSelection, err := selectionApp.Search(instructionsSelection, codeCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	codeList := codeSelection.List()
	if len(codeList) != 1 {
		t.Errorf("%d instructions were expected, %d returned", 1, len(codeList))
		return
	}

	// assignment execution:
	executionAmount := uint(12)
	executionCriteria := createCommandApp.AssignmentExecution(0, &executionAmount)
	executionSelection, err := selectionApp.Search(instructionsSelection, executionCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executionList := executionSelection.List()
	if len(executionList) != 4 {
		t.Errorf("%d instructions were expected, %d returned", 4, len(executionList))
		return
	}

	// execution:
	execAmount := uint(12)
	execCriteria := createCommandApp.Execution(0, &execAmount)
	execSelection, err := selectionApp.Search(instructionsSelection, execCriteria)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	execList := execSelection.List()
	if len(execList) != 5 {
		t.Errorf("%d instructions were expected, %d returned", 5, len(execList))
		return
	}
}
