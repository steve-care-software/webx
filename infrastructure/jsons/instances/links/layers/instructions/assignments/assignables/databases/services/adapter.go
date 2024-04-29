package services

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases/services"
)

// Adapter represents an adapter
type Adapter struct {
	builder services.Builder
}

func createAdapter(
	builder services.Builder,
) services.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins services.Service) ([]byte, error) {
	ptr, err := app.ServiceToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (services.Service, error) {
	ins := new(Service)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToService(*ins)
}

// ServiceToStruct converts a service to struct
func (app *Adapter) ServiceToStruct(ins services.Service) (*Service, error) {
	out := Service{}
	if ins.IsBegin() {
		out.IsBegin = true
	}

	return &out, nil
}

// StructToService converts a struct to service
func (app *Adapter) StructToService(str Service) (services.Service, error) {
	builder := app.builder.Create()
	if str.IsBegin {
		builder.IsBegin()
	}

	return builder.Now()
}
