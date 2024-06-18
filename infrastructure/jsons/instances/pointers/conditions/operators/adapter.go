package operators

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"
)

// Adapter represents the operator adapter
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

// ToBytes convert instance to bytes
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

// ToInstance convert bytes to instance
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
	return &Operator{
		IsAnd: ins.IsAnd(),
		IsOr:  ins.IsOr(),
		IsXor: ins.IsXor(),
	}, nil
}

// StructToOperator converts a struct to operator
func (app *Adapter) StructToOperator(str Operator) (operators.Operator, error) {
	builder := app.builder.Create()
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
