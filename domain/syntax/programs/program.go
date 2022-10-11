package programs

import "github.com/steve-care-software/syntax/domain/syntax/programs/applications"

type program struct {
	assignments []applications.Assignment
	outputs     []string
}

func createProgram(
	assignments []applications.Assignment,
) Program {
	return createProgramInternally(assignments, nil)
}

func createProgramWithOutputs(
	assignments []applications.Assignment,
	outputs []string,
) Program {
	return createProgramInternally(assignments, outputs)
}

func createProgramInternally(
	assignments []applications.Assignment,
	outputs []string,
) Program {
	out := program{
		assignments: assignments,
		outputs:     outputs,
	}

	return &out
}

// Assignments returns the assignments
func (obj *program) Assignments() []applications.Assignment {
	return obj.assignments
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *program) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *program) Outputs() []string {
	return obj.outputs
}
