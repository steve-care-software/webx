package links

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers"
	json_links "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links"
)

// Adapter represents an adapter
type Adapter struct {
	layerAdapter *json_layers.Adapter
	linkAdapter  *json_links.Adapter
	builder      links.Builder
}

func createAdapter(
	layerAdapter *json_layers.Adapter,
	linkAdapter *json_links.Adapter,
	builder links.Builder,
) links.Adapter {
	out := Adapter{
		layerAdapter: layerAdapter,
		linkAdapter:  linkAdapter,
		builder:      builder,
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
	ptrSource, err := app.linkAdapter.LinkToStruct(ins.Source())
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(ins.Input())
	out := Link{
		Input:  encoded,
		Source: *ptrSource,
	}

	if ins.HasLayers() {
		layers, err := app.layerAdapter.LayersToStruct(ins.Layers())
		if err != nil {
			return nil, err
		}

		out.Layers = layers
	}

	return &out, nil
}

// StructToLink converts a struct to link
func (app *Adapter) StructToLink(str Link) (links.Link, error) {
	sourceIns, err := app.linkAdapter.StructToLink(str.Source)
	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(str.Input)
	if err != nil {
		return nil, err
	}

	builder := app.builder.WithInput(decoded).WithSource(sourceIns)
	if str.Layers != nil && len(str.Layers) > 0 {
		layersIns, err := app.layerAdapter.StructToLayers(str.Layers)
		if err != nil {
			return nil, err
		}

		builder.WithLayers(layersIns)
	}

	return builder.Now()
}
