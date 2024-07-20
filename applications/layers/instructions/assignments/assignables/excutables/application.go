package executables

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_executables "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executables"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	localAppBuilder   applications.LocalBuilder
	remoteAppBuilder  applications.RemoteBuilder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	localAppBuilder applications.LocalBuilder,
	remoteAppBuilder applications.RemoteBuilder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		localAppBuilder:   localAppBuilder,
		remoteAppBuilder:  remoteAppBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable instruction_executables.Executable) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsLocal() {
		localVar := assignable.Local()
		localPath, err := frame.FetchList(localVar)
		if err != nil {
			return nil, nil, err
		}

		path := []string{}
		values := localPath.List()
		for _, oneValue := range values {
			if !oneValue.IsString() {
				code := failures.CouldNotFetchStringFromList
				return nil, &code, nil
			}

			pDirName := oneValue.String()
			path = append(path, *pDirName)
		}

		retApp, err := app.localAppBuilder.Create().WithBasePath(path).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithApplication(retApp)
	}

	if assignable.IsRemote() {
		remoteVar := assignable.Remote()
		remoteHost, err := frame.FetchString(remoteVar)
		if err != nil {
			code := failures.CouldNotFetchStringFromFrame
			return nil, &code, nil
		}

		retApp, err := app.remoteAppBuilder.Create().WithHost(remoteHost).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithApplication(retApp)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
