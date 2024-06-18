package heads

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
)

// Adapter represents an head adapter
type Adapter struct {
	builder heads.Builder
}

func createAdapter(
	builder heads.Builder,
) heads.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins heads.Head) ([]byte, error) {
	ptr, err := app.HeadToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (heads.Head, error) {
	ins := new(Head)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToHead(*ins)
}

// HeadToStruct converts an head to struct
func (app *Adapter) HeadToStruct(ins heads.Head) (*Head, error) {
	out := Head{
		Path:        ins.Path(),
		Description: ins.Description(),
		IsActive:    false,
	}

	if ins.IsActive() {
		out.IsActive = true
	}

	return &out, nil
}

// StructToHead converts a struct to head
func (app *Adapter) StructToHead(str Head) (heads.Head, error) {
	builder := app.builder.Create().
		WithPath(str.Path).
		WithDescription(str.Description)

	if str.IsActive {
		builder.IsActive()
	}

	return builder.Now()
}
