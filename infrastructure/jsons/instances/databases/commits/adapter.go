package commits

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions"
)

// Adapter represents the commit adapter
type Adapter struct {
	actionAdapter *json_actions.Adapter
	builder       commits.Builder
	hashAdapter   hash.Adapter
}

func createAdapter(
	actionAdapter *json_actions.Adapter,
	builder commits.Builder,
	hashAdapter hash.Adapter,
) commits.Adapter {
	out := Adapter{
		actionAdapter: actionAdapter,
		builder:       builder,
		hashAdapter:   hashAdapter,
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
	ptr, err := app.actionAdapter.ActionsToStruct(ins.Content().Actions())
	if err != nil {
		return nil, err
	}

	content := ins.Content()
	out := Commit{
		Description: content.Description(),
		Actions:     ptr,
	}

	if ins.HasParent() {
		out.Parent = ins.Parent().String()
	}

	return &out, nil
}

// StructToCommit converts a struct to commit
func (app *Adapter) StructToCommit(str Commit) (commits.Commit, error) {
	ins, err := app.actionAdapter.StructToActions(str.Actions)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().
		WithDescription(str.Description).
		WithActions(ins)

	if str.Parent != "" {
		pHash, err := app.hashAdapter.FromString(str.Parent)
		if err != nil {
			return nil, err
		}

		builder.WithParent(*pHash)
	}

	return builder.Now()
}
