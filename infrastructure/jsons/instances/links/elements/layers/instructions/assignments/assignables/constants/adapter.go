package constants

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/constants"
)

// Adapter represents an adapter
type Adapter struct {
	builder constants.Builder
}

func createAdapter(
	builder constants.Builder,
) constants.Adapter {
	out := Adapter{
		builder: builder,
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

// ConstantToStruct converts a constant to struct
func (app *Adapter) ConstantToStruct(ins constants.Constant) (*Constant, error) {
	out := Constant{}
	if ins.IsBool() {
		out.Boolean = ins.Bool()
	}

	if ins.IsBytes() {
		value := ins.Bytes()
		out.Bytes = base64.StdEncoding.EncodeToString(value)
	}

	return &out, nil
}

// StructToConstant converts a struct to constant
func (app *Adapter) StructToConstant(str Constant) (constants.Constant, error) {
	builder := app.builder.Create()
	if str.Boolean != nil {
		builder.WithBool(*str.Boolean)
	}

	if str.Bytes != "" {
		decoded, err := base64.StdEncoding.DecodeString(str.Bytes)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(decoded)
	}

	return builder.Now()
}
