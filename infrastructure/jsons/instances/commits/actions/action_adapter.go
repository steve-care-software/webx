package actions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
	json_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/pointers"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/resources"
)

// ActionAdapter represents an action adapter
type ActionAdapter struct {
	resourceAdapter *json_resources.Adapter
	pointerAdapter  *json_pointers.Adapter
	builder         actions.ActionBuilder
}

func createActionAdapter(
	resourceAdapter *json_resources.Adapter,
	pointerAdapter *json_pointers.Adapter,
	builder actions.ActionBuilder,
) actions.ActionAdapter {
	out := ActionAdapter{
		resourceAdapter: resourceAdapter,
		pointerAdapter:  pointerAdapter,
		builder:         builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *ActionAdapter) ToBytes(ins actions.Action) ([]byte, error) {
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
func (app *ActionAdapter) ToInstance(bytes []byte) (actions.Action, error) {
	ins := new(Action)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToAction(*ins)
}

// ActionToStruct converts an action to struct
func (app *ActionAdapter) ActionToStruct(ins actions.Action) (*Action, error) {
	output := Action{}
	if ins.HasInsert() {
		ptrStr, err := app.resourceAdapter.ResourceToStruct(ins.Insert())
		if err != nil {
			return nil, err
		}

		output.Insert = ptrStr
	}

	if ins.HasDelete() {
		retStr := app.pointerAdapter.PointerToStruct(ins.Delete())
		output.Delete = &retStr
	}

	return &output, nil
}

// StructToAction converts a struct to action
func (app *ActionAdapter) StructToAction(str Action) (actions.Action, error) {
	builder := app.builder.Create()
	if str.Insert != nil {
		retIns, err := app.resourceAdapter.StructToResource(*str.Insert)
		if err != nil {
			return nil, err
		}

		builder.WithInsert(retIns)
	}

	if str.Delete != nil {
		retIns, err := app.pointerAdapter.StructToPointer(*str.Delete)
		if err != nil {
			return nil, err
		}

		builder.WithDelete(retIns)
	}

	return builder.Now()
}
