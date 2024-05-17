package signatures

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	voteApp           votes.Application
	signApp           signs.Application
	assignableBuilder stacks.AssignableBuilder
	signerFactory     signers.Factory
}

func createApplication(
	voteApp votes.Application,
	signApp signs.Application,
	assignableBuilder stacks.AssignableBuilder,
	signerFactory signers.Factory,
) Application {
	out := application{
		voteApp:           voteApp,
		signApp:           signApp,
		assignableBuilder: assignableBuilder,
		signerFactory:     signerFactory,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable signatures.Signature) (stacks.Assignable, *uint, error) {
	if assignable.IsFetchPublicKey() {
		pkVar := assignable.FetchPublicKey()
		pk, err := frame.FetchSigner(pkVar)
		if err != nil {
			code := failures.CouldNotFetchSignerPrivateKeyFromFrame
			return nil, &code, err
		}

		pubKey := pk.PublicKey()
		ins, err := app.assignableBuilder.Create().WithSignerPubKey(pubKey).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil
	}

	if assignable.IsSign() {
		sign := assignable.Sign()
		return app.signApp.Execute(frame, sign)
	}

	if assignable.IsVote() {
		vote := assignable.Vote()
		return app.voteApp.Execute(frame, vote)
	}

	signer := app.signerFactory.Create()
	ins, err := app.assignableBuilder.Create().WithSigner(signer).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
