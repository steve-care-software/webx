package bytes

import (
	go_bytes "bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
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
func (app *Adapter) ToInstance(data []byte) (bytes.Bytes, error) {
	ins := new(Bytes)
	decoder := json.NewDecoder(go_bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
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
