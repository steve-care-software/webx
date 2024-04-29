package accounts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts"
	json_communications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	json_credentials "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/credentials"
	json_encryptions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/encryptions"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
)

// Adapter represents an adapter
type Adapter struct {
	communicationAdapter *json_communications.Adapter
	credentialsAdapter   *json_credentials.Adapter
	encryptionAdapter    *json_encryptions.Adapter
	retrieveAdapter      *json_retrieves.Adapter
	builder              accounts.Builder
}

func createAdapter(
	communicationAdapter *json_communications.Adapter,
	credentialsAdapter *json_credentials.Adapter,
	encryptionAdapter *json_encryptions.Adapter,
	retrieveAdapter *json_retrieves.Adapter,
	builder accounts.Builder,
) accounts.Adapter {
	out := Adapter{
		communicationAdapter: communicationAdapter,
		credentialsAdapter:   credentialsAdapter,
		encryptionAdapter:    encryptionAdapter,
		retrieveAdapter:      retrieveAdapter,
		builder:              builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins accounts.Account) ([]byte, error) {
	ptr, err := app.AccountToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (accounts.Account, error) {
	ins := new(Account)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToAccount(*ins)
}

// AccountToStruct converts an account to struct
func (app *Adapter) AccountToStruct(ins accounts.Account) (*Account, error) {
	out := Account{}
	if ins.IsList() {
		out.List = ins.List()
	}

	if ins.IsCredentials() {
		str := app.credentialsAdapter.CredentialsToStruct(ins.Credentials())
		out.Credentials = &str
	}

	if ins.IsRetrieve() {
		str := app.retrieveAdapter.RetrieveToStruct(ins.Retrieve())
		out.Retrieve = &str
	}

	if ins.IsCommunication() {
		ptr, err := app.communicationAdapter.CommunicationToStruct(ins.Communication())
		if err != nil {
			return nil, err
		}

		out.Communication = ptr
	}

	if ins.IsEncryption() {
		ptr, err := app.encryptionAdapter.EncryptionToStruct(ins.Encryption())
		if err != nil {
			return nil, err
		}

		out.Encryption = ptr
	}

	return &out, nil
}

// StructToAccount converts a struct to account
func (app *Adapter) StructToAccount(str Account) (accounts.Account, error) {
	builder := app.builder.Create()
	if str.List != "" {
		builder.WithList(str.List)
	}

	if str.Credentials != nil {
		ins, err := app.credentialsAdapter.StructToCredentials(*str.Credentials)
		if err != nil {
			return nil, err
		}

		builder.WithCredentials(ins)
	}

	if str.Retrieve != nil {
		ins, err := app.retrieveAdapter.StructToRetrieve(*str.Retrieve)
		if err != nil {
			return nil, err
		}

		builder.WithRetrieve(ins)
	}

	if str.Communication != nil {
		ins, err := app.communicationAdapter.StructToCommunication(*str.Communication)
		if err != nil {
			return nil, err
		}

		builder.WithCommunication(ins)
	}

	if str.Encryption != nil {
		ins, err := app.encryptionAdapter.StructToEncryption(*str.Encryption)
		if err != nil {
			return nil, err
		}

		builder.WithEncryption(ins)
	}

	return builder.Now()
}
