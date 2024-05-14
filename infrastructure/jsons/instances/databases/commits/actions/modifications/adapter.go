package modifications

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions/modifications/deletes"
)

// Adapter represents a modification adapter
type Adapter struct {
	deleteAdapter       *json_deletes.Adapter
	modificationBuilder modifications.ModificationBuilder
	builder             modifications.Builder
}

func createAdapter(
	deleteAdapter *json_deletes.Adapter,
	modificationBuilder modifications.ModificationBuilder,
	builder modifications.Builder,
) modifications.Adapter {
	out := Adapter{
		deleteAdapter:       deleteAdapter,
		modificationBuilder: modificationBuilder,
		builder:             builder,
	}

	return &out
}

// ToBytes returns instance to bytes
func (app *Adapter) ToBytes(ins modifications.Modifications) ([]byte, error) {
	ptr, err := app.ModificationsToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance returns bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (modifications.Modifications, error) {
	ins := new([]Modification)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToModifications(*ins)
}

// ModificationsToStruct converts a modifications to struct
func (app *Adapter) ModificationsToStruct(ins modifications.Modifications) ([]Modification, error) {
	out := []Modification{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.ModificationToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructToModifications converts a struct to modifications
func (app *Adapter) StructToModifications(list []Modification) (modifications.Modifications, error) {
	output := []modifications.Modification{}
	for _, oneStr := range list {
		ins, err := app.StructToModification(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// ModificationToStruct converts a modification to struct
func (app *Adapter) ModificationToStruct(ins modifications.Modification) (*Modification, error) {
	out := Modification{}
	if ins.IsInsert() {
		encoded := base64.StdEncoding.EncodeToString(ins.Insert())
		out.Insert = encoded
	}

	if ins.IsDelete() {
		str := app.deleteAdapter.DeleteToStruct(ins.Delete())
		out.Delete = &str
	}

	return &out, nil
}

// StructToModification converts a struct to modification
func (app *Adapter) StructToModification(str Modification) (modifications.Modification, error) {
	builder := app.modificationBuilder.Create()
	if str.Insert != "" {
		decoded, err := base64.StdEncoding.DecodeString(str.Insert)
		if err != nil {
			return nil, err
		}

		builder.WithInsert(decoded)
	}

	if str.Delete != nil {
		ins, err := app.deleteAdapter.StructToDelete(*str.Delete)
		if err != nil {
			return nil, err
		}

		builder.WithDelete(ins)
	}

	return builder.Now()
}
