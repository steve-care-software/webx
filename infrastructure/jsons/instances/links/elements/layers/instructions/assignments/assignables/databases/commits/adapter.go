package commits

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/commits"
)

// Adapter represents an adapter
type Adapter struct {
	builder commits.Builder
}

func createAdapter(
	builder commits.Builder,
) commits.Adapter {
	out := Adapter{
		builder: builder,
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
	out := Commit{
		Description: ins.Description(),
		Actions:     ins.Actions(),
	}

	if ins.HashParent() {
		out.Parent = ins.Parent()
	}

	return &out, nil
}

// StructToCommit converts a struct to commit
func (app *Adapter) StructToCommit(str Commit) (commits.Commit, error) {
	builder := app.builder.Create().
		WithDescription(str.Description).
		WithActions(str.Actions)

	if str.Parent != "" {
		builder.WithParent(str.Parent)
	}

	return builder.Now()
}
