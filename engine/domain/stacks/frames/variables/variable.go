package variables

type variable struct {
	name            string
	value           any
	kind            uint8
	replaceIfExists bool
}

func createVariable(
	name string,
	value any,
	kind uint8,
	replaceIfExists bool,
) Variable {
	out := variable{
		name:            name,
		value:           value,
		kind:            kind,
		replaceIfExists: replaceIfExists,
	}

	return &out
}

// Name returns the name
func (obj *variable) Name() string {
	return obj.name
}

// Value returns the value
func (obj *variable) Value() any {
	return obj.value
}

// Kind returns the kind
func (obj *variable) Kind() uint8 {
	return obj.kind
}

// ReplaceIfExists returns true if replaceIfExists, false otherwise
func (obj *variable) ReplaceIfExists() bool {
	return obj.replaceIfExists
}
