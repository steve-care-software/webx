package credentials

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/credentials"
)

type Adapter struct {
	builder credentials.Builder
}

func createAdapter(
	builder credentials.Builder,
) credentials.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins credentials.Credentials) ([]byte, error) {
	str := app.CredentialsToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (credentials.Credentials, error) {
	ins := new(Credentials)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCredentials(*ins)
}

// CredentialsToStruct converts a credentials to struct
func (app *Adapter) CredentialsToStruct(ins credentials.Credentials) Credentials {
	return Credentials{
		Username: ins.Username(),
		Password: ins.Password(),
	}
}

// StructToCredentials converts a struct to credentils
func (app *Adapter) StructToCredentials(str Credentials) (credentials.Credentials, error) {
	return app.builder.Create().
		WithUsername(str.Username).
		WithPassword(str.Password).
		Now()
}
