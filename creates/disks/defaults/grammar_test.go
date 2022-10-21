package defaults

import (
	"fmt"
	"testing"

	"github.com/steve-care-software/webx/applications/grammars"
)

func TestGrammar_Success(t *testing.T) {
	grammarIns, err := NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests).Grammar().Execute()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	grammarApp := grammars.NewApplication()
	coverages, err := grammarApp.Coverages(grammarIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if coverages != nil {
		list := coverages.List()
		for _, oneCoverage := range list {
			tokenName := oneCoverage.Token().Name()
			executionsList := oneCoverage.Executions().List()
			for execIdx, oneExecution := range executionsList {
				expectation := oneExecution.Expectation()
				content := expectation.Content()
				result := oneExecution.Result()
				expected := expectation.IsValid()

				path := fmt.Sprintf("%s.%d=%v", tokenName, execIdx, string(content))
				if expected {
					if result.IsTree() {
						if result.Tree().Block().HasSuccessful() {
							continue
						}

						t.Errorf("the test suite expected to be successful, but there was no successful result, path: %s", path)
						continue
					}

					t.Errorf("the test suite expected to be successful, but there was an error while executing the test suite (path: %s): %s", path, result.Error())
					continue
				}

				if result.IsError() {
					continue
				}

				t.Errorf("the test suite expected to be unsuccessful, but the result was successful, path: %s", path)
			}

		}
	}

	uncovered, err := grammarApp.Uncovered(grammarIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	for tokenName, lines := range uncovered {
		for lineIdx, line := range lines {
			for elIdx, element := range line {
				t.Errorf("token: %s, line: %d, element: %d, name: %s - is uncovered", tokenName, lineIdx, elIdx, element)
			}
		}
	}
}
