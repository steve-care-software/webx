package resources

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
)

// Adapter represents an adapter
type Adapter struct {
	builder         resources.Builder
	instanceAdapter instances.Adapter
}

func createAdapter(
	builder resources.Builder,
	instanceAdapter instances.Adapter,
) resources.Adapter {
	out := Adapter{
		builder:         builder,
		instanceAdapter: instanceAdapter,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins resources.Resource) ([]byte, error) {
	str, err := app.ResourceToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(str)
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
	bytes, err := app.instanceAdapter.ToBytes(ins.Path(), ins.Instance())
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	return &Resource{
		Path:     ins.Path(),
		Instance: encoded,
	}, nil
}

// StructToResource converts a struct to resource
func (app *Adapter) StructToResource(str Resource) (resources.Resource, error) {
	decoded, err := base64.StdEncoding.DecodeString(str.Instance)
	if err != nil {
		return nil, err
	}

	ins, err := app.instanceAdapter.ToInstance(str.Path, decoded)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithPath(str.Path).
		WithInstance(ins).
		Now()
}
