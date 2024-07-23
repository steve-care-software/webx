package validates

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// Adapter represents the adapter
type Adapter struct {
	builder validates.Builder
}

func createAdapter(
	builder validates.Builder,
) validates.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts validates to bytes
func (app *Adapter) ToBytes(ins validates.Validate) ([]byte, error) {
	ptr, err := app.ValidateToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to validates
func (app *Adapter) ToInstance(data []byte) (validates.Validate, error) {
	ins := new(Validate)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToValidate(*ins)
}

// ValidateToStruct converts a validate to struct
func (app *Adapter) ValidateToStruct(ins validates.Validate) (*Validate, error) {
	return &Validate{
		Signature: ins.Signature(),
		Message:   ins.Message(),
		PublicKey: ins.PublicKey(),
	}, nil
}

// StructToValidate converts a struct to validate
func (app *Adapter) StructToValidate(str Validate) (validates.Validate, error) {
	return app.builder.Create().
		WithSignature(str.Signature).
		WithMessage(str.Message).
		WithPublicKey(str.PublicKey).
		Now()
}
