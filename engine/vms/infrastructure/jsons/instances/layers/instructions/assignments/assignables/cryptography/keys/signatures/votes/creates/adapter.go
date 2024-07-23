package creates

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
)

// Adapter represents the adapter
type Adapter struct {
	builder creates.Builder
}

func createAdapter(
	builder creates.Builder,
) creates.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins creates.Create) ([]byte, error) {
	ptr, err := app.CreateToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (creates.Create, error) {
	ins := new(Create)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCreate(*ins)
}

// CreateToStruct converts a create to struct
func (app *Adapter) CreateToStruct(ins creates.Create) (*Create, error) {
	return &Create{
		Message:    ins.Message(),
		Ring:       ins.Ring(),
		PrivateKey: ins.PrivateKey(),
	}, nil
}

// StructToCreate converts a struct to create
func (app *Adapter) StructToCreate(str Create) (creates.Create, error) {
	return app.builder.Create().
		WithMessage(str.Message).
		WithRing(str.Ring).
		WithPrivateKey(str.PrivateKey).
		Now()
}
