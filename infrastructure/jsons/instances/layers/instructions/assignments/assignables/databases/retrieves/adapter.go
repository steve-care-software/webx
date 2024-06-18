package retrieves

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/retrieves"
)

// Adapter represents a retrieve adapter
type Adapter struct {
	builder retrieves.Builder
}

func createAdapter(
	builder retrieves.Builder,
) retrieves.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins retrieves.Retrieve) ([]byte, error) {
	ptr, err := app.RetrieveToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instanxce
func (app *Adapter) ToInstance(bytes []byte) (retrieves.Retrieve, error) {
	ins := new(Retrieve)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToRetrieve(*ins)
}

// RetrieveToStruct converts a retrieve to struct
func (app *Adapter) RetrieveToStruct(ins retrieves.Retrieve) (*Retrieve, error) {
	out := Retrieve{}
	if ins.IsList() {
		out.IsList = true
	}

	if ins.IsExists() {
		out.Exists = ins.Exists()
	}

	if ins.IsRetrieve() {
		out.Retrieve = ins.Retrieve()
	}

	return &out, nil
}

// StructToRetrieve converts a struct to retrieve
func (app *Adapter) StructToRetrieve(str Retrieve) (retrieves.Retrieve, error) {
	builder := app.builder.Create()
	if str.IsList {
		builder.IsList()
	}

	if str.Exists != "" {
		builder.WithExists(str.Exists)
	}

	if str.Retrieve != "" {
		builder.WithRetrieve(str.Retrieve)
	}

	return builder.Now()
}
