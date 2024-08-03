package assignables

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executables"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists"
)

type builder struct {
	hashAdapter hash.Adapter
	bytes       bytes.Bytes
	constant    constants.Constant
	crypto      cryptography.Cryptography
	compiler    compilers.Compiler
	execution   executions.Execution
	list        lists.List
	executable  executables.Executable
	variable    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		constant:    nil,
		crypto:      nil,
		compiler:    nil,
		execution:   nil,
		list:        nil,
		executable:  nil,
		variable:    "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes bytes.Bytes) Builder {
	app.bytes = bytes
	return app
}

// WithConsant adds a constant to the builder
func (app *builder) WithConsant(constant constants.Constant) Builder {
	app.constant = constant
	return app
}

// WithCryptography adds a cryptography to the builder
func (app *builder) WithCryptography(cryptography cryptography.Cryptography) Builder {
	app.crypto = cryptography
	return app
}

// WithCompiler adds a compiler to the builder
func (app *builder) WithCompiler(compiler compilers.Compiler) Builder {
	app.compiler = compiler
	return app
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution executions.Execution) Builder {
	app.execution = execution
	return app
}

// WithList adds a list to the builder
func (app *builder) WithList(list lists.List) Builder {
	app.list = list
	return app
}

// WithExecutable adds an executable to the builder
func (app *builder) WithExecutable(executable executables.Executable) Builder {
	app.executable = executable
	return app
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// Now builds a new Assignable instance
func (app *builder) Now() (Assignable, error) {
	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes.Hash().Bytes())
	}

	if app.constant != nil {
		data = append(data, []byte("constant"))
		data = append(data, app.constant.Hash().Bytes())
	}

	if app.crypto != nil {
		data = append(data, []byte("crypto"))
		data = append(data, app.crypto.Hash().Bytes())
	}

	if app.compiler != nil {
		data = append(data, []byte("compiler"))
		data = append(data, app.compiler.Hash().Bytes())
	}

	if app.execution != nil {
		data = append(data, []byte("execution"))
		data = append(data, app.execution.Hash().Bytes())
	}

	if app.list != nil {
		data = append(data, []byte("list"))
		data = append(data, app.list.Hash().Bytes())
	}

	if app.executable != nil {
		data = append(data, []byte("executable"))
		data = append(data, app.executable.Hash().Bytes())
	}

	if app.variable != "" {
		data = append(data, []byte("variable"))
		data = append(data, []byte(app.variable))
	}

	if len(data) != 2 {
		return nil, errors.New("the Assignable is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createAssignableWithBytes(*pHash, app.bytes), nil
	}

	if app.constant != nil {
		return createAssignableWithConstant(*pHash, app.constant), nil
	}

	if app.crypto != nil {
		return createAssignableWithCryptography(*pHash, app.crypto), nil
	}

	if app.list != nil {
		return createAssignableWithList(*pHash, app.list), nil
	}

	if app.execution != nil {
		return createAssignableWithExecution(*pHash, app.execution), nil
	}

	if app.executable != nil {
		return createAssignableWithExecutable(*pHash, app.executable), nil
	}

	if app.variable != "" {
		return createAssignableWithVariable(*pHash, app.variable), nil
	}

	return createAssignableWithCompiler(*pHash, app.compiler), nil
}
