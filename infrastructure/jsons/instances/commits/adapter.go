package commits

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/commits"
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions"
)

// Adapter represents a commit adapter
type Adapter struct {
	actionsAdapter   *json_actions.Adapter
	signatureAdapter signers.SignatureAdapter
	contentBuilder   commits.ContentBuilder
	builder          commits.Builder
}

func createAdapter(
	actionsAdapter *json_actions.Adapter,
	signatureAdapter signers.SignatureAdapter,
	contentBuilder commits.ContentBuilder,
	builder commits.Builder,
) commits.Adapter {
	out := Adapter{
		actionsAdapter,
		signatureAdapter,
		contentBuilder,
		builder,
	}
	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins commits.Commit) ([]byte, error) {
	ptr, err := app.CommitToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (commits.Commit, error) {
	ins := new(Commit)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCommit(*ins)
}

// CommitToStruct converts a commit to struct
func (app *Adapter) CommitToStruct(ins commits.Commit) (*Commit, error) {
	contentPtr, err := app.contentToStruct(ins.Content())
	if err != nil {
		return nil, err
	}

	return &Commit{
		Content:   *contentPtr,
		Signature: ins.Signature().String(),
	}, nil
}

// StructToCommit converts a struct to commit
func (app *Adapter) StructToCommit(str Commit) (commits.Commit, error) {
	content, err := app.structToContent(str.Content)
	if err != nil {
		return nil, err
	}

	signature, err := app.signatureAdapter.ToSignature(str.Signature)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithContent(content).
		WithSignature(signature).
		Now()
}

func (app *Adapter) contentToStruct(ins commits.Content) (*Content, error) {
	actionsStr, err := app.actionsAdapter.ActionsToStructs(ins.Actions())
	if err != nil {
		return nil, err
	}

	output := Content{
		Actions: actionsStr,
	}

	if ins.HasPrevious() {
		prevPtr, err := app.CommitToStruct(ins.Previous())
		if err != nil {
			return nil, err
		}

		output.Previous = prevPtr
	}

	return &output, nil
}

func (app *Adapter) structToContent(str Content) (commits.Content, error) {
	actionsIns, err := app.actionsAdapter.StructsToActions(str.Actions)
	if err != nil {
		return nil, err
	}

	builder := app.contentBuilder.Create().WithActions(actionsIns)
	if str.Previous != nil {
		prevIns, err := app.StructToCommit(*str.Previous)
		if err != nil {
			return nil, err
		}

		builder.WithPrevious(prevIns)
	}

	return builder.Now()
}
