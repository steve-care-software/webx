package communications

import (
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	execSignApp       signs.Application
	execVoteApp       votes.Application
	signerFactory     signers.Factory
	accountBuilder    accounts.Builder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	execSignApp signs.Application,
	execVoteApp votes.Application,
	signerFactory signers.Factory,
	accountBuilder accounts.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		execSignApp:       execSignApp,
		execVoteApp:       execVoteApp,
		signerFactory:     signerFactory,
		accountBuilder:    accountBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable communications.Communication) (stacks.Assignable, *uint, error) {
	if assignable.IsSign() {
		sign := assignable.Sign()
		return app.execSignApp.Execute(frame, sign)
	}

	if assignable.IsVote() {
		vote := assignable.Vote()
		return app.execVoteApp.Execute(frame, vote)
	}

	builder := app.assignableBuilder.Create()
	if assignable.IsGenerateRing() {
		ring := []signers.PublicKey{}
		amountVariable := assignable.GenerateRing()
		pAmount, err := frame.FetchUnsignedInt(amountVariable)
		if err != nil {
			code := failures.CouldNotFetchGenerateRingFromFrame
			return nil, &code, err
		}

		casted := int(*pAmount)
		for i := 0; i < casted; i++ {
			pubKey := app.signerFactory.Create().PublicKey()
			ring = append(ring, pubKey)
		}

		account, err := app.accountBuilder.Create().WithRing(ring).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithAccount(account)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
