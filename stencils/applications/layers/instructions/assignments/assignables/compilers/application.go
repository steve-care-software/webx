package compilers

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
)

type application struct {
	instanceAdapter   instances.Adapter
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	instanceAdapter instances.Adapter,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		instanceAdapter:   instanceAdapter,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable compilers.Compiler) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsCompile() {
		comVar := assignable.Compile()
		data, err := frame.FetchBytes(comVar)
		if err != nil {
			code := failures.CouldNotFetchCompileFromFrame
			return nil, &code, err
		}

		ins, err := app.instanceAdapter.ToInstance(data)
		if err != nil {
			code := failures.CouldNotCompileBytesToInstance
			return nil, &code, err
		}

		builder.WithInstance(ins)
	}

	if assignable.IsDecompile() {
		decVar := assignable.Decompile()
		ins, err := frame.FetchInstance(decVar)
		if err != nil {
			code := failures.CouldNotFetchDecompileFromFrame
			return nil, &code, err
		}

		data, err := app.instanceAdapter.ToBytes(ins)
		if err != nil {
			code := failures.CouldNotDecompileInstanceToBytes
			return nil, &code, err
		}

		builder.WithBytes(data)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
