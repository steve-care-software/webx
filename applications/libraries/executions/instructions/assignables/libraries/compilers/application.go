package compilers

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_libraries "github.com/steve-care-software/datastencil/domain/stacks/libraries"
)

type application struct {
	instanceAdapter   orms.Adapter
	libraryBuilder    stacks_libraries.Builder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	instanceAdapter orms.Adapter,
	libraryBuilder stacks_libraries.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		instanceAdapter:   instanceAdapter,
		libraryBuilder:    libraryBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable compilers.Compiler) (stacks.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsCompile() {
		comVar := assignable.Compile()
		data, err := frame.FetchBytes(comVar)
		if err != nil {
			return nil, err
		}

		ins, err := app.instanceAdapter.ToInstance(data)
		if err != nil {
			return nil, err
		}

		lib, err := app.libraryBuilder.Create().
			WithInstance(ins).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithLibrary(lib)
	}

	if assignable.IsDecompile() {
		decVar := assignable.Decompile()
		ins, err := frame.FetchInstance(decVar)
		if err != nil {
			return nil, err
		}

		data, err := app.instanceAdapter.ToData(ins)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(data)
	}

	return builder.Now()
}
