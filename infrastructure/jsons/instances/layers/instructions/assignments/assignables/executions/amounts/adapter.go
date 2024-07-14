package amounts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
)

// Adapter represents the amount adapter
type Adapter struct {
	builder amounts.Builder
}

func createAdapter(
	builder amounts.Builder,
) amounts.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins amounts.Amount) ([]byte, error) {
	str := app.AmountToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (amounts.Amount, error) {
	ins := new(Amount)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToAmount(*ins)
}

// AmountToStruct converts an amount to struct
func (app *Adapter) AmountToStruct(ins amounts.Amount) Amount {
	return Amount{
		Context: ins.Context(),
		Return:  ins.Return(),
	}
}

// StructToAmount converts a struct to amounts
func (app *Adapter) StructToAmount(str Amount) (amounts.Amount, error) {
	return app.builder.Create().
		WithContext(str.Context).
		WithReturn(str.Return).
		Now()
}
