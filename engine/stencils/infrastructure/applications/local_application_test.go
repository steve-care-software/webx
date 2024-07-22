package applications

import (
	"os"
	"path/filepath"
	"testing"

	json_executions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions"
)

func TestLocalApplication_withoutInput_isPro_Success(t *testing.T) {
	basePath := []string{
		"test_files",
		"files",
	}

	dbPath := []string{
		"database",
		"mydatabase.data",
	}

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	localApp, err := NewLocalApplicationBuilder().Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	script := `{
		"instructions": [
			{
				"assignment": {
					"name": "myReturneValue",
					"assignable": {
						"constant": {
							"bytes": [45, 65, 56, 6]
						}
					}
				}
			}
		],
		"output": {
			"variable": "myReturneValue",
			"kind": {
				"prompt": true
			}
		}
	}`

	pContext, err := localApp.Init(dbPath, "myName", "This is a description")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes, err := localApp.Execute(*pContext, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExecution, err := json_executions.NewAdapter().BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result := retExecution.Result()
	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	output := result.Success().Output()
	if output.HasExecute() {
		t.Errorf("the output was expected to NOT execute")
		return
	}

	if len(output.Input()) != 4 {
		t.Errorf("the output's input was expected to contain 4 bytes")
		return
	}

	err = localApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
