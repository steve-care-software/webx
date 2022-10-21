package coverages

import (
	"github.com/steve-care-software/webx/domain/grammars"
)

type execution struct {
	expectation grammars.Suite
	result      Result
}

func createExecution(
	expectation grammars.Suite,
	result Result,
) Execution {
	out := execution{
		expectation: expectation,
		result:      result,
	}

	return &out
}

// Expectation returns the expectation
func (obj *execution) Expectation() grammars.Suite {
	return obj.expectation
}

// Result returns the result
func (obj *execution) Result() Result {
	return obj.result
}
