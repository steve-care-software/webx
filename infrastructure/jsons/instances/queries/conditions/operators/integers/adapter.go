package integers

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/integers"
)

// Adapter represents the integer adapter
type Adapter struct {
	builder integers.Builder
}

func createAdapter(
	builder integers.Builder,
) integers.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins integers.Integer) ([]byte, error) {
	ptr, err := app.IntegerToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to insatnce
func (app *Adapter) ToInstance(bytes []byte) (integers.Integer, error) {
	ins := new(Integer)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInteger(*ins)
}

// IntegerToStruct converts an integer to struct
func (app *Adapter) IntegerToStruct(ins integers.Integer) (*Integer, error) {
	output := Integer{}
	if ins.IsBiggerThan() {
		output.IsBiggerThan = true
	}

	if ins.IsSmallerThan() {
		output.IsSmalerThan = true
	}

	if ins.IsEqual() {
		output.IsEqual = true
	}

	return &output, nil
}

// StructToInteger converts a struct to integer
func (app *Adapter) StructToInteger(str Integer) (integers.Integer, error) {
	builder := app.builder.Create()
	if str.IsBiggerThan {
		builder.IsBiggerThan()
	}

	if str.IsSmalerThan {
		builder.IsSmallerThan()
	}

	if str.IsEqual {
		builder.IsEqual()
	}

	return builder.Now()
}
