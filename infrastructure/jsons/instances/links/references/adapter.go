package references

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
)

// Adapter represents the adapter
type Adapter struct {
	hashAdapter      hash.Adapter
	referenceBuilder references.ReferenceBuilder
	builder          references.Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	referenceBuilder references.ReferenceBuilder,
	builder references.Builder,
) references.Adapter {
	out := Adapter{
		hashAdapter:      hashAdapter,
		referenceBuilder: referenceBuilder,
		builder:          builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins references.References) ([]byte, error) {
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

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (references.References, error) {
	ins := new([]Reference)
	err := json.Unmarshal(bytes, ins)
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
		Variable:   ins.Variable(),
		Identifier: ins.Identifier().String(),
	}, nil
}

// StructToReference converts a struct to reference
func (app *Adapter) StructToReference(str Reference) (references.Reference, error) {
	pHash, err := app.hashAdapter.FromString(str.Identifier)
	if err != nil {
		return nil, err
	}

	return app.referenceBuilder.Create().
		WithVariable(str.Variable).
		WithIdentifier(*pHash).
		Now()
}
