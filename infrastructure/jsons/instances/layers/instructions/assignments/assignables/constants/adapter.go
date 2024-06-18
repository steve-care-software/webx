package constants

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/constants"
)

// Adapter represents an adapter
type Adapter struct {
	builder         constants.Builder
	constantBuilder constants.ConstantBuilder
}

func createAdapter(
	builder constants.Builder,
	constantBuilder constants.ConstantBuilder,
) constants.Adapter {
	out := Adapter{
		builder:         builder,
		constantBuilder: constantBuilder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins constants.Constant) ([]byte, error) {
	ptr, err := app.ConstantToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (constants.Constant, error) {
	ins := new(Constant)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToConstant(*ins)
}

// ConstantsToStruct converts a constant sto struct
func (app *Adapter) ConstantsToStruct(ins constants.Constants) ([]Constant, error) {
	out := []Constant{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.ConstantToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructToConstants converts a struct to constants
func (app *Adapter) StructToConstants(list []Constant) (constants.Constants, error) {
	output := []constants.Constant{}
	for _, oneStr := range list {
		ins, err := app.StructToConstant(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// ConstantToStruct converts a constant to struct
func (app *Adapter) ConstantToStruct(ins constants.Constant) (*Constant, error) {
	out := Constant{}
	if ins.IsBool() {
		out.Boolean = ins.Bool()
	}

	if ins.IsString() {
		out.String = ins.String()
	}

	if ins.IsInt() {
		out.Int = ins.Int()
	}

	if ins.IsUint() {
		out.Uint = ins.Uint()
	}

	if ins.IsFloat() {
		out.Float = ins.Float()
	}

	if ins.IsList() {
		ptr, err := app.ConstantsToStruct(ins.List())
		if err != nil {
			return nil, err
		}

		out.List = ptr
	}

	return &out, nil
}

// StructToConstant converts a struct to constant
func (app *Adapter) StructToConstant(str Constant) (constants.Constant, error) {
	builder := app.constantBuilder.Create()
	if str.Boolean != nil {
		builder.WithBool(*str.Boolean)
	}

	if str.String != nil {
		builder.WithString(*str.String)
	}

	if str.Int != nil {
		builder.WithInt(*str.Int)
	}

	if str.Uint != nil {
		builder.WithUint(*str.Uint)
	}

	if str.Float != nil {
		builder.WithFloat(*str.Float)
	}

	if str.List != nil {
		ins, err := app.StructToConstants(str.List)
		if err != nil {
			return nil, err
		}

		builder.WithList(ins)
	}

	return builder.Now()
}
