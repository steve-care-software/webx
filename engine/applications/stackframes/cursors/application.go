package cursors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

type application struct {
	variablesBuilder variables.Builder
	variableBuilder  variables.VariableBuilder
	variables        []variables.Variable
	values           map[uint8][]any
}

func createApplication(
	variablesBuilder variables.Builder,
	variableBuilder variables.VariableBuilder,
) Application {
	out := application{
		variablesBuilder: variablesBuilder,
		variableBuilder:  variableBuilder,
		variables:        []variables.Variable{},
		values:           map[uint8][]any{},
	}

	out.ClearAll()
	return &out
}

// Fingerprint fingerprints the cursor
func (app *application) Fingerprint() (variables.Variables, error) {
	return app.variablesBuilder.Create().
		WithList(app.variables).
		Now()
}

// Amount returns the amount of values a kind contains
func (app *application) Amount(kind uint8) (*uint, error) {
	if list, ok := app.values[kind]; ok {
		amount := uint(len(list))
		return &amount, nil
	}

	str := fmt.Sprintf(invalidKindErrPattern, kind)
	return nil, errors.New(str)
}

// Fetch fetches a value from an index and a kind
func (app *application) Fetch(index uint, kind uint8) (any, error) {
	if list, ok := app.values[kind]; ok {
		amount := uint(len(list))
		if index < amount {
			return list[index], nil
		}

		str := fmt.Sprintf("the provided kind (%d) only contains %d values, therefore the index (%d) is invalid", kind, amount, index)
		return nil, errors.New(str)
	}

	str := fmt.Sprintf(invalidKindErrPattern, kind)
	return nil, errors.New(str)
}

// Push pushes a value into a kind
func (app *application) Push(value any, kind uint8) error {
	if list, ok := app.values[kind]; ok {
		app.values[kind] = append(list, value)
		return nil
	}

	str := fmt.Sprintf(invalidKindErrPattern, kind)
	return errors.New(str)
}

// PushAsStringBytes pushes a string bytes as value into a kind
func (app *application) PushAsStringBytes(valueStrAsBytes []byte, kind uint8) error {
	casted, err := app.castToKind(valueStrAsBytes, kind)
	if err != nil {
		return err
	}

	return app.Push(casted, kind)
}

func (app *application) castToKind(valueStrAsBytes []byte, kind uint8) (any, error) {
	if kind == variables.KindUint8 ||
		kind == variables.KindUint16 ||
		kind == variables.KindUint32 ||
		kind == variables.KindUint64 ||
		kind == variables.KindInt8 ||
		kind == variables.KindInt16 ||
		kind == variables.KindInt32 ||
		kind == variables.KindInt64 {
		value, err := strconv.Atoi(string(valueStrAsBytes))
		if err != nil {
			return nil, err
		}

		switch kind {
		case variables.KindUint8:
			return uint8(value), nil
		case variables.KindUint16:
			return uint8(value), nil
		case variables.KindUint32:
			return uint8(value), nil
		case variables.KindUint64:
			return uint8(value), nil
		case variables.KindInt8:
			return uint8(value), nil
		case variables.KindInt16:
			return uint8(value), nil
		case variables.KindInt32:
			return uint8(value), nil
		case variables.KindInt64:
			return uint8(value), nil
		}
	}

	if kind == variables.KindFloat32 ||
		kind == variables.KindFloat64 {
		value, err := strconv.ParseFloat(string(valueStrAsBytes), 10)
		if err != nil {
			return nil, err
		}

		switch kind {
		case variables.KindFloat32:
			return float32(value), nil
		case variables.KindFloat64:
			return float64(value), nil
		}
	}

	if kind == variables.KindBool {
		str := strings.ToLower(string(valueStrAsBytes))
		if str == "true" {
			return true, nil
		}

		return false, nil
	}

	if kind == variables.KindString {
		return string(valueStrAsBytes), nil
	}

	str := fmt.Sprintf("the provided kind (%d) do not support the push as string bytes method", kind)
	return nil, errors.New(str)
}

/*

const (
	// KindUint8 represents the uint8
	KindUint8 (uint8) = iota

	// KindUint16 represents the uint16
	KindUint16

	// KindUint32 represents the uint32
	KindUint32

	// KindUint64 represents the uint64
	KindUint64

	// KindInt8 represents the int8
	KindInt8

	// KindInt16 represents the int16
	KindInt16

	// KindInt32 represents the int32
	KindInt32

	// KindInt64 represents the int64
	KindInt64

	// KindFloat32 represents the float32
	KindFloat32

	// KindFloat64 represents the float64
	KindFloat64

	// KindBool represents the bool
	KindBool

	// KindString represents the string
	KindString

	// KindStack represents the stack
	KindStack
)

*/
// Save saves a value to a variable
func (app *application) Save(index uint, kind uint8, variable string, replaceIfExists bool) error {
	value, err := app.Fetch(index, kind)
	if err != nil {
		return err
	}

	builder := app.variableBuilder.Create().
		WithName(variable).
		WithKind(kind).
		WithValue(value)

	if replaceIfExists {
		builder.ReplaceIfExists()
	}

	variableIns, err := builder.Now()
	if err != nil {
		return err
	}

	app.variables = append(app.variables, variableIns)
	return nil
}

// Remove removes a value from its list
func (app *application) Remove(index uint, kind uint8) error {
	if list, ok := app.values[kind]; ok {
		app.values[kind] = append(list[:index], list[index+1:]...)
		return nil
	}

	str := fmt.Sprintf(invalidKindErrPattern, kind)
	return errors.New(str)
}

// Clear clears all values of a kind
func (app *application) Clear(kind uint8) error {
	if _, ok := app.values[kind]; ok {
		app.values[kind] = []any{}
		return nil
	}

	str := fmt.Sprintf(invalidKindErrPattern, kind)
	return errors.New(str)
}

// ClearAll clears all values in all kinds
func (app *application) ClearAll() error {
	app.variables = []variables.Variable{}
	app.values = map[uint8][]any{
		variables.KindUint8:   []any{},
		variables.KindUint16:  []any{},
		variables.KindUint32:  []any{},
		variables.KindUint64:  []any{},
		variables.KindInt8:    []any{},
		variables.KindInt16:   []any{},
		variables.KindInt32:   []any{},
		variables.KindInt64:   []any{},
		variables.KindFloat32: []any{},
		variables.KindFloat64: []any{},
		variables.KindBool:    []any{},
		variables.KindString:  []any{},
	}

	return nil
}
