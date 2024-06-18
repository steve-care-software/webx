package conditions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
	json_operators "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/conditions/operators"
)

// Adapter represents an adapter
type Adapter struct {
	operatorAdapter    *json_operators.Adapter
	builder            conditions.Builder
	resourceBuilder    conditions.ResourceBuilder
	comparisonsBuilder conditions.ComparisonsBuilder
	comparisonBuilder  conditions.ComparisonBuilder
}

func createAdapter(
	operatorAdapter *json_operators.Adapter,
	builder conditions.Builder,
	resourceBuilder conditions.ResourceBuilder,
	comparisonsBuilder conditions.ComparisonsBuilder,
	comparisonBuilder conditions.ComparisonBuilder,
) conditions.Adapter {
	out := Adapter{
		operatorAdapter:    operatorAdapter,
		builder:            builder,
		resourceBuilder:    resourceBuilder,
		comparisonsBuilder: comparisonsBuilder,
		comparisonBuilder:  comparisonBuilder,
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

// ToInstance converts bytes to instance
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
	ptrResource, err := app.ResourceToStruct(ins.Resource())
	if err != nil {
		return nil, err
	}

	output := Condition{
		Resource: *ptrResource,
	}

	if ins.HasComparisons() {
		ptrComparisons, err := app.ComparisonsToStruct(ins.Comparisons())
		if err != nil {
			return nil, err
		}

		output.Comparisons = *&ptrComparisons
	}

	return &output, nil
}

// StructToCondition converts a struct to condition
func (app *Adapter) StructToCondition(str Condition) (conditions.Condition, error) {
	resourceIns, err := app.StructToResource(str.Resource)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithResource(resourceIns)
	if str.Comparisons != nil && len(str.Comparisons) > 0 {
		comparisonsIns, err := app.StructToComparisons(str.Comparisons)
		if err != nil {
			return nil, err
		}

		builder.WithComparisons(comparisonsIns)
	}

	return builder.Now()
}

// ResourceToStruct converts a resource to struct
func (app *Adapter) ResourceToStruct(ins conditions.Resource) (*Resource, error) {
	return &Resource{
		Path:         ins.Path(),
		MustBeLoaded: ins.MustBeLoaded(),
	}, nil
}

// StructToResource converts a struct to resource
func (app *Adapter) StructToResource(str Resource) (conditions.Resource, error) {
	builder := app.resourceBuilder.Create().WithPath(str.Path)
	if str.MustBeLoaded {
		builder.MustBeLoaded()
	}

	return builder.Now()
}

// ComparisonsToStruct converts a comparisons to struct
func (app *Adapter) ComparisonsToStruct(ins conditions.Comparisons) ([]Comparison, error) {
	output := []Comparison{}
	list := ins.List()
	for _, oneComparison := range list {
		ptr, err := app.ComparisonToStruct(oneComparison)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToComparisons converts a struct to comparisons
func (app *Adapter) StructToComparisons(str []Comparison) (conditions.Comparisons, error) {
	list := []conditions.Comparison{}
	for _, oneStruct := range str {
		ins, err := app.StructToComparison(oneStruct)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.comparisonsBuilder.Create().
		WithList(list).
		Now()
}

// ComparisonToStruct converts a comparison to struct
func (app *Adapter) ComparisonToStruct(ins conditions.Comparison) (*Comparison, error) {
	ptrOperator, err := app.operatorAdapter.OperatorToStruct(ins.Operator())
	if err != nil {
		return nil, err
	}

	ptrCondition, err := app.ConditionToStruct(ins.Condition())
	if err != nil {
		return nil, err
	}

	return &Comparison{
		Operator:  *ptrOperator,
		Condition: *ptrCondition,
	}, nil
}

// StructToComparison converts a struct to comparison
func (app *Adapter) StructToComparison(str Comparison) (conditions.Comparison, error) {
	operatorIns, err := app.operatorAdapter.StructToOperator(str.Operator)
	if err != nil {
		return nil, err
	}

	conditionIns, err := app.StructToCondition(str.Condition)
	if err != nil {
		return nil, err
	}

	return app.comparisonBuilder.Create().
		WithOperator(operatorIns).
		WithCondition(conditionIns).
		Now()
}
