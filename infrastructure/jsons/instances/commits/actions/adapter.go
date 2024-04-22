package actions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

// Adapter represents an adapter
type Adapter struct {
	actionAdapter *ActionAdapter
	builder       actions.Builder
}

func createAdapter(
	actionAdapter *ActionAdapter,
	builder actions.Builder,
) actions.Adapter {
	out := Adapter{
		actionAdapter: actionAdapter,
		builder:       builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins actions.Actions) ([]byte, error) {
	ptr, err := app.ActionsToStructs(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (actions.Actions, error) {
	ins := new([]Action)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructsToActions(*ins)
}

// ActionsToStructs converts actions to structs
func (app *Adapter) ActionsToStructs(ins actions.Actions) ([]Action, error) {
	output := []Action{}
	list := ins.List()
	for _, oneAction := range list {
		ptrStr, err := app.actionAdapter.ActionToStruct(oneAction)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptrStr)
	}

	return output, nil
}

// StructsToActions converts structs to actions
func (app *Adapter) StructsToActions(str []Action) (actions.Actions, error) {
	list := []actions.Action{}
	for _, oneStr := range str {
		retIns, err := app.actionAdapter.StructToAction(oneStr)
		if err != nil {
			return nil, err
		}

		list = append(list, retIns)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}
