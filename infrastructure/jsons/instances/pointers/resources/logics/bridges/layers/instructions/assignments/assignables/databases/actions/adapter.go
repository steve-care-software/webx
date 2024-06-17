package actions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/actions"
)

// Adapter represents the adapter
type Adapter struct {
	builder actions.Builder
}

func createAdapter(
	builder actions.Builder,
) actions.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins actions.Action) ([]byte, error) {
	ptr, err := app.ActionToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (actions.Action, error) {
	ins := new(Action)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToAction(*ins)
}

// ActionToStruct converts an action to struct
func (app *Adapter) ActionToStruct(ins actions.Action) (*Action, error) {
	return &Action{
		Path:          ins.Path(),
		Modifications: ins.Modifications(),
	}, nil
}

// StructToAction converts a struct to action
func (app *Adapter) StructToAction(str Action) (actions.Action, error) {
	return app.builder.Create().
		WithPath(str.Path).
		WithModifications(str.Modifications).
		Now()
}
