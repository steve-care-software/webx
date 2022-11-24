package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type pointerAdapter struct {
	builder PointerBuilder
}

func createPointerAdapter(
	builder PointerBuilder,
) PointerAdapter {
	out := pointerAdapter{
		builder: builder,
	}

	return &out
}

// ToContent converts pointer to bytes
func (app *pointerAdapter) ToContent(ins Pointer) ([]byte, error) {
	fromBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fromBytes, uint64(ins.From()))

	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(ins.Length()))

	output := []byte{}
	output = append(output, []byte(fromBytes)...)
	output = append(output, []byte(lengthBytes)...)
	return output, nil
}

// ToPointer converts bytes to pointer
func (app *pointerAdapter) ToPointer(content []byte) (Pointer, error) {
	if len(content) != pointerSize {
		str := fmt.Sprintf("the content was expected to contain %d bytes in order to convert to a Pointer instance, %d provided", pointerSize, len(content))
		return nil, errors.New(str)
	}

	from := binary.LittleEndian.Uint64(content[:8])
	length := binary.LittleEndian.Uint64(content[8:])
	return app.builder.Create().
		From(uint(from)).
		WithLength(uint(length)).
		Now()
}
