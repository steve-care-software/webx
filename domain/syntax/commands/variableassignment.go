package commands

import "github.com/steve-care-software/syntax/domain/syntax/criterias"

type variableAssignment struct {
	assignee criterias.Criteria
	value    Value
}

func createVariableAssignment(
	assignee criterias.Criteria,
	value Value,
) VariableAssignment {
	out := variableAssignment{
		assignee: assignee,
		value:    value,
	}

	return &out
}

// Assignee returns the assignee
func (obj *variableAssignment) Assignee() criterias.Criteria {
	return obj.assignee
}

// Value returns the value
func (obj *variableAssignment) Value() Value {
	return obj.value
}
