package conditions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions/resources"
)

// Adapter represents an adapter
type Adapter struct {
	resourceAdapter *json_resources.Adapter
	builder         conditions.Builder
}

func createAdapter(
	resourceAdapter *json_resources.Adapter,
	builder conditions.Builder,
) conditions.Adapter {
	out := Adapter{
		resourceAdapter: resourceAdapter,
		builder:         builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins conditions.Condition) ([]byte, error) {
	ptr, err := app.ConditionToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to condition
func (app *Adapter) ToInstance(bytes []byte) (conditions.Condition, error) {
	ins := new(Condition)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCondition(*ins)
}

// ConditionToStruct converts a condition to struct
func (app *Adapter) ConditionToStruct(ins conditions.Condition) (*Condition, error) {
	ptr, err := app.resourceAdapter.ResourceToStruct(ins.Resource())
	if err != nil {
		return nil, err
	}

	out := Condition{
		Resource: *ptr,
	}

	if ins.HasNext() {
		ptr, err := app.ConditionToStruct(ins.Next())
		if err != nil {
			return nil, err
		}

		out.Next = ptr
	}

	return &out, nil
}

// StructToCondition converts a struct to condition
func (app *Adapter) StructToCondition(str Condition) (conditions.Condition, error) {
	resource, err := app.resourceAdapter.StructToResource(str.Resource)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithResource(resource)
	if str.Next != nil {
		ins, err := app.StructToCondition(*str.Next)
		if err != nil {
			return nil, err
		}

		builder.WithNext(ins)
	}

	return builder.Now()
}
