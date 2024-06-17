package bridges

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers"
)

// Adapter represents an adapter
type Adapter struct {
	layerAdapter  *json_layers.Adapter
	builder       bridges.Builder
	bridgeBuilder bridges.BridgeBuilder
}

func createAdapter(
	layerAdapter *json_layers.Adapter,
	builder bridges.Builder,
	bridgeBuilder bridges.BridgeBuilder,
) bridges.Adapter {
	out := Adapter{
		layerAdapter:  layerAdapter,
		builder:       builder,
		bridgeBuilder: bridgeBuilder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins bridges.Bridges) ([]byte, error) {
	ptr, err := app.BridgesToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (bridges.Bridges, error) {
	ins := new([]Bridge)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToBridges(*ins)
}

// BridgesToStruct converts a bridges to struct
func (app *Adapter) BridgesToStruct(ins bridges.Bridges) ([]Bridge, error) {
	list := ins.List()
	output := []Bridge{}
	for _, oneBridge := range list {
		ptr, err := app.BridgeToStruct(oneBridge)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToBridges converts a struct to bridges
func (app *Adapter) StructToBridges(str []Bridge) (bridges.Bridges, error) {
	list := []bridges.Bridge{}
	for _, oneStr := range str {
		ins, err := app.StructToBridge(oneStr)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}

// BridgeToStruct converts a bridge to struct
func (app *Adapter) BridgeToStruct(ins bridges.Bridge) (*Bridge, error) {
	ptr, err := app.layerAdapter.LayerToStruct(ins.Layer())
	if err != nil {
		return nil, err
	}

	return &Bridge{
		Path:  ins.Path(),
		Layer: *ptr,
	}, nil
}

// StructToBridge converts a struct to bridge
func (app *Adapter) StructToBridge(str Bridge) (bridges.Bridge, error) {
	layer, err := app.layerAdapter.StructToLayer(str.Layer)
	if err != nil {
		return nil, err
	}

	return app.bridgeBuilder.Create().
		WithPath(str.Path).
		WithLayer(layer).
		Now()
}
