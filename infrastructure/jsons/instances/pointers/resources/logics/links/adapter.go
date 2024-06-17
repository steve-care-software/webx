package links

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
	json_elements "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/links/elements"
	json_references "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/links/references"
)

// Adapter represents the adapter
type Adapter struct {
	elementAdapter   *json_elements.Adapter
	referenceAdapter *json_references.Adapter
	builder          links.Builder
}

func createAdapter(
	elementAdapter *json_elements.Adapter,
	referenceAdapter *json_references.Adapter,
	builder links.Builder,
) links.Adapter {
	out := Adapter{
		elementAdapter:   elementAdapter,
		referenceAdapter: referenceAdapter,
		builder:          builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins links.Link) ([]byte, error) {
	ptr, err := app.LinkToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (links.Link, error) {
	ins := new(Link)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToLink(*ins)
}

// LinkToStruct converts a link to struct
func (app *Adapter) LinkToStruct(ins links.Link) (*Link, error) {
	strElements, err := app.elementAdapter.ElementsToStruct(ins.Elements())
	if err != nil {
		return nil, err
	}

	output := Link{
		Elements: strElements,
	}

	if ins.HasReferences() {
		strReferences, err := app.referenceAdapter.ReferencesToStruct(ins.References())
		if err != nil {
			return nil, err
		}

		output.References = strReferences
	}

	return &output, nil
}

// StructToLink converts a struct to link
func (app *Adapter) StructToLink(str Link) (links.Link, error) {
	elements, err := app.elementAdapter.StructToElements(str.Elements)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithElements(elements)
	if str.References != nil {
		references, err := app.referenceAdapter.StructToReferences(str.References)
		if err != nil {
			return nil, err
		}

		builder.WithReferences(references)
	}

	return builder.Now()
}
