package pointers

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers"
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/conditions"
)

// Adapter represents a pointer adapter
type Adapter struct {
	conditionAdapter *json_conditions.Adapter
	builder          pointers.Builder
	pointerBuilder   pointers.PointerBuilder
}

func createAdapter(
	conditionAdapter *json_conditions.Adapter,
	builder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
) pointers.Adapter {
	out := Adapter{
		conditionAdapter: conditionAdapter,
		builder:          builder,
		pointerBuilder:   pointerBuilder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins pointers.Pointers) ([]byte, error) {
	ptr, err := app.PointersToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (pointers.Pointers, error) {
	ins := new([]Pointer)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToPointers(*ins)
}

// PointersToStruct converts a pointers to struct
func (app *Adapter) PointersToStruct(ins pointers.Pointers) ([]Pointer, error) {
	output := []Pointer{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.PointerToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToPointers converts a struct to pointers
func (app *Adapter) StructToPointers(str []Pointer) (pointers.Pointers, error) {
	list := []pointers.Pointer{}
	for _, oneStruct := range str {
		ins, err := app.StructToPointer(oneStruct)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}

// PointerToStruct converts a pointer to struct
func (app *Adapter) PointerToStruct(ins pointers.Pointer) (*Pointer, error) {
	output := Pointer{
		Path:     ins.Path(),
		IsActive: ins.IsActive(),
	}

	if ins.HasLoader() {
		ptr, err := app.conditionAdapter.ConditionToStruct(ins.Loader())
		if err != nil {
			return nil, err
		}

		output.Loader = ptr
	}

	if ins.HasCanceller() {
		ptr, err := app.conditionAdapter.ConditionToStruct(ins.Canceller())
		if err != nil {
			return nil, err
		}

		output.Canceller = ptr
	}

	return &output, nil
}

// StructToPointer converts a struct to pointer
func (app *Adapter) StructToPointer(str Pointer) (pointers.Pointer, error) {
	builder := app.pointerBuilder.Create().WithPath(str.Path)
	if str.IsActive {
		builder.IsActive()
	}

	if str.Loader != nil {
		ins, err := app.conditionAdapter.StructToCondition(*str.Loader)
		if err != nil {
			return nil, err
		}

		builder.WithLoader(ins)
	}

	if str.Canceller != nil {
		ins, err := app.conditionAdapter.StructToCondition(*str.Canceller)
		if err != nil {
			return nil, err
		}

		builder.WithCanceller(ins)
	}

	return builder.Now()
}
