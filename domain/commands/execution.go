package commands

import "github.com/steve-care-software/syntax/domain/bytes/criterias"

type execution struct {
	application criterias.Criteria
	assignee    criterias.Criteria
}

func createExecution(
	application criterias.Criteria,
) Execution {
	return createExecutionInternally(application, nil)
}

func createExecutionWithAssignee(
	application criterias.Criteria,
	assignee criterias.Criteria,
) Execution {
	return createExecutionInternally(application, assignee)
}

func createExecutionInternally(
	application criterias.Criteria,
	assignee criterias.Criteria,
) Execution {
	out := execution{
		application: application,
		assignee:    assignee,
	}

	return &out
}

// Application returns the application
func (obj *execution) Application() criterias.Criteria {
	return obj.application
}

// HasAssignee returns true if there is an assignee, false otherwise
func (obj *execution) HasAssignee() bool {
	return obj.assignee != nil
}

// Assignee returns the assignee, if any
func (obj *execution) Assignee() criterias.Criteria {
	return obj.assignee
}
