package encryptions

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	json_decrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	json_encrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// Adapter represents the adapter
type Adapter struct {
	encryptAdapter *json_encrypts.Adapter
	decryptAdapter *json_decrypts.Adapter
	builder        encryptions.Builder
}

func createAdapter(
	encryptAdapter *json_encrypts.Adapter,
	decryptAdapter *json_decrypts.Adapter,
	builder encryptions.Builder,
) encryptions.Adapter {
	out := Adapter{
		encryptAdapter: encryptAdapter,
		decryptAdapter: decryptAdapter,
		builder:        builder,
	}

	return &out
}

// ToBytes converts instance to bytes
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

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(data []byte) (encryptions.Encryption, error) {
	ins := new(Encryption)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToEncryption(*ins)
}

// EncryptionToStruct converts an encryption to struct
func (app *Adapter) EncryptionToStruct(ins encryptions.Encryption) (*Encryption, error) {
	out := Encryption{}
	if ins.IsGeneratePrivateKey() {
		out.IsGeneratePrivateKey = true
	}

	if ins.IsFetchPublicKey() {
		out.FetchPublicKey = ins.FetchPublicKey()
	}

	if ins.IsEncrypt() {
		ptr, err := app.encryptAdapter.EncryptToStruct(ins.Encrypt())
		if err != nil {
			return nil, err
		}

		out.Encrypt = ptr
	}

	if ins.IsDecrypt() {
		ptr, err := app.decryptAdapter.DecryptToStruct(ins.Decrypt())
		if err != nil {
			return nil, err
		}

		out.Decrypt = ptr
	}

	return &out, nil
}

// StructToEncryption converts a struct to encryption
func (app *Adapter) StructToEncryption(str Encryption) (encryptions.Encryption, error) {
	builder := app.builder.Create()
	if str.IsGeneratePrivateKey {
		builder.IsGeneratePrivateKey()
	}

	if str.FetchPublicKey != "" {
		builder.WithFetchPublicKey(str.FetchPublicKey)
	}

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
