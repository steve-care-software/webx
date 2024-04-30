package assignables

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables"
	json_bytes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
)

// Adapter represents an adapter
type Adapter struct {
	bytesAdapter        *json_bytes.Adapter
	compilerAdapter     *json_compiler.Adapter
	constantAdapter     *json_constants.Adapter
	cryptographyAdapter *json_cryptography.Adapter
	builder             assignables.Builder
}

func createAdapter(
	bytesAdapter *json_bytes.Adapter,
	compilerAdapter *json_compiler.Adapter,
	constantAdapter *json_constants.Adapter,
	cryptographyAdapter *json_cryptography.Adapter,
	builder assignables.Builder,
) assignables.Adapter {
	out := Adapter{
		bytesAdapter:        bytesAdapter,
		compilerAdapter:     compilerAdapter,
		constantAdapter:     constantAdapter,
		cryptographyAdapter: cryptographyAdapter,
		builder:             builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins assignables.Assignable) ([]byte, error) {
	ptr, err := app.AssignableToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (assignables.Assignable, error) {
	ins := new(Assignable)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToAssignable(*ins)
}

// AssignableToStruct converts an assignable to struct
func (app *Adapter) AssignableToStruct(ins assignables.Assignable) (*Assignable, error) {
	out := Assignable{}
	if ins.IsBytes() {
		ptr, err := app.bytesAdapter.BytesToStruct(ins.Bytes())
		if err != nil {
			return nil, err
		}

		out.Bytes = ptr
	}

	if ins.IsConstant() {
		ptr, err := app.constantAdapter.ConstantToStruct(ins.Constant())
		if err != nil {
			return nil, err
		}

		out.Constant = ptr
	}

	if ins.IsCryptography() {
		ptr, err := app.cryptographyAdapter.CryptographyToStruct(ins.Cryptography())
		if err != nil {
			return nil, err
		}

		out.Cryptography = ptr
	}

	if ins.IsCompiler() {
		ptr, err := app.compilerAdapter.CompilerToStruct(ins.Compiler())
		if err != nil {
			return nil, err
		}

		out.Compiler = ptr
	}

	return &out, nil
}

// StructToAssignable converts a struct to assignable
func (app *Adapter) StructToAssignable(str Assignable) (assignables.Assignable, error) {
	builder := app.builder.Create()
	if str.Bytes != nil {
		ins, err := app.bytesAdapter.StructToBytes(*str.Bytes)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(ins)
	}

	if str.Constant != nil {
		ins, err := app.constantAdapter.StructToConstant(*str.Constant)
		if err != nil {
			return nil, err
		}

		builder.WithConsant(ins)
	}

	if str.Cryptography != nil {
		ins, err := app.cryptographyAdapter.StructToCryptography(*str.Cryptography)
		if err != nil {
			return nil, err
		}

		builder.WithCryptography(ins)
	}

	if str.Compiler != nil {
		ins, err := app.compilerAdapter.StructToCompiler(*str.Compiler)
		if err != nil {
			return nil, err
		}

		builder.WithCompiler(ins)
	}

	return builder.Now()
}
