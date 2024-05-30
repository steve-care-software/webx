package resources

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/origins/resources"
)

// Adapter represents an adapter
type Adapter struct {
	hashAdapter hash.Adapter
	builder     resources.Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder resources.Builder,
) resources.Adapter {
	out := Adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins resources.Resource) ([]byte, error) {
	ptr, err := app.ResourceToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (resources.Resource, error) {
	ins := new(Resource)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToResource(*ins)
}

// ResourceToStruct converts a resource to struct
func (app *Adapter) ResourceToStruct(ins resources.Resource) (*Resource, error) {
	output := Resource{
		Layer: ins.Layer().String(),
	}

	if ins.IsMandatory() {
		output.IsMandatory = true
	}

	return &output, nil
}

// StructToResource converts a struct to resource
func (app *Adapter) StructToResource(str Resource) (resources.Resource, error) {
	pHash, err := app.hashAdapter.FromString(str.Layer)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithLayer(*pHash)
	if str.IsMandatory {
		builder.IsMandatory()
	}

	return builder.Now()
}
