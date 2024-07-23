package compilers

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/compilers"
)

// Adapter represents an adapter
type Adapter struct {
	builder compilers.Builder
}

func createAdapter(
	builder compilers.Builder,
) compilers.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts a compiler to bytes
func (app *Adapter) ToBytes(ins compilers.Compiler) ([]byte, error) {
	ptr, err := app.CompilerToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to compiler
func (app *Adapter) ToInstance(data []byte) (compilers.Compiler, error) {
	ins := new(Compiler)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToCompiler(*ins)
}

// CompilerToStruct converts a compiler to struct
func (app *Adapter) CompilerToStruct(ins compilers.Compiler) (*Compiler, error) {
	out := Compiler{}
	if ins.IsCompile() {
		out.Compile = ins.Compile()
	}

	if ins.IsDecompile() {
		out.Decompile = ins.Decompile()
	}

	return &out, nil
}

// StructToCompiler converts a struct to compiler
func (app *Adapter) StructToCompiler(str Compiler) (compilers.Compiler, error) {
	builder := app.builder.Create()
	if str.Compile != "" {
		builder.WithCompile(str.Compile)
	}

	if str.Decompile != "" {
		builder.WithDecompile(str.Decompile)
	}

	return builder.Now()
}
