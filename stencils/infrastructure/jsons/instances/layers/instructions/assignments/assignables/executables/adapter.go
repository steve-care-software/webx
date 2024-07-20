package executables

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executables"
)

// Adapter represents the executable adapter
type Adapter struct {
	builder executables.Builder
}

func createAdapter(
	builder executables.Builder,
) executables.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instancte to bytes
func (app *Adapter) ToBytes(ins executables.Executable) ([]byte, error) {
	str := app.ExecutableToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to sn instance
func (app *Adapter) ToInstance(data []byte) (executables.Executable, error) {
	ins := new(Executable)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToExecutable(*ins)
}

// ExecutableToStruct converts an executable to struct
func (app *Adapter) ExecutableToStruct(ins executables.Executable) Executable {
	out := Executable{}
	if ins.IsLocal() {
		out.Local = ins.Local()
	}

	if ins.IsRemote() {
		out.Remote = ins.Remote()
	}

	return out
}

// StructToAssignment converts a struct to assignment
func (app *Adapter) StructToExecutable(str Executable) (executables.Executable, error) {
	builder := app.builder
	if str.Local != "" {
		builder.WithLocal(str.Local)
	}

	if str.Remote != "" {
		builder.WithRemote(str.Remote)
	}

	return builder.Now()
}
