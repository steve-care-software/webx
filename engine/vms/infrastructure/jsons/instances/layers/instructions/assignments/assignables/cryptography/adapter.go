package cryptography

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography"
	json_decrypts "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	json_encrypts "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	json_keys "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys"
)

// Adapter represents a cryptography adapter
type Adapter struct {
	encryptAdapter *json_encrypts.Adapter
	decryptAdapter *json_decrypts.Adapter
	keyAdapter     *json_keys.Adapter
	builder        cryptography.Builder
}

func createAdapter(
	encryptAdapter *json_encrypts.Adapter,
	decryptAdapter *json_decrypts.Adapter,
	keyAdapter *json_keys.Adapter,
	builder cryptography.Builder,
) cryptography.Adapter {
	out := Adapter{
		encryptAdapter: encryptAdapter,
		decryptAdapter: decryptAdapter,
		keyAdapter:     keyAdapter,
		builder:        builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins cryptography.Cryptography) ([]byte, error) {
	ptr, err := app.CryptographyToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (cryptography.Cryptography, error) {
	ins := new(Cryptography)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCryptography(*ins)
}

// CryptographyToStruct converts a cryptography to struct
func (app *Adapter) CryptographyToStruct(ins cryptography.Cryptography) (*Cryptography, error) {
	out := Cryptography{}
	if ins.IsEncrypt() {
		str := app.encryptAdapter.EncryptToStruct(ins.Encrypt())
		out.Encrypt = &str
	}

	if ins.IsDecrypt() {
		str := app.decryptAdapter.DecryptToStruct(ins.Decrypt())
		out.Decrypt = &str
	}

	if ins.IsKey() {
		ptr, err := app.keyAdapter.KeyToStruct(ins.Key())
		if err != nil {
			return nil, err
		}

		out.Key = ptr
	}

	return &out, nil
}

// StructToCryptography converts a struct to cryptography
func (app *Adapter) StructToCryptography(str Cryptography) (cryptography.Cryptography, error) {
	builder := app.builder.Create()
	if str.Encrypt != nil {
		ins, err := app.encryptAdapter.StructToEncrypt(*str.Encrypt)
		if err != nil {
			return nil, err
		}

		builder.WithEncrypt(ins)
	}

	if str.Decrypt != nil {
		ins, err := app.decryptAdapter.StructToDecrypt(*str.Decrypt)
		if err != nil {
			return nil, err
		}

		builder.WithDecrypt(ins)
	}

	if str.Key != nil {
		ins, err := app.keyAdapter.StructToKey(*str.Key)
		if err != nil {
			return nil, err
		}

		builder.WithKey(ins)
	}

	return builder.Now()
}
