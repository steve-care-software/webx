package cardinalities

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

// Adapter represents the cardinality adapter
type Adapter struct {
	builder cardinalities.Builder
}

func createAdapter(
	builder cardinalities.Builder,
) cardinalities.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins cardinalities.Cardinality) ([]byte, error) {
	str := app.CardinalityToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(data []byte) (cardinalities.Cardinality, error) {
	ins := new(Cardinality)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCardinality(*ins)
}

// CardinalityToStruct converts an instance to struct
func (app *Adapter) CardinalityToStruct(ins cardinalities.Cardinality) Cardinality {
	out := Cardinality{
		Min: ins.Min(),
	}

	if ins.HasMax() {
		pMax := ins.Max()
		out.Max = pMax
	}

	return out
}

// StructToCardinality converts a struct to cardinality
func (app *Adapter) StructToCardinality(str Cardinality) (cardinalities.Cardinality, error) {
	builder := app.builder.Create().WithMin(str.Min)
	if str.Max != nil {
		builder.WithMax(*str.Max)
	}

	return builder.Now()
}
