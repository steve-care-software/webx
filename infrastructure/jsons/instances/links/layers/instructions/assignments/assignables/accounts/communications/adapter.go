package communications

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	json_signs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/communications/signs"
	json_votes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// Adapter represents an adapter
type Adapter struct {
	signAdapter *json_signs.Adapter
	voteAdapter *json_votes.Adapter
	builder     communications.Builder
}

func createAdapter(
	signAdapter *json_signs.Adapter,
	voteAdapter *json_votes.Adapter,
	builder communications.Builder,
) communications.Adapter {
	out := Adapter{
		signAdapter: signAdapter,
		voteAdapter: voteAdapter,
		builder:     builder,
	}

	return &out
}

// ToBytes convert instance to bytes
func (app *Adapter) ToBytes(ins communications.Communication) ([]byte, error) {
	ptr, err := app.CommunicationToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance convert bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (communications.Communication, error) {
	ins := new(Communication)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCommunication(*ins)
}

// CommunicationToStruct converts a communication to struct
func (app *Adapter) CommunicationToStruct(ins communications.Communication) (*Communication, error) {
	output := Communication{}
	if ins.IsGenerateRing() {
		output.GenerateRing = ins.GenerateRing()
	}

	if ins.IsSign() {
		pSign, err := app.signAdapter.SignToStruct(ins.Sign())
		if err != nil {
			return nil, err
		}

		output.Sign = pSign
	}

	if ins.IsVote() {
		pVote, err := app.voteAdapter.VoteToStruct(ins.Vote())
		if err != nil {
			return nil, err
		}

		output.Vote = pVote
	}

	return &output, nil
}

// StructToCommunication converts a struct to communication
func (app *Adapter) StructToCommunication(str Communication) (communications.Communication, error) {
	builder := app.builder.Create()
	if str.GenerateRing != "" {
		builder.WithGenerateRing(str.GenerateRing)
	}

	if str.Sign != nil {
		sign, err := app.signAdapter.StructToSign(*str.Sign)
		if err != nil {
			return nil, err
		}

		builder.WithSign(sign)
	}

	if str.Vote != nil {
		vote, err := app.voteAdapter.StructToVote(*str.Vote)
		if err != nil {
			return nil, err
		}

		builder.WithVote(vote)
	}

	return builder.Now()
}
