package votes

import (
	"math/rand"
	"time"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	accountBuilder    stacks_accounts.Builder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	accountBuilder stacks_accounts.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		accountBuilder:    accountBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, *uint, error) {
	messageVar := assignable.Message()
	message, err := frame.FetchBytes(messageVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	ringVar := assignable.Ring()
	ring, err := frame.FetchRing(ringVar)
	if err != nil {
		code := failures.CouldNotFetchRingFromFrame
		return nil, &code, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		code := failures.CouldNotFetchAccountFromFrame
		return nil, &code, err
	}

	// add the signer public key to the ring:
	signer := account.Signer()
	ring = append(ring, signer.PublicKey())

	// shuffle using fisher-yates algo:
	rand.Seed(time.Now().UnixNano())
	for i := len(ring) - 1; i > 0; i-- {
		rdn := rand.Intn(i + 1)
		ring[i], ring[rdn] = ring[rdn], ring[i]
	}

	// vote:
	vote, err := signer.Vote(string(message), ring)
	if err != nil {
		return nil, nil, err
	}

	retStackAccount, err := app.accountBuilder.Create().
		WithVote(vote).
		Now()

	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithAccount(retStackAccount).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil

}
