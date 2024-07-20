package jsons

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/states/domain/databases/pointers"
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type pointerAdapter struct {
	metaDataAdapter *MetaDataAdapter
	builder         pointers.Builder
	hashAdapter     hash.Adapter
}

func createPointerAdapter(
	metaDataAdapter *MetaDataAdapter,
	builder pointers.Builder,
	hashAdapter hash.Adapter,
) pointers.Adapter {
	out := pointerAdapter{
		metaDataAdapter: metaDataAdapter,
		builder:         builder,
		hashAdapter:     hashAdapter,
	}

	return &out
}

// ToBytes converts pointer to bytes
func (app *pointerAdapter) ToBytes(ins pointers.Pointer) ([]byte, error) {
	str := app.PointerToStruct(ins)
	return json.Marshal(str)
}

// ToInstance converts bytes to pointer
func (app *pointerAdapter) ToInstance(bytes []byte) (pointers.Pointer, error) {
	ptr := new(Pointer)
	err := json.Unmarshal(bytes, ptr)
	if err != nil {
		return nil, err
	}

	return app.StructToPointer(*ptr)
}

// StructToPointer converts struct to pointer
func (app *pointerAdapter) StructToPointer(str Pointer) (pointers.Pointer, error) {
	decoded, err := base64.StdEncoding.DecodeString(str.Head)
	if err != nil {
		return nil, err
	}

	pHead, err := app.hashAdapter.FromBytes(decoded)
	if err != nil {
		return nil, err
	}

	metaData, err := app.metaDataAdapter.StructToMetaData(str.MetaData)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithHead(*pHead).
		WithMetaData(metaData).
		Now()
}

// PointerToStruct converts pointer to struct
func (app *pointerAdapter) PointerToStruct(ins pointers.Pointer) Pointer {
	head := base64.StdEncoding.EncodeToString(ins.Head().Bytes())
	metaDataStr := app.metaDataAdapter.MetaDataToStruct(ins.MetaData())
	return Pointer{
		Head:     head,
		MetaData: metaDataStr,
	}
}
