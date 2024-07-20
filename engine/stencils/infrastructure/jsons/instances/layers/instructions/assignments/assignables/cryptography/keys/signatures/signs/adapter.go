package signs

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	json_creates "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	json_validates "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// Adapter represents the adapter
type Adapter struct {
	createAdapter   *json_creates.Adapter
	validateAdapter *json_validates.Adapter
	builder         signs.Builder
}

func createAdapter(
	createAdapter *json_creates.Adapter,
	validateAdapter *json_validates.Adapter,
	builder signs.Builder,
) signs.Adapter {
	out := Adapter{
		createAdapter:   createAdapter,
		validateAdapter: validateAdapter,
		builder:         builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins signs.Sign) ([]byte, error) {
	ptr, err := app.SignToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (signs.Sign, error) {
	ins := new(Sign)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToSign(*ins)
}

// SignToStruct converts a sign to struct
func (app *Adapter) SignToStruct(ins signs.Sign) (*Sign, error) {
	out := Sign{}
	if ins.IsCreate() {
		ptr, err := app.createAdapter.CreateToStruct(ins.Create())
		if err != nil {
			return nil, err
		}

		out.Create = ptr
	}

	if ins.IsValidate() {
		ptr, err := app.validateAdapter.ValidateToStruct(ins.Validate())
		if err != nil {
			return nil, err
		}

		out.Validate = ptr
	}

	return &out, nil
}

// StructToSign converts a struct to sign
func (app *Adapter) StructToSign(str Sign) (signs.Sign, error) {
	builder := app.builder.Create()
	if str.Create != nil {
		ins, err := app.createAdapter.StructToCreate(*str.Create)
		if err != nil {
			return nil, err
		}

		builder.WithCreate(ins)
	}

	if str.Validate != nil {
		ins, err := app.validateAdapter.StructToValidate(*str.Validate)
		if err != nil {
			return nil, err
		}

		builder.WithValidate(ins)
	}

	return builder.Now()
}
