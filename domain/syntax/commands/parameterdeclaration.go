package commands

import "github.com/steve-care-software/syntax/domain/syntax/criterias"

type parameterDeclaration struct {
	input  criterias.Criteria
	output criterias.Criteria
}

func createParameterDeclaration(
	input criterias.Criteria,
	output criterias.Criteria,
) ParameterDeclaration {
	out := parameterDeclaration{
		input:  input,
		output: output,
	}

	return &out
}

// Input returns the input
func (obj *parameterDeclaration) Input() criterias.Criteria {
	return obj.input
}

// Output returns the output
func (obj *parameterDeclaration) Output() criterias.Criteria {
	return obj.output
}
