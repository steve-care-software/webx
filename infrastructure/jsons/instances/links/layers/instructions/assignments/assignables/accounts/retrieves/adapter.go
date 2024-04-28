package retrieves

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
)

// Adapter represents the adapter
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
	str := app.RetrieveToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (retrieves.Retrieve, error) {
	ins := new(Retrieve)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToRetrieve(*ins)
}

// RetrieveToStruct converts a retrieve to struct
func (app *Adapter) RetrieveToStruct(ins retrieves.Retrieve) Retrieve {
	return Retrieve{
		Password:    ins.Password(),
		Credentials: ins.Credentials(),
	}
}

// StructToRetrieve converts a struct to retrieve
func (app *Adapter) StructToRetrieve(str Retrieve) (retrieves.Retrieve, error) {
	return app.builder.Create().
		WithPassword(str.Password).
		WithCredentials(str.Credentials).
		Now()
}
