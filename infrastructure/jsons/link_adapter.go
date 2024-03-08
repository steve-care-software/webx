package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/operators"
	origins_resources "github.com/steve-care-software/datastencil/domain/libraries/links/origins/resources"
	structs "github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries/links"
)

type linkAdapter struct {
	hashAdapter              hash.Adapter
	builder                  links.Builder
	linkBuilder              links.LinkBuilder
	elementsBuilder          links.ElementsBuilder
	elementBuilder           links.ElementBuilder
	conditionBuilder         links.ConditionBuilder
	conditionValueBuilder    links.ConditionValueBuilder
	conditionResourceBuilder resources.Builder
	originBuilder            origins.Builder
	originValueBuilder       origins.ValueBuilder
	originResourceBuilder    origins_resources.Builder
	operatorBuilder          operators.Builder
}

func createLinkAdapter(
	hashAdapter hash.Adapter,
	builder links.Builder,
	linkBuilder links.LinkBuilder,
	elementsBuilder links.ElementsBuilder,
	elementBuilder links.ElementBuilder,
	conditionBuilder links.ConditionBuilder,
	conditionValueBuilder links.ConditionValueBuilder,
	conditionResourceBuilder resources.Builder,
	originBuilder origins.Builder,
	originValueBuilder origins.ValueBuilder,
	originResourceBuilder origins_resources.Builder,
	operatorBuilder operators.Builder,
) links.LinkAdapter {
	out := linkAdapter{
		hashAdapter:              hashAdapter,
		builder:                  builder,
		linkBuilder:              linkBuilder,
		elementsBuilder:          elementsBuilder,
		elementBuilder:           elementBuilder,
		conditionBuilder:         conditionBuilder,
		conditionValueBuilder:    conditionValueBuilder,
		conditionResourceBuilder: conditionResourceBuilder,
		originBuilder:            originBuilder,
		originValueBuilder:       originValueBuilder,
		originResourceBuilder:    originResourceBuilder,
		operatorBuilder:          operatorBuilder,
	}

	return &out
}

// ToData converts link to data
func (app *linkAdapter) ToData(ins links.Link) ([]byte, error) {
	str := app.toStructLink(ins)
	data, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ToInstance converts data to link instance
func (app *linkAdapter) ToInstance(data []byte) (links.Link, error) {
	ins := structs.Link{}
	err := json.Unmarshal(data, &ins)
	if err != nil {
		return nil, err
	}

	return app.toInstanceLink(ins)
}

func (app *linkAdapter) toInstanceLink(ins structs.Link) (links.Link, error) {
	origin, err := app.toInstanceOrigin(ins.Origin)
	if err != nil {
		return nil, err
	}

	elements, err := app.toInstanceElements(ins.Elements)
	if err != nil {
		return nil, err
	}

	return app.linkBuilder.Create().WithOrigin(origin).WithElements(elements).Now()
}

func (app *linkAdapter) toStructLink(ins links.Link) structs.Link {
	return structs.Link{
		Origin:   app.toStructOrigin(ins.Origin()),
		Elements: app.toStructElements(ins.Elements()),
	}
}

func (app *linkAdapter) toInstanceElements(list []structs.Element) (links.Elements, error) {
	output := []links.Element{}
	for _, oneElement := range list {
		ins, err := app.toInstanceElement(oneElement)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.elementsBuilder.Create().
		WithList(output).
		Now()
}

func (app *linkAdapter) toStructElements(ins links.Elements) []structs.Element {
	output := []structs.Element{}
	list := ins.List()
	for _, oneElement := range list {
		output = append(output, app.toStructElement(oneElement))
	}

	return output
}

func (app *linkAdapter) toInstanceElement(str structs.Element) (links.Element, error) {
	pHash, err := app.hashAdapter.FromString(str.Layer)
	if err != nil {
		return nil, err
	}

	builder := app.elementBuilder.Create().WithLayer(*pHash)
	if str.Condition != nil {
		condition, err := app.toInstanceCondition(*str.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *linkAdapter) toStructElement(ins links.Element) structs.Element {
	output := structs.Element{
		Layer: ins.Layer().String(),
	}

	if ins.HasCondition() {
		condition := app.toStructCondition(ins.Condition())
		output.Condition = &condition
	}

	return output
}

func (app *linkAdapter) toInstanceCondition(str structs.Condition) (links.Condition, error) {
	resource, err := app.toInstanceConditionResource(str.Resource)
	if err != nil {
		return nil, err
	}

	builder := app.conditionBuilder.Create().WithResource(resource)
	if str.Next != nil {
		value, err := app.toInstanceConditionValue(*str.Next)
		if err != nil {
			return nil, err
		}

		builder.WithNext(value)
	}

	return builder.Now()
}

func (app *linkAdapter) toStructCondition(ins links.Condition) structs.Condition {
	output := structs.Condition{
		Resource: app.toStructConditionResource(ins.Resource()),
	}

	if ins.HasNext() {
		value := app.toStructConditionValue(ins.Next())
		output.Next = &value
	}

	return output
}

func (app *linkAdapter) toInstanceConditionValue(str structs.ConditionValue) (links.ConditionValue, error) {
	builder := app.conditionValueBuilder.Create()
	if str.Resource != nil {
		resource, err := app.toInstanceConditionResource(*str.Resource)
		if err != nil {
			return nil, err
		}

		builder.WithResource(resource)
	}

	if str.Condition != nil {
		condition, err := app.toInstanceCondition(*str.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *linkAdapter) toStructConditionValue(ins links.ConditionValue) structs.ConditionValue {
	output := structs.ConditionValue{}
	if ins.IsResource() {
		resource := app.toStructConditionResource(ins.Resource())
		output.Resource = &resource
	}

	if ins.IsCondition() {
		condition := app.toStructCondition(ins.Condition())
		output.Condition = &condition
	}

	return output
}

func (app *linkAdapter) toInstanceConditionResource(str structs.ConditionResource) (resources.Resource, error) {
	builder := app.conditionResourceBuilder.Create().WithCode(str.Code)
	if str.IsRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	return builder.Now()
}

func (app *linkAdapter) toStructConditionResource(ins resources.Resource) structs.ConditionResource {
	return structs.ConditionResource{
		Code:            ins.Code(),
		IsRaisedInLayer: ins.IsRaisedInLayer(),
	}
}

func (app *linkAdapter) toInstanceOrigin(str structs.Origin) (origins.Origin, error) {
	resource, err := app.toInstanceOriginResource(str.Resource)
	if err != nil {
		return nil, err
	}

	operator, err := app.toInstanceOperator(str.Operator)
	if err != nil {
		return nil, err
	}

	next, err := app.toInstanceOriginValue(str.Next)
	if err != nil {
		return nil, err
	}

	return app.originBuilder.Create().
		WithResource(resource).
		WithOperator(operator).
		WithNext(next).
		Now()
}

func (app *linkAdapter) toStructOrigin(ins origins.Origin) structs.Origin {
	return structs.Origin{
		Resource: app.toStructOriginResource(ins.Resource()),
		Operator: app.toStructOperator(ins.Operator()),
		Next:     app.toStructOriginValue(ins.Next()),
	}
}

func (app *linkAdapter) toInstanceOriginValue(str structs.OriginValue) (origins.Value, error) {
	builder := app.originValueBuilder.Create()
	if str.Resource != nil {
		originValue, err := app.toInstanceOriginResource(*str.Resource)
		if err != nil {
			return nil, err
		}

		builder.WithResource(originValue)
	}

	if str.Origin != nil {
		origin, err := app.toInstanceOrigin(*str.Origin)
		if err != nil {
			return nil, err
		}

		builder.WithOrigin(origin)
	}

	return builder.Now()
}

func (app *linkAdapter) toStructOriginValue(ins origins.Value) structs.OriginValue {
	output := structs.OriginValue{}
	if ins.IsResource() {
		resource := app.toStructOriginResource(ins.Resource())
		output.Resource = &resource
	}

	if ins.IsOrigin() {
		origin := app.toStructOrigin(ins.Origin())
		output.Origin = &origin
	}

	return output
}

func (app *linkAdapter) toInstanceOriginResource(str structs.Resource) (origins_resources.Resource, error) {
	pHash, err := app.hashAdapter.FromString(str.Layer)
	if err != nil {
		return nil, err
	}

	builder := app.originResourceBuilder.Create().WithLayer(*pHash)
	if str.IsMandatory {
		builder.IsMandatory()
	}

	return builder.Now()
}

func (app *linkAdapter) toStructOriginResource(ins origins_resources.Resource) structs.Resource {
	return structs.Resource{
		Layer:       ins.Layer().String(),
		IsMandatory: ins.IsMandatory(),
	}
}

func (app *linkAdapter) toInstanceOperator(str structs.Operator) (operators.Operator, error) {
	builder := app.operatorBuilder.Create()
	if str.IsAnd {
		builder.IsAnd()
	}

	if str.IsOr {
		builder.IsOr()
	}

	if str.IsXor {
		builder.IsXor()
	}

	return builder.Now()
}

func (app *linkAdapter) toStructOperator(ins operators.Operator) structs.Operator {
	return structs.Operator{
		IsAnd: ins.IsAnd(),
		IsOr:  ins.IsOr(),
		IsXor: ins.IsXor(),
	}
}
