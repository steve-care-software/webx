package applications

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"
	json_executions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions"
)

type testExecution func(t *testing.T, execution executions.Execution)
type testCase struct {
	input  string
	path   []string
	result testExecution
}

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

	testCases := []testCase{
		{
			input: `{
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
			}`,
			result: func(t *testing.T, execution executions.Execution) {
				retResult := execution.Result()
				if !retResult.IsSuccess() {
					t.Errorf("the result was expected to be successful")
					return
				}

				output := retResult.Success().Output()
				if output.HasExecute() {
					t.Errorf("the output was expected to NOT execute")
					return
				}

				if len(output.Input()) != 4 {
					t.Errorf("the output's input was expected to contain 4 bytes")
					return
				}
			},
		},
		{
			input: `012test`,
			path: []string{
				".",
				"..",
				"..",
				"..",
				"..",
				"scripts",
				"extractions",
				"numbers",
				"number",
				"index.json",
			},
			result: func(t *testing.T, execution executions.Execution) {
				retResult := execution.Result()
				if !retResult.IsSuccess() {
					t.Errorf("the result was expected to be successful")
					return
				}

				output := retResult.Success().Output()
				if output.HasExecute() {
					t.Errorf("the output was expected to NOT execute")
					return
				}

				if !bytes.Equal(output.Input(), []byte{0}) {
					t.Errorf("the output's input was expected to contain 4 bytes")
					return
				}
			},
		},
	}

	for index, oneTestCase := range testCases {
		pContext, err := localApp.Init(dbPath, "myName", "This is a description")
		if err != nil {
			t.Errorf("%d, the error was expected to be nil, error returned: %s", index, err.Error())
			return
		}

		hasPath := oneTestCase.path != nil && len(oneTestCase.path) > 0
		if oneTestCase.result != nil && !hasPath {
			retBytes, err := localApp.Execute(*pContext, []byte(oneTestCase.input))
			if err != nil {
				t.Errorf("%d, the error was expected to be nil, error returned: %s", index, err.Error())
				return
			}

			retExecution, err := json_executions.NewAdapter().BytesToInstance(retBytes)
			if err != nil {
				t.Errorf("%d, the error was expected to be nil, error returned: %s", index, err.Error())
				return
			}

			fn := oneTestCase.result
			fn(t, retExecution)
		}

		if oneTestCase.path != nil && hasPath {
			retBytes, err := localApp.ExecuteLayer(*pContext, []byte(oneTestCase.input), oneTestCase.path)
			if err != nil {
				t.Errorf("%d, the error was expected to be nil, error returned: %s", index, err.Error())
				return
			}

			retExecution, err := json_executions.NewAdapter().BytesToInstance(retBytes)
			if err != nil {
				t.Errorf("%d, the error was expected to be nil, error returned: %s", index, err.Error())
				return
			}

			fn := oneTestCase.result
			fn(t, retExecution)
		}

		err = localApp.Commit(*pContext)
		if err != nil {
			t.Errorf("%d, the error was expected to be nil, error returned: %s", index, err.Error())
			return
		}

		os.RemoveAll(filepath.Join(basePath...))
	}

}
