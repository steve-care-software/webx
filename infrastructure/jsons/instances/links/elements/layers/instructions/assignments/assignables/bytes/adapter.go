package bytes

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/bytes"
)

// Adapter represents an adapter
type Adapter struct {
	builder bytes.Builder
}

func createAdapter(
	builder bytes.Builder,
) bytes.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins bytes.Bytes) ([]byte, error) {
	ptr, err := app.BytesToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (bytes.Bytes, error) {
	ins := new(Bytes)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToBytes(*ins)
}

// BytesToStruct converts a bytes to struct
func (app *Adapter) BytesToStruct(ins bytes.Bytes) (*Bytes, error) {
	out := Bytes{}
	if ins.IsJoin() {
		out.Join = ins.Join()
	}

	if ins.IsCompare() {
		out.Compare = ins.Compare()
	}

	if ins.IsHashBytes() {
		out.Hash = ins.HashBytes()
	}

	return &out, nil
}

// StructToBytes converts a struct to bytes
func (app *Adapter) StructToBytes(str Bytes) (bytes.Bytes, error) {
	builder := app.builder.Create()
	if str.Join != nil {
		builder.WithJoin(str.Join)
	}

	if str.Compare != nil {
		builder.WithCompare(str.Compare)
	}

	if str.Hash != "" {
		builder.WithHashBytes(str.Hash)
	}

	return builder.Now()
}
