package links

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
	json_elements "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements"
	json_origins "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins"
	json_references "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/references"
)

// Adapter represents the adapter
type Adapter struct {
	elementAdapter   *json_elements.Adapter
	originAdapter    *json_origins.Adapter
	referenceAdapter *json_references.Adapter
	linkBuilder      links.LinkBuilder
	builder          links.Builder
}

func createAdapter(
	elementAdapter *json_elements.Adapter,
	originAdapter *json_origins.Adapter,
	referenceAdapter *json_references.Adapter,
	linkBuilder links.LinkBuilder,
	builder links.Builder,
) links.Adapter {
	out := Adapter{
		elementAdapter:   elementAdapter,
		originAdapter:    originAdapter,
		referenceAdapter: referenceAdapter,
		linkBuilder:      linkBuilder,
		builder:          builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins links.Links) ([]byte, error) {
	ptr, err := app.LinksToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (links.Links, error) {
	ins := new([]Link)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToLinks(*ins)
}

// LinksToStruct converts a links to struct
func (app *Adapter) LinksToStruct(ins links.Links) ([]Link, error) {
	output := []Link{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.LinkToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToLinks converts a struct to links
func (app *Adapter) StructToLinks(list []Link) (links.Links, error) {
	output := []links.Link{}
	for _, oneStr := range list {
		ins, err := app.StructToLink(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// LinkToStruct converts a link to struct
func (app *Adapter) LinkToStruct(ins links.Link) (*Link, error) {
	ptrOrigin, err := app.originAdapter.OriginToStruct(ins.Origin())
	if err != nil {
		return nil, err
	}

	strElements, err := app.elementAdapter.ElementsToStruct(ins.Elements())
	if err != nil {
		return nil, err
	}

	output := Link{
		Origin:   *ptrOrigin,
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
	origin, err := app.originAdapter.StructToOrigin(str.Origin)
	if err != nil {
		return nil, err
	}

	elements, err := app.elementAdapter.StructToElements(str.Elements)
	if err != nil {
		return nil, err
	}

	builder := app.linkBuilder.Create().WithOrigin(origin).WithElements(elements)
	if str.References != nil {
		references, err := app.referenceAdapter.StructToReferences(str.References)
		if err != nil {
			return nil, err
		}

		builder.WithReferences(references)
	}

	return builder.Now()
}
