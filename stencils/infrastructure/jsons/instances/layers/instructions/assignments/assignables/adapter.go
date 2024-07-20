package assignables

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables"
	json_bytes "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography"
	json_executables "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executables"
	json_executions "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions"
	json_lists "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists"
)

// Adapter represents an adapter
type Adapter struct {
	executionAdapter    *json_executions.Adapter
	bytesAdapter        *json_bytes.Adapter
	compilerAdapter     *json_compiler.Adapter
	constantAdapter     *json_constants.Adapter
	cryptographyAdapter *json_cryptography.Adapter
	listAdapter         *json_lists.Adapter
	executableAdapter   *json_executables.Adapter
	builder             assignables.Builder
}

func createAdapter(
	executionAdapter *json_executions.Adapter,
	bytesAdapter *json_bytes.Adapter,
	compilerAdapter *json_compiler.Adapter,
	constantAdapter *json_constants.Adapter,
	cryptographyAdapter *json_cryptography.Adapter,
	listAdapter *json_lists.Adapter,
	executableAdapter *json_executables.Adapter,
	builder assignables.Builder,
) assignables.Adapter {
	out := Adapter{
		executionAdapter:    executionAdapter,
		bytesAdapter:        bytesAdapter,
		compilerAdapter:     compilerAdapter,
		constantAdapter:     constantAdapter,
		cryptographyAdapter: cryptographyAdapter,
		listAdapter:         listAdapter,
		executableAdapter:   executableAdapter,
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
func (app *Adapter) ToInstance(data []byte) (assignables.Assignable, error) {
	ins := new(Assignable)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
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

	if ins.IsList() {
		ptr, err := app.listAdapter.ListToStruct(ins.List())
		if err != nil {
			return nil, err
		}

		out.List = ptr
	}

	if ins.IsExecution() {
		execution := app.executionAdapter.ExecutionToStruct(ins.Execution())
		out.Execution = &execution
	}

	if ins.IsExecutable() {
		executable := app.executableAdapter.ExecutableToStruct(ins.Executable())
		out.Executable = &executable
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

	if str.List != nil {
		ins, err := app.listAdapter.StructToList(*str.List)
		if err != nil {
			return nil, err
		}

		builder.WithList(ins)
	}

	if str.Execution != nil {
		ins, err := app.executionAdapter.StructToExecution(*str.Execution)
		if err != nil {
			return nil, err
		}

		builder.WithExecution(ins)
	}

	if str.Executable != nil {
		ins, err := app.executableAdapter.StructToExecutable(*str.Executable)
		if err != nil {
			return nil, err
		}

		builder.WithExecutable(ins)
	}

	return builder.Now()
}
