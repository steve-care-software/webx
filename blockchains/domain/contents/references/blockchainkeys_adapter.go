package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type blockchainKeysAdapter struct {
	adapter BlockchainKeyAdapter
	builder BlockchainKeysBuilder
}

func createBlockchainKeysAdapter(
	adapter BlockchainKeyAdapter,
	builder BlockchainKeysBuilder,
) BlockchainKeysAdapter {
	out := blockchainKeysAdapter{
		adapter: adapter,
		builder: builder,
	}

	return &out
}

// ToContent converts BlockchainKeys to bytes
func (app *blockchainKeysAdapter) ToContent(ins BlockchainKeys) ([]byte, error) {
	list := ins.List()
	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(len(list)))

	output := []byte{}
	output = append(output, lengthBytes...)

	for _, oneBlockchainKey := range list {
		content, err := app.adapter.ToContent(oneBlockchainKey)
		if err != nil {
			return nil, err
		}

		output = append(output, content...)
	}

	return output, nil
}

// ToBlockchainKeys converts bytes to BlockchainKeys
func (app *blockchainKeysAdapter) ToBlockchainKeys(content []byte) (BlockchainKeys, error) {
	smallest := 8 + blockchainKeySize
	if len(content) < smallest {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a BlockchainKey instance, %d provided", smallest, len(content))
		return nil, errors.New(str)
	}

	list := []BlockchainKey{}
	length := int(binary.LittleEndian.Uint64(content[:8]))
	for i := 0; i < length; i++ {
		beginsOn := 8 + (i * blockchainKeySize)
		endsOn := beginsOn + blockchainKeySize
		ins, err := app.adapter.ToBlockchainKey(content[beginsOn:endsOn])
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().WithList(list).Now()
}
