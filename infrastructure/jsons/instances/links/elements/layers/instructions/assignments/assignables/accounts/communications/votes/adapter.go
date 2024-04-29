package votes

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// Adapter represents an adapter
type Adapter struct {
	builder votes.Builder
}

func createAdapter(
	builder votes.Builder,
) votes.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins votes.Vote) ([]byte, error) {
	ptr, err := app.VoteToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (votes.Vote, error) {
	ins := new(Vote)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToVote(*ins)
}

// VoteToStruct converts a vote to struct
func (app *Adapter) VoteToStruct(ins votes.Vote) (*Vote, error) {
	return &Vote{
		Message: ins.Message(),
		Ring:    ins.Ring(),
		Account: ins.Account(),
	}, nil
}

// StructToVote converts a struct to vote
func (app *Adapter) StructToVote(str Vote) (votes.Vote, error) {
	return app.builder.Create().
		WithMessage(str.Message).
		WithRing(str.Ring).
		WithAccount(str.Account).
		Now()
}
