package decrypts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
)

// Adapter represents the decrypt adapter
type Adapter struct {
	builder decrypts.Builder
}

func createAdapter(
	builder decrypts.Builder,
) decrypts.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts decrypt to bytes
func (app *Adapter) ToBytes(ins decrypts.Decrypt) ([]byte, error) {
	str := app.DecryptToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to decrypt
func (app *Adapter) ToInstance(bytes []byte) (decrypts.Decrypt, error) {
	ins := new(Decrypt)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToDecrypt(*ins)
}

// DecryptToStruct converts a decrypt to struct
func (app *Adapter) DecryptToStruct(ins decrypts.Decrypt) Decrypt {
	return Decrypt{
		Cipher:   ins.Cipher(),
		Password: ins.Password(),
	}
}

// StructToDecrypt converts a struct to decrypt
func (app *Adapter) StructToDecrypt(str Decrypt) (decrypts.Decrypt, error) {
	return app.builder.Create().
		WithCipher(str.Cipher).
		WithPassword(str.Password).
		Now()
}
