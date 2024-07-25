package omissions

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	json_elements "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/elements"
)

// Adapter represents an adapter
type Adapter struct {
	elementAdapter *json_elements.Adapter
	builder        omissions.Builder
}

func createAdapter(
	elementAdapter *json_elements.Adapter,
	builder omissions.Builder,
) omissions.Adapter {
	out := Adapter{
		elementAdapter: elementAdapter,
		builder:        builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins omissions.Omission) ([]byte, error) {
	ptr, err := app.OmissionToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (omissions.Omission, error) {
	ins := new(Omission)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToOmission(*ins)
}

// OmissionToStruct converts an instance to struct
func (app *Adapter) OmissionToStruct(ins omissions.Omission) (*Omission, error) {
	out := Omission{}
	if ins.HasPrefix() {
		ptr, err := app.elementAdapter.ElementToStruct(ins.Prefix())
		if err != nil {
			return nil, err
		}

		out.Prefix = ptr
	}

	if ins.HasSuffix() {
		ptr, err := app.elementAdapter.ElementToStruct(ins.Suffix())
		if err != nil {
			return nil, err
		}

		out.Suffix = ptr
	}

	return &out, nil
}

// StructToOmission converts a struct to omission
func (app *Adapter) StructToOmission(str Omission) (omissions.Omission, error) {
	builder := app.builder.Create()
	if str.Prefix != nil {
		ins, err := app.elementAdapter.StructToElement(*str.Prefix)
		if err != nil {
			return nil, err
		}

		builder.WithPrefix(ins)
	}

	if str.Suffix != nil {
		ins, err := app.elementAdapter.StructToElement(*str.Suffix)
		if err != nil {
			return nil, err
		}

		builder.WithSuffix(ins)
	}

	return builder.Now()
}
