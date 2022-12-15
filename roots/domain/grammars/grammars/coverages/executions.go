package coverages

type executions struct {
	list []Execution
}

func createExecutions(
	list []Execution,
) Executions {
	out := executions{
		list: list,
	}

	return &out
}

// List returns the executions
func (obj *executions) List() []Execution {
	return obj.list
}

// ContainsError returns true if it contains an error, false otherwise
func (obj *executions) ContainsError() bool {
	for _, oneExecution := range obj.list {
		if !oneExecution.Result().IsError() {
			continue
		}

		return true
	}

	return false
}
