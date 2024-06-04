package creates

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
	}
	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable creates.Create) (stacks.Assignable, *uint, error) {
	msgVar := assignable.Message()
	msg, err := frame.FetchBytes(msgVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	pkName := assignable.PrivateKey()
	pk, err := frame.FetchSigner(pkName)
	if err != nil {
		code := failures.CouldNotFetchSignerPrivateKeyFromFrame
		return nil, &code, err
	}

	signature, err := pk.Sign(string(msg))
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithSignature(signature).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
