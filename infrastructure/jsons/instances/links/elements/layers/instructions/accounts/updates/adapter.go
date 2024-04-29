package updates

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates"
	json_criterias "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/accounts/updates/criterias"
)

// Adapter represents an adapter
type Adapter struct {
	criteriaAdapter *json_criterias.Adapter
	builder         updates.Builder
}

func createAdapter(
	criteriaAdapter *json_criterias.Adapter,
	builder updates.Builder,
) updates.Adapter {
	out := Adapter{
		criteriaAdapter: criteriaAdapter,
		builder:         builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins updates.Update) ([]byte, error) {
	ptr, err := app.UpdateToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (updates.Update, error) {
	ins := new(Update)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToUpdate(*ins)
}

// UpdateToStruct converts an update to struct
func (app *Adapter) UpdateToStruct(ins updates.Update) (*Update, error) {
	ptrCriteria, err := app.criteriaAdapter.CriteriaToStruct(ins.Criteria())
	if err != nil {
		return nil, err
	}

	return &Update{
		Criteria:    *ptrCriteria,
		Credentials: ins.Credentials(),
	}, nil
}

// StructToUpdate converts a struct to update
func (app *Adapter) StructToUpdate(str Update) (updates.Update, error) {
	criteria, err := app.criteriaAdapter.StructToCriteria(str.Criteria)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithCriteria(criteria).
		WithCredentials(str.Credentials).
		Now()
}
