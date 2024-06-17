package encrypts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/encrypts"
)

// Adapter represents the encrypt adapter
type Adapter struct {
	builder encrypts.Builder
}

func createAdapter(
	builder encrypts.Builder,
) encrypts.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts encrypt to bytes
func (app *Adapter) ToBytes(ins encrypts.Encrypt) ([]byte, error) {
	str := app.EncryptToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to encrypt
func (app *Adapter) ToInstance(bytes []byte) (encrypts.Encrypt, error) {
	ins := new(Encrypt)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToEncrypt(*ins)
}

// EncryptToStruct converts a encrypt to struct
func (app *Adapter) EncryptToStruct(ins encrypts.Encrypt) Encrypt {
	return Encrypt{
		Message:  ins.Message(),
		Password: ins.Password(),
	}
}

// StructToEncrypt converts a struct to encrypt
func (app *Adapter) StructToEncrypt(str Encrypt) (encrypts.Encrypt, error) {
	return app.builder.Create().
		WithMessage(str.Message).
		WithPassword(str.Password).
		Now()
}
