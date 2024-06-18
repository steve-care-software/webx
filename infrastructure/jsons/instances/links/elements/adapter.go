package elements

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions"
)

// Adapter represents the adapter
type Adapter struct {
	conditionAdapter *json_conditions.Adapter
	elementBuilder   elements.ElementBuilder
	builder          elements.Builder
}

func createAdapter(
	conditionAdapter *json_conditions.Adapter,
	elementBuilder elements.ElementBuilder,
	builder elements.Builder,
) elements.Adapter {
	out := Adapter{
		conditionAdapter: conditionAdapter,
		elementBuilder:   elementBuilder,
		builder:          builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins elements.Elements) ([]byte, error) {
	ptr, err := app.ElementsToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (elements.Elements, error) {
	ins := new([]Element)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToElements(*ins)
}

// ElementsToStruct converts elements to struct
func (app *Adapter) ElementsToStruct(ins elements.Elements) ([]Element, error) {
	output := []Element{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.ElementToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToElements converts a struct to elements
func (app *Adapter) StructToElements(list []Element) (elements.Elements, error) {
	output := []elements.Element{}
	for _, oneStr := range list {
		ins, err := app.StructToElement(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// ElementToStruct converts element to struct
func (app *Adapter) ElementToStruct(ins elements.Element) (*Element, error) {
	output := Element{
		Layer: ins.Layer(),
	}

	if ins.HasCondition() {
		ptr, err := app.conditionAdapter.ConditionToStruct(ins.Condition())
		if err != nil {
			return nil, err
		}

		output.Condition = ptr
	}

	return &output, nil
}

// StructToElement converts a struct to element
func (app *Adapter) StructToElement(str Element) (elements.Element, error) {
	builder := app.elementBuilder.Create().WithLayer(str.Layer)
	if str.Condition != nil {
		condition, err := app.conditionAdapter.StructToCondition(*str.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}
