package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type contentKeysAdapter struct {
	adapter ContentKeyAdapter
	builder ContentKeysBuilder
}

func createContentKeysAdapter(
	adapter ContentKeyAdapter,
	builder ContentKeysBuilder,
) ContentKeysAdapter {
	out := contentKeysAdapter{
		adapter: adapter,
		builder: builder,
	}

	return &out
}

// ToContent converts ContentKeys to bytes
func (app *contentKeysAdapter) ToContent(ins ContentKeys) ([]byte, error) {
	list := ins.List()
	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(len(list)))

	output := []byte{}
	output = append(output, lengthBytes...)

	for _, oneContentKey := range list {
		content, err := app.adapter.ToContent(oneContentKey)
		if err != nil {
			return nil, err
		}

		output = append(output, content...)
	}

	return output, nil
}

// ToContentKeys converts bytes to ContentKeys
func (app *contentKeysAdapter) ToContentKeys(content []byte) (ContentKeys, error) {
	smallest := 8 + contentKeySize
	if len(content) < smallest {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a ContentKey instance, %d provided", smallest, len(content))
		return nil, errors.New(str)
	}

	list := []ContentKey{}
	length := int(binary.LittleEndian.Uint64(content[:8]))
	for i := 0; i < length; i++ {
		beginsOn := 8 + (i * contentKeySize)
		endsOn := beginsOn + contentKeySize
		ins, err := app.adapter.ToContentKey(content[beginsOn:endsOn])
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().WithList(list).Now()
}
