package operators

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
)

// Adapter represents the adapter
type Adapter struct {
	builder operators.Builder
}

func createAdapter(
	builder operators.Builder,
) operators.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
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
	if ins.IsAnd() {
		output.And = true
	}

	if ins.IsOr() {
		output.Or = true
	}

	if ins.IsXor() {
		output.Xor = true
	}

	return &output, nil
}

// StructToOperator converts a struct to operator
func (app *Adapter) StructToOperator(str Operator) (operators.Operator, error) {
	builder := app.builder.Create()
	if str.And {
		builder.IsAnd()
	}

	if str.Or {
		builder.IsOr()
	}

	if str.Xor {
		builder.IsXor()
	}

	return builder.Now()
}
