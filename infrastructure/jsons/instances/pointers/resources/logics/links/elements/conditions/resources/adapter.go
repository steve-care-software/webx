package resources

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions/resources"
)

// Adapter represents the adapter
type Adapter struct {
	builder resources.Builder
}

func createAdapter(
	builder resources.Builder,
) resources.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes represents the instance to bytes
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

// ToInstance represents the bytes to instance
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
	out := Resource{
		Code: ins.Code(),
	}

	if ins.IsRaisedInLayer() {
		out.IsRaisedInLayer = true
	}

	return &out, nil
}

// StructToResource converts a struct to resource
func (app *Adapter) StructToResource(str Resource) (resources.Resource, error) {
	builder := app.builder.Create().WithCode(str.Code)
	if str.IsRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	return builder.Now()
}
