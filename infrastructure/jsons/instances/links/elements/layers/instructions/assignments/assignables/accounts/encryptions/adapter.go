package encryptions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions"
	json_decrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	json_encrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

// Adapter represents an adapter
type Adapter struct {
	decryptAdapter *json_decrypts.Adapter
	encryptAdapter *json_encrypts.Adapter
	builder        encryptions.Builder
}

func createAdapter(
	decryptAdapter *json_decrypts.Adapter,
	encryptAdapter *json_encrypts.Adapter,
	builder encryptions.Builder,
) encryptions.Adapter {
	out := Adapter{
		decryptAdapter: decryptAdapter,
		encryptAdapter: encryptAdapter,
		builder:        builder,
	}

	return &out
}

// ToBytes converts encryption to bytes
func (app *Adapter) ToBytes(ins encryptions.Encryption) ([]byte, error) {
	ptr, err := app.EncryptionToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to encryption
func (app *Adapter) ToInstance(bytes []byte) (encryptions.Encryption, error) {
	ins := new(Encryption)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToEncryption(*ins)
}

// EncryptionToStruct converts an encryption to struct
func (app *Adapter) EncryptionToStruct(ins encryptions.Encryption) (*Encryption, error) {
	out := Encryption{}
	if ins.IsDecrypt() {
		str := app.decryptAdapter.DecryptToStruct(ins.Decrypt())
		out.Decrypt = &str
	}

	if ins.IsEncrypt() {
		str := app.encryptAdapter.EncryptToStruct(ins.Encrypt())
		out.Encrypt = &str
	}

	return &out, nil
}

// StructToEncryption converts a struct to encryption
func (app *Adapter) StructToEncryption(str Encryption) (encryptions.Encryption, error) {
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

	return builder.Now()
}
