package constants

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
)

type application struct {
	assignablesBuilder stacks.AssignablesBuilder
	assignableBuilder  stacks.AssignableBuilder
}

func createApplication(
	assignablesBuilder stacks.AssignablesBuilder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		assignablesBuilder: assignablesBuilder,
		assignableBuilder:  assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(assignable constants.Constant) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsBytes() {
		bytes := assignable.Bytes()
		builder.WithBytes(bytes)
	}

	if assignable.IsBool() {
		pBool := assignable.Bool()
		builder.WithBool(*pBool)
	}

	if assignable.IsString() {
		pValue := assignable.String()
		builder.WithString(*pValue)
	}

	if assignable.IsFloat() {
		pValue := assignable.Float()
		builder.WithFloat(*pValue)
	}

	if assignable.IsInt() {
		pValue := assignable.Int()
		builder.WithInt(*pValue)
	}

	if assignable.IsUint() {
		pValue := assignable.Uint()
		builder.WithUnsignedInt(*pValue)
	}

	if assignable.IsList() {
		output := []stacks.Assignable{}
		list := assignable.List().List()
		for _, oneAssignable := range list {
			retAssignable, pCode, err := app.Execute(oneAssignable)
			if err != nil {
				return nil, nil, err
			}

			if pCode != nil {
				return nil, pCode, nil
			}

			output = append(output, retAssignable)
		}

		stackAssignables, err := app.assignablesBuilder.Create().
			WithList(output).
			Now()

		if err != nil {
			return nil, nil, err
		}

		builder.WithList(stackAssignables)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
