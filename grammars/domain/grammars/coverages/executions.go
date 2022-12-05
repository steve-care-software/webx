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
