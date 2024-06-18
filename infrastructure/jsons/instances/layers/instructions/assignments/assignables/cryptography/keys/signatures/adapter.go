package signatures

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	json_signs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	json_votes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
)

// Adapter represents the adapter
type Adapter struct {
	signAdapter *json_signs.Adapter
	voteAdapter *json_votes.Adapter
	builder     signatures.Builder
}

func createAdapter(
	signAdapter *json_signs.Adapter,
	voteAdapter *json_votes.Adapter,
	builder signatures.Builder,
) signatures.Adapter {
	out := Adapter{
		signAdapter: signAdapter,
		voteAdapter: voteAdapter,
		builder:     builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins signatures.Signature) ([]byte, error) {
	ptr, err := app.SignatureToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (signatures.Signature, error) {
	ins := new(Signature)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToSignature(*ins)
}

// SignatureToStruct converts a signature to struct
func (app *Adapter) SignatureToStruct(ins signatures.Signature) (*Signature, error) {
	out := Signature{}
	if ins.IsGeneratePrivateKey() {
		out.IsGeneratePrivateKey = true
	}

	if ins.IsFetchPublicKey() {
		out.FetchPublicKey = ins.FetchPublicKey()
	}

	if ins.IsSign() {
		ptr, err := app.signAdapter.SignToStruct(ins.Sign())
		if err != nil {
			return nil, err
		}

		out.Sign = ptr
	}

	if ins.IsVote() {
		ptr, err := app.voteAdapter.VoteToStruct(ins.Vote())
		if err != nil {
			return nil, err
		}

		out.Vote = ptr
	}

	return &out, nil
}

// StructToSignature converts a struct to signature
func (app *Adapter) StructToSignature(str Signature) (signatures.Signature, error) {
	builder := app.builder.Create()
	if str.IsGeneratePrivateKey {
		builder.IsGeneratePrivateKey()
	}

	if str.FetchPublicKey != "" {
		builder.WithFetchPublicKey(str.FetchPublicKey)
	}

	if str.Sign != nil {
		ins, err := app.signAdapter.StructToSign(*str.Sign)
		if err != nil {
			return nil, err
		}

		builder.WithSign(ins)
	}

	if str.Vote != nil {
		ins, err := app.voteAdapter.StructToVote(*str.Vote)
		if err != nil {
			return nil, err
		}

		builder.WithVote(ins)
	}

	return builder.Now()
}
