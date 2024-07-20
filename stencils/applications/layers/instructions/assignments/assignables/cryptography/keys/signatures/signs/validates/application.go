package validates

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
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
func (app *application) Execute(frame stacks.Frame, assignable validates.Validate) (stacks.Assignable, *uint, error) {
	sigName := assignable.Signature()
	sig, err := frame.FetchSignature(sigName)
	if err != nil {
		code := failures.CouldNotFetchSignatureFromFrame
		return nil, &code, err
	}

	msgVar := assignable.Message()
	msg, err := frame.FetchBytes(msgVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	pubKeyName := assignable.PublicKey()
	expectedPubKey, err := frame.FetchSignerPubKey(pubKeyName)
	if err != nil {
		code := failures.CouldNotFetchSignerPrivateKeyFromFrame
		return nil, &code, err
	}

	validated := sig.PublicKey(string(msg)).Equals(expectedPubKey)
	ins, err := app.assignableBuilder.Create().WithBool(validated).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
