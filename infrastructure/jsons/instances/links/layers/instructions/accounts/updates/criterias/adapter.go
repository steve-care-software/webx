package criterias

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/updates/criterias"
)

// Adapter represents an adapter
type Adapter struct {
	builder criterias.Builder
}

func createAdapter(
	builder criterias.Builder,
) criterias.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins criterias.Criteria) ([]byte, error) {
	ptr, err := app.CriteriaToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to an instance
func (app *Adapter) ToInstance(bytes []byte) (criterias.Criteria, error) {
	ins := new(Criteria)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCriteria(*ins)
}

// CriteriaToStruct converts a criteria to struct
func (app *Adapter) CriteriaToStruct(ins criterias.Criteria) (*Criteria, error) {
	output := Criteria{
		ChangeSigner:    ins.ChangeSigner(),
		ChangeEncryptor: ins.ChangeEncryptor(),
	}

	if ins.HasUsername() {
		output.Username = ins.Username()
	}

	if ins.HasPassword() {
		output.Password = ins.Password()
	}

	return &output, nil
}

// StructToCriteria converts a struct to criteria
func (app *Adapter) StructToCriteria(str Criteria) (criterias.Criteria, error) {
	builder := app.builder.Create()
	if str.ChangeSigner {
		builder.ChangeSigner()
	}

	if str.ChangeEncryptor {
		builder.ChangeEncryptor()
	}

	if str.Username != "" {
		builder.WithUsername(str.Username)
	}

	if str.Password != "" {
		builder.WithPassword(str.Password)
	}

	return builder.Now()
}
