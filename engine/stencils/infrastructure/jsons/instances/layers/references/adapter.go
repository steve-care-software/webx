package references

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/references"
)

// Adapter represents the adapter
type Adapter struct {
	referenceBuilder references.ReferenceBuilder
	builder          references.Builder
}

func createAdapter(
	referenceBuilder references.ReferenceBuilder,
	builder references.Builder,
) references.Adapter {
	out := Adapter{
		referenceBuilder: referenceBuilder,
		builder:          builder,
	}

	return &out
}

// InstanceToBytes converts an instance to bytes
func (app *Adapter) InstanceToBytes(ins references.Reference) ([]byte, error) {
	ptr, err := app.ReferenceToStruct(ins)
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
func (app *Adapter) BytesToInstance(data []byte) (references.Reference, error) {
	ins := new(Reference)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToReference(*ins)
}

// InstancesToBytes converts an instances to bytes
func (app *Adapter) InstancesToBytes(ins references.References) ([]byte, error) {
	ptr, err := app.ReferencesToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// BytesToInstances converts bytes to instances
func (app *Adapter) BytesToInstances(data []byte) (references.References, error) {
	ins := new([]Reference)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToReferences(*ins)
}

// ReferencesToStruct converts references to struct
func (app *Adapter) ReferencesToStruct(ins references.References) ([]Reference, error) {
	output := []Reference{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.ReferenceToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToReferences converts a struct to references
func (app *Adapter) StructToReferences(list []Reference) (references.References, error) {
	output := []references.Reference{}
	for _, oneStr := range list {
		ins, err := app.StructToReference(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// ReferenceToStruct converts a reference to struct
func (app *Adapter) ReferenceToStruct(ins references.Reference) (*Reference, error) {
	return &Reference{
		Variable: ins.Variable(),
		Path:     ins.Path(),
	}, nil
}

// StructToReference converts a struct to reference
func (app *Adapter) StructToReference(str Reference) (references.Reference, error) {
	return app.referenceBuilder.Create().
		WithVariable(str.Variable).
		WithPath(str.Path).
		Now()
}
