package creates

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
)

// Adapter represents an adapter
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

// ToBytes converts instance to bytes
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
func (app *Adapter) ToInstance(bytes []byte) (creates.Create, error) {
	ins := new(Create)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCreate(*ins)
}

// CreateToStruct converts a create to struct
func (app *Adapter) CreateToStruct(ins creates.Create) (*Create, error) {
	return &Create{
		Message:    ins.Message(),
		PrivateKey: ins.PrivateKey(),
	}, nil
}

// StructToCreate converts a struct to create
func (app *Adapter) StructToCreate(str Create) (creates.Create, error) {
	return app.builder.Create().
		WithMessage(str.Message).
		WithPrivateKey(str.PrivateKey).
		Now()
}
