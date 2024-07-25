package elements

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
)

// Adapter represents an adapter
type Adapter struct {
	hashAdapter    hash.Adapter
	builder        elements.Builder
	elementBuilder elements.ElementBuilder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder elements.Builder,
	elementBuilder elements.ElementBuilder,
) elements.Adapter {
	out := Adapter{
		hashAdapter:    hashAdapter,
		builder:        builder,
		elementBuilder: elementBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *Adapter) InstancesToBytes(ins elements.Elements) ([]byte, error) {
	ptr, err := app.ElementsToStructs(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// BytesToInstances converts bytes to instance
func (app *Adapter) BytesToInstances(data []byte) (elements.Elements, error) {
	ins := new([]Element)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructsToElements(*ins)
}

// InstanceToBytes converts instance to bytes
func (app *Adapter) InstanceToBytes(ins elements.Element) ([]byte, error) {
	ptr, err := app.ElementToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// BytesToInstance converts bytes to instance
func (app *Adapter) BytesToInstance(data []byte) (elements.Element, error) {
	ins := new(Element)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToElement(*ins)
}

// ElementsToStruct converts an instance to structs
func (app *Adapter) ElementsToStructs(ins elements.Elements) ([]Element, error) {
	list := ins.List()
	out := []Element{}
	for _, oneIns := range list {
		ptr, err := app.ElementToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructsToElements converts a struct to elements
func (app *Adapter) StructsToElements(str []Element) (elements.Elements, error) {
	list := []elements.Element{}
	builder := app.builder.Create()
	for _, oneStr := range str {
		ins, err := app.StructToElement(oneStr)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return builder.WithList(list).
		Now()
}

// ElementToStruct converts an instance to struct
func (app *Adapter) ElementToStruct(ins elements.Element) (*Element, error) {
	out := Element{}
	if ins.IsBytes() {
		out.Bytes = ins.Bytes()
	}

	if ins.IsLayer() {
		out.Layer = ins.Layer().String()
	}

	if ins.IsString() {
		out.String = ins.String()
	}

	return &out, nil
}

// StructToElement converts a struct to element
func (app *Adapter) StructToElement(str Element) (elements.Element, error) {
	builder := app.elementBuilder.Create()
	if str.Layer != "" {
		pHash, err := app.hashAdapter.FromString(str.Layer)
		if err != nil {
			return nil, err
		}

		builder.WithLayer(*pHash)
	}

	if str.String != "" {
		builder.WithString(str.String)
	}

	if str.Bytes != nil && len(str.Bytes) > 0 {
		builder.WithBytes(str.Bytes)
	}

	return builder.Now()
}
