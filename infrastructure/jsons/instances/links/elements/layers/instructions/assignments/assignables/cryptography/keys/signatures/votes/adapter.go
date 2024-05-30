package votes

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	json_creates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	json_validates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// Adapter represents an adapter
type Adapter struct {
	createAdapter   *json_creates.Adapter
	validateAdapter *json_validates.Adapter
	builder         votes.Builder
}

func createAdapter(
	createAdapter *json_creates.Adapter,
	validateAdapter *json_validates.Adapter,
	builder votes.Builder,
) votes.Adapter {
	out := Adapter{
		createAdapter:   createAdapter,
		validateAdapter: validateAdapter,
		builder:         builder,
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
	out := Vote{}
	if ins.IsCreate() {
		ptr, err := app.createAdapter.CreateToStruct(ins.Create())
		if err != nil {
			return nil, err
		}

		out.Create = ptr
	}

	if ins.IsValidate() {
		ptr, err := app.validateAdapter.ValidateToStruct(ins.Validate())
		if err != nil {
			return nil, err
		}

		out.Validate = ptr
	}

	return &out, nil
}

// StructToVote converts a struct to vote
func (app *Adapter) StructToVote(str Vote) (votes.Vote, error) {
	builder := app.builder.Create()
	if str.Create != nil {
		ins, err := app.createAdapter.StructToCreate(*str.Create)
		if err != nil {
			return nil, err
		}

		builder.WithCreate(ins)
	}

	if str.Validate != nil {
		ins, err := app.validateAdapter.StructToValidate(*str.Validate)
		if err != nil {
			return nil, err
		}

		builder.WithValidate(ins)
	}

	return builder.Now()
}
