package accounts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts"
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/accounts/inserts"
	json_updates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/accounts/updates"
)

// Adapter represents an adapter
type Adapter struct {
	insertAdapter *json_inserts.Adapter
	updateAdapter *json_updates.Adapter
	builder       accounts.Builder
}

func createAdapter(
	insertAdapter *json_inserts.Adapter,
	updateAdapter *json_updates.Adapter,
	builder accounts.Builder,
) accounts.Adapter {
	out := Adapter{
		insertAdapter: insertAdapter,
		updateAdapter: updateAdapter,
		builder:       builder,
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
	output := Account{}
	if ins.IsInsert() {
		str := app.insertAdapter.InsertToStruct(ins.Insert())
		output.Insert = &str
	}

	if ins.IsUpdate() {
		ptr, err := app.updateAdapter.UpdateToStruct(ins.Update())
		if err != nil {
			return nil, err
		}

		output.Update = ptr
	}

	if ins.IsDelete() {
		output.Delete = ins.Delete()
	}

	return &output, nil
}

// StructToAccount converts a struct to account
func (app *Adapter) StructToAccount(str Account) (accounts.Account, error) {
	builder := app.builder.Create()
	if str.Insert != nil {
		insert, err := app.insertAdapter.StructToInsert(*str.Insert)
		if err != nil {
			return nil, err
		}

		builder.WithInsert(insert)
	}

	if str.Update != nil {
		update, err := app.updateAdapter.StructToUpdate(*str.Update)
		if err != nil {
			return nil, err
		}

		builder.WithUpdate(update)
	}

	if str.Delete != "" {
		builder.WithDelete(str.Delete)
	}

	return builder.Now()
}
