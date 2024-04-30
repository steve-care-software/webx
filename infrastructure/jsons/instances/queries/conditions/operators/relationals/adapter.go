package relationals

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/relationals"
)

// Adapter represents an adapter
type Adapter struct {
	builder relationals.Builder
}

func createAdapter(
	builder relationals.Builder,
) relationals.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins relationals.Relational) ([]byte, error) {
	ptr, err := app.RelationalToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to relational
func (app *Adapter) ToInstance(bytes []byte) (relationals.Relational, error) {
	ins := new(Relational)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToRelational(*ins)
}

// RelationalToStruct converts a relational to struct
func (app *Adapter) RelationalToStruct(ins relationals.Relational) (*Relational, error) {
	output := Relational{}
	if ins.IsAnd() {
		output.IsAnd = true
	}

	if ins.IsOr() {
		output.IsOr = true
	}

	return &output, nil
}

// StructToRelational converts a struct to relational
func (app *Adapter) StructToRelational(str Relational) (relationals.Relational, error) {
	builder := app.builder.Create()
	if str.IsAnd {
		builder.IsAnd()
	}

	if str.IsOr {
		builder.IsOr()
	}

	return builder.Now()
}
