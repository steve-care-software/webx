package signs

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
)

// Adapter represents an adapter
type Adapter struct {
	builder signs.Builder
}

func createAdapter(
	builder signs.Builder,
) signs.Adapter {
	out := Adapter{
		builder: builder,
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
func (app *Adapter) ToInstance(bytes []byte) (signs.Sign, error) {
	ins := new(Sign)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToSign(*ins)
}

// SignToStruct converts a sign to struct
func (app *Adapter) SignToStruct(ins signs.Sign) (*Sign, error) {
	return &Sign{
		Message: ins.Message(),
		Account: ins.Account(),
	}, nil
}

// StructToSign converts a struct to sign
func (app *Adapter) StructToSign(str Sign) (signs.Sign, error) {
	return app.builder.Create().
		WithMessage(str.Message).
		WithAccount(str.Account).
		Now()
}
