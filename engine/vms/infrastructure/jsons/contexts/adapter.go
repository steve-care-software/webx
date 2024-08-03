package contexts

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/contexts"
)

type adapter struct {
	builder     contexts.Builder
	hashAdapter hash.Adapter
}

func createAdapter(
	builder contexts.Builder,
	hashAdapter hash.Adapter,
) contexts.Adapter {
	out := adapter{
		builder:     builder,
		hashAdapter: hashAdapter,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *adapter) ToBytes(ins contexts.Context) ([]byte, error) {
	ptr, err := app.contextToStruct(ins)
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
func (app *adapter) ToInstance(data []byte) (contexts.Context, error) {
	ins := new(Context)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.structToContext(*ins)
}

func (app *adapter) contextToStruct(ins contexts.Context) (*Context, error) {
	executions := []string{}
	list := ins.Executions()
	for _, oneIns := range list {
		executions = append(executions, oneIns.String())
	}

	context := Context{
		Identifier: ins.Identifier(),
		Executions: executions,
	}

	if ins.HasHead() {
		context.Head = ins.Head().String()
	}

	return &context, nil
}

func (app *adapter) structToContext(str Context) (contexts.Context, error) {
	hashes := []hash.Hash{}
	for _, oneStr := range str.Executions {
		pHash, err := app.hashAdapter.FromString(oneStr)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *pHash)
	}

	builder := app.builder.Create().
		WithIdentifier(str.Identifier).
		WithExecutions(hashes)

	if str.Head != "" {
		pHead, err := app.hashAdapter.FromString(str.Head)
		if err != nil {
			return nil, err
		}

		builder.WithHead(*pHead)
	}

	return builder.Now()
}
