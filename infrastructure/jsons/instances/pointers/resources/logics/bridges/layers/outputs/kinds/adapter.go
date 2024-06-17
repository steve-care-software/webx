package kinds

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs/kinds"
)

// Adapter represents the adapter
type Adapter struct {
	builder kinds.Builder
}

func createAdapter(
	builder kinds.Builder,
) kinds.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins kinds.Kind) ([]byte, error) {
	ptr, err := app.KindToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (kinds.Kind, error) {
	ins := new(Kind)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToKind(*ins)
}

// KindToStruct converts a kind to struct
func (app *Adapter) KindToStruct(ins kinds.Kind) (*Kind, error) {
	out := Kind{}
	if ins.IsContinue() {
		out.Continue = ins.IsContinue()
	}

	if ins.IsPrompt() {
		out.Prompt = ins.IsPrompt()
	}

	return &out, nil
}

// StructToKind converts a struct to kind
func (app *Adapter) StructToKind(str Kind) (kinds.Kind, error) {
	builder := app.builder.Create()
	if str.Continue {
		builder.IsContinue()
	}

	if str.Prompt {
		builder.IsPrompt()
	}

	return builder.Now()
}
