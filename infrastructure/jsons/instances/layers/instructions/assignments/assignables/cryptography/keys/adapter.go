package keys

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	json_encryption "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	json_signatures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

// Adapter represents an adapter
type Adapter struct {
	encryptionAdapter *json_encryption.Adapter
	signatureAdapter  *json_signatures.Adapter
	builder           keys.Builder
}

func createAdapter(
	encryptionAdapter *json_encryption.Adapter,
	signatureAdapter *json_signatures.Adapter,
	builder keys.Builder,
) keys.Adapter {
	out := Adapter{
		encryptionAdapter: encryptionAdapter,
		signatureAdapter:  signatureAdapter,
		builder:           builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins keys.Key) ([]byte, error) {
	ptr, err := app.KeyToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (keys.Key, error) {
	ins := new(Key)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToKey(*ins)
}

// KeyToStruct converts a key to struct
func (app *Adapter) KeyToStruct(ins keys.Key) (*Key, error) {
	out := Key{}
	if ins.IsEncryption() {
		ptr, err := app.encryptionAdapter.EncryptionToStruct(ins.Encryption())
		if err != nil {
			return nil, err
		}

		out.Encryption = ptr
	}

	if ins.IsSignature() {
		ptr, err := app.signatureAdapter.SignatureToStruct(ins.Signature())
		if err != nil {
			return nil, err
		}

		out.Signature = ptr
	}

	return &out, nil
}

// StructToKey converts a struct to key
func (app *Adapter) StructToKey(str Key) (keys.Key, error) {
	builder := app.builder.Create()
	if str.Encryption != nil {
		ins, err := app.encryptionAdapter.StructToEncryption(*str.Encryption)
		if err != nil {
			return nil, err
		}

		builder.WithEncryption(ins)
	}

	if str.Signature != nil {
		ins, err := app.signatureAdapter.StructToSignature(*str.Signature)
		if err != nil {
			return nil, err
		}

		builder.WithSignature(ins)
	}

	return builder.Now()
}
