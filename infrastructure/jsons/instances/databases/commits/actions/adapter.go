package actions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	json_modifications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions/modifications"
)

// Adapter represents the adapter
type Adapter struct {
	modificationAdapter *json_modifications.Adapter
	actionBuilder       actions.ActionBuilder
	builder             actions.Builder
}

func createAdapter(
	modificationAdapter *json_modifications.Adapter,
	actionBuilder actions.ActionBuilder,
	builder actions.Builder,
) actions.Adapter {
	out := Adapter{
		modificationAdapter: modificationAdapter,
		actionBuilder:       actionBuilder,
		builder:             builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins actions.Actions) ([]byte, error) {
	ptr, err := app.ActionsToStruct(ins)
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

	return app.StructToActions(*ins)
}

// ActionsToStruct converts an actions to struct
func (app *Adapter) ActionsToStruct(ins actions.Actions) ([]Action, error) {
	out := []Action{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.ActionToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructToActions converts a struct to actions
func (app *Adapter) StructToActions(list []Action) (actions.Actions, error) {
	output := []actions.Action{}
	for _, oneStr := range list {
		ins, err := app.StructToAction(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// ActionToStruct converts an action to struct
func (app *Adapter) ActionToStruct(ins actions.Action) (*Action, error) {
	str, err := app.modificationAdapter.ModificationsToStruct(ins.Modifications())
	if err != nil {
		return nil, err
	}

	return &Action{
		Path:          ins.Path(),
		Modifications: str,
	}, nil
}

// StructToAction converts a struct to action
func (app *Adapter) StructToAction(str Action) (actions.Action, error) {
	ins, err := app.modificationAdapter.StructToModifications(str.Modifications)
	if err != nil {
		return nil, err
	}

	return app.actionBuilder.Create().
		WithPath(str.Path).
		WithModifications(ins).
		Now()
}
