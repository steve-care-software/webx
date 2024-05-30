package origins

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/origins"
	json_operators "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins/operators"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins/resources"
)

// Adapter represents the adapter
type Adapter struct {
	resourceAdapter *json_resources.Adapter
	operatorAdapter *json_operators.Adapter
	valueBuilder    origins.ValueBuilder
	builder         origins.Builder
}

func createAdapter(
	resourceAdapter *json_resources.Adapter,
	operatorAdapter *json_operators.Adapter,
	valueBuilder origins.ValueBuilder,
	builder origins.Builder,
) origins.Adapter {
	out := Adapter{
		resourceAdapter: resourceAdapter,
		operatorAdapter: operatorAdapter,
		valueBuilder:    valueBuilder,
		builder:         builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins origins.Origin) ([]byte, error) {
	ptr, err := app.OriginToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (origins.Origin, error) {
	ins := new(Origin)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToOrigin(*ins)
}

// OriginToStruct converts an origin to struct
func (app *Adapter) OriginToStruct(ins origins.Origin) (*Origin, error) {
	ptrResource, err := app.resourceAdapter.ResourceToStruct(ins.Resource())
	if err != nil {
		return nil, err
	}

	ptrOperator, err := app.operatorAdapter.OperatorToStruct(ins.Operator())
	if err != nil {
		return nil, err
	}

	ptrNext, err := app.ValueToStruct(ins.Next())
	if err != nil {
		return nil, err
	}

	return &Origin{
		Resource: *ptrResource,
		Operator: *ptrOperator,
		Next:     *ptrNext,
	}, nil
}

// StructToOrigin converts a struct to orgin
func (app *Adapter) StructToOrigin(str Origin) (origins.Origin, error) {
	resource, err := app.resourceAdapter.StructToResource(str.Resource)
	if err != nil {
		return nil, err
	}

	operator, err := app.operatorAdapter.StructToOperator(str.Operator)
	if err != nil {
		return nil, err
	}

	next, err := app.StructToValue(str.Next)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithResource(resource).
		WithOperator(operator).
		WithNext(next).
		Now()
}

// ValueToStruct converts a value to struct
func (app *Adapter) ValueToStruct(ins origins.Value) (*Value, error) {
	output := Value{}
	if ins.IsOrigin() {
		ptr, err := app.OriginToStruct(ins.Origin())
		if err != nil {
			return nil, err
		}

		output.Origin = ptr
	}

	if ins.IsResource() {
		ptr, err := app.resourceAdapter.ResourceToStruct(ins.Resource())
		if err != nil {
			return nil, err
		}

		output.Resource = ptr
	}

	return &output, nil
}

// StructToValue converts a struct to value
func (app *Adapter) StructToValue(str Value) (origins.Value, error) {
	builder := app.valueBuilder.Create()
	if str.Origin != nil {
		ins, err := app.StructToOrigin(*str.Origin)
		if err != nil {
			return nil, err
		}

		builder.WithOrigin(ins)
	}

	if str.Resource != nil {
		ins, err := app.resourceAdapter.StructToResource(*str.Resource)
		if err != nil {
			return nil, err
		}

		builder.WithResource(ins)
	}

	return builder.Now()
}
