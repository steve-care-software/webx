package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/states/domain/databases/metadatas"
)

// MetaDataAdapter represents a metaData adapter
type MetaDataAdapter struct {
	builder metadatas.Builder
}

func createMetaDataAdapter(
	builder metadatas.Builder,
) metadatas.Adapter {
	out := MetaDataAdapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *MetaDataAdapter) ToBytes(ins metadatas.MetaData) ([]byte, error) {
	str := app.MetaDataToStruct(ins)
	return json.Marshal(str)
}

// ToInstance converts bytes to instance
func (app *MetaDataAdapter) ToInstance(bytes []byte) (metadatas.MetaData, error) {
	ptr := new(MetaData)
	err := json.Unmarshal(bytes, ptr)
	if err != nil {
		return nil, err
	}

	return app.StructToMetaData(*ptr)
}

// StructToMetaData converts struct to metadata
func (app *MetaDataAdapter) StructToMetaData(str MetaData) (metadatas.MetaData, error) {
	return app.builder.Create().
		WithName(str.Name).
		WithDescription(str.Description).
		WithPath(str.Path).
		Now()
}

// MetaDataToStruct converts metaData to struct
func (app *MetaDataAdapter) MetaDataToStruct(ins metadatas.MetaData) MetaData {
	return MetaData{
		Name:        ins.Name(),
		Description: ins.Description(),
		Path:        ins.Path(),
	}
}
