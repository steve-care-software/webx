package encrypts

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// Adapter represents an adapter
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

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins encrypts.Encrypt) ([]byte, error) {
	ptr, err := app.EncryptToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (encrypts.Encrypt, error) {
	ins := new(Encrypt)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToEncrypt(*ins)
}

// EncryptToStruct converts an encrypt to struct
func (app *Adapter) EncryptToStruct(ins encrypts.Encrypt) (*Encrypt, error) {
	return &Encrypt{
		Message:   ins.Message(),
		PublicKey: ins.PublicKey(),
	}, nil
}

// StructToEncrypt converts a struct to encrypt
func (app *Adapter) StructToEncrypt(str Encrypt) (encrypts.Encrypt, error) {
	return app.builder.Create().
		WithMessage(str.Message).
		WithPublicKey(str.PublicKey).
		Now()
}
