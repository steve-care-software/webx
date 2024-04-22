package pointers

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
)

// Adapter represents an adapter
type Adapter struct {
	hashAdapter hash.Adapter
	builder     pointers.Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder pointers.Builder,
) pointers.Adapter {
	out := Adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToBytes converts an intance to bytes
func (app *Adapter) ToBytes(ins pointers.Pointer) ([]byte, error) {
	str := app.PointerToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (pointers.Pointer, error) {
	ins := new(Pointer)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToPointer(*ins)
}

// PointerToStruct converts a pointer to struct
func (app *Adapter) PointerToStruct(ins pointers.Pointer) Pointer {
	return Pointer{
		Path:       ins.Path(),
		Identifier: ins.Identifier().String(),
	}
}

// StructToPointer converts a struct to pointer
func (app *Adapter) StructToPointer(str Pointer) (pointers.Pointer, error) {
	pHash, err := app.hashAdapter.FromString(str.Identifier)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithIdentifier(*pHash).
		WithPath(str.Path).
		Now()
}
