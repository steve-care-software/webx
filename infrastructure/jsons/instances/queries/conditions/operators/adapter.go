package operators

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	json_integers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/queries/conditions/operators/integers"
	json_relationals "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/queries/conditions/operators/relationals"
)

// Adapter represents an adapter
type Adapter struct {
	integerAdapter    *json_integers.Adapter
	relationalAdapter *json_relationals.Adapter
	builder           operators.Builder
}

func createAdapter(
	integerAdapter *json_integers.Adapter,
	relationalAdapter *json_relationals.Adapter,
	builder operators.Builder,
) operators.Adapter {
	out := Adapter{
		integerAdapter:    integerAdapter,
		relationalAdapter: relationalAdapter,
		builder:           builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins operators.Operator) ([]byte, error) {
	ptr, err := app.OperatorToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (operators.Operator, error) {
	ins := new(Operator)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToOperator(*ins)
}

// OperatorToStruct converts an operator to struct
func (app *Adapter) OperatorToStruct(ins operators.Operator) (*Operator, error) {
	output := Operator{}
	if ins.IsEqual() {
		output.IsEqual = true
	}

	if ins.IsInteger() {
		ptr, err := app.integerAdapter.IntegerToStruct(ins.Integer())
		if err != nil {
			return nil, err
		}

		output.Integer = ptr
	}

	if ins.IsRelational() {
		ptr, err := app.relationalAdapter.RelationalToStruct(ins.Relational())
		if err != nil {
			return nil, err
		}

		output.Relational = ptr
	}

	return &output, nil
}

// StructToOperator converts a struct to operator
func (app *Adapter) StructToOperator(str Operator) (operators.Operator, error) {
	builder := app.builder.Create()
	if str.IsEqual {
		builder.IsEqual()
	}

	if str.Integer != nil {
		ins, err := app.integerAdapter.StructToInteger(*str.Integer)
		if err != nil {
			return nil, err
		}

		builder.WithInteger(ins)
	}

	if str.Relational != nil {
		ins, err := app.relationalAdapter.StructToRelational(*str.Relational)
		if err != nil {
			return nil, err
		}

		builder.WithRelational(ins)
	}

	return builder.Now()
}
