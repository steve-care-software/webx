package commands

import "github.com/steve-care-software/syntax/domain/bytes/criterias"

type parameterDeclaration struct {
	input  criterias.Criteria
	output criterias.Criteria
	name   criterias.Criteria
}

func createParameterDeclaration(
	input criterias.Criteria,
	output criterias.Criteria,
	name criterias.Criteria,
) ParameterDeclaration {
	out := parameterDeclaration{
		input:  input,
		output: output,
		name:   name,
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

// Name returns the name
func (obj *parameterDeclaration) Name() criterias.Criteria {
	return obj.name
}
