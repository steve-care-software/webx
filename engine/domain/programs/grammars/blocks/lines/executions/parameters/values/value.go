package values

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values/references"

type value struct {
	reference references.Reference
	bytes     []byte
}

func createValueWithReference(
	reference references.Reference,
) Value {
	return createValueInternally(reference, nil)
}

func createValueWithBytes(
	bytes []byte,
) Value {
	return createValueInternally(nil, bytes)
}

func createValueInternally(
	reference references.Reference,
	bytes []byte,
) Value {
	out := value{
		reference: reference,
		bytes:     bytes,
	}

	return &out
}

// IsReference returns true if there is a reference, false otherwise
func (obj *value) IsReference() bool {
	return obj.reference != nil
}

// Reference returns the reference, if any
func (obj *value) Reference() references.Reference {
	return obj.reference
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *value) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *value) Bytes() []byte {
	return obj.bytes
}
