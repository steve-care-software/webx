package commands

import "github.com/steve-care-software/webx/domain/criterias"

type execution struct {
	application criterias.Criteria
	assignee    criterias.Criteria
}

func createExecution(
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

// Assignee returns the assignee, if any
func (obj *execution) Assignee() criterias.Criteria {
	return obj.assignee
}
